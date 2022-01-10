package cmd

import (
	"fmt"
	"sort"
	"strings"
)

const (
	lineLength = 80
	indentSize = 8
)

func helpFormatFlags(flags []flag) string {
	// for there no flags then there are no flags
	if len(flags) < 1 {
		return "NONE\n"
	}

	var len1, len2 int

	lines := make([]string, len(flags), len(flags))
	rows := make([][3]string, len(flags), len(flags))

	// get the width maximum of each component
	for i, flag := range flags {
		var col1, col2 string

		if flag.short != "" {
			col1 = fmt.Sprintf("%s,", flag.short)
		}

		if flag.option == "" {
			col2 = fmt.Sprintf("%s", flag.long)
		} else {
			col2 = fmt.Sprintf("%s %s", flag.long, flag.option)
		}

		if l := len(col1); l > len1 {
			len1 = l
		}

		if l := len(col2); l > len2 {
			len2 = l
		}

		rows[i] = [3]string{col1, col2, flag.description}
	}

	// format each component no that we know how wide the columns are
	for i, row := range rows {
		lines[i] = fmt.Sprintf("%-*s %-*s  %s", len1, row[0], len2, row[1], row[2])
	}

	return "\t" + strings.Join(lines, "\n\t") + "\n"
}

func helpFormatText(msg string) string {
	lines := []string{}

	for _, paragraph := range strings.Split(msg, "\n") {
		var line string

		words := strings.Fields(paragraph)

		for _, word := range words {
			if len(line)+len(word)+1 < lineLength-indentSize {
				if len(line) > 0 {
					line += " " + word
				} else {
					line = word
				}
			} else {
				lines = append(lines, line)
				line = word
			}
		}

		lines = append(lines, line)
	}

	return "\t" + strings.Join(lines, "\n\t") + "\n"
}

type command struct {
	name        string
	usage       string
	description string
	flags       []flag
	other       map[string]string
}

func (c command) help(ctx *Context) string {
	// format the base part of the help message
	messages := []string{
		fmt.Sprintf("Usage: %s %s %s\n", ctx.Self, c.name, c.usage),
		fmt.Sprintf("%s\n", c.description),
	}

	// create section for flags and options
	if c.flags != nil {
		messages = append(messages, "Options:\n"+helpFormatFlags(c.flags))
	}

	// start list of all help message sections
	sections := make(map[string]string, len(c.other))
	headings := make([]string, len(sections), len(sections))

	// include any other sections, e.g. "Examples", "Notes", "Bugs" etc
	if c.other != nil {
		for k, v := range c.other {
			headings = append(headings, k)
			sections[k] = helpFormatText(v)
		}
	}

	// sort these other headings
	sort.Strings(headings)

	// add the headings to the message in order
	for _, heading := range headings {
		messages = append(messages, heading+":\n"+sections[heading])
	}

	return strings.Join(messages, "\n")
}

var (
	quickstartCommand command = command{
		name:        "quickstart",
		usage:       "[-cfhnuv] [--disclose-identity] [--follow-news]",
		description: "Quickstart wizard for setting up twtxt.",
		flags: []flag{
			configFlag,
			discloseIdentityFlag,
			fileFlag,
			followNewsFlag,
			helpFlag,
			nickFlag,
			urlFlag,
			verboseFlag,
			versionFlag,
		},
	}
	timelineCommand command = command{
		name:        "timeline",
		usage:       "[-chv] [--limit COUNT] [--sort ascending|descending]",
		description: "Retrieve your personal timeline.",
		flags: []flag{
			configFlag,
			helpFlag,
			limitFlag,
			sortFlag,
			verboseFlag,
			versionFlag,
		},
	}
	followingCommand command = command{
		name:        "following",
		usage:       "[-chv]",
		description: "View the sources that you are following.",
		flags: []flag{
			configFlag,
			helpFlag,
			versionFlag,
			verboseFlag,
		},
	}
	followCommand command = command{
		name:        "follow",
		usage:       "[-chv] [--replace] SOURCE [SOURCES...]",
		description: "Add a new source to your following.",
		flags: []flag{
			configFlag,
			helpFlag,
			replaceFlag,
			verboseFlag,
			versionFlag,
		},
		other: map[string]string{
			"Sources": "At least one SOURCE must be given (unless called with -h), each SOURCE consists of a NICK and a URL. Allowed formats are NICK@URL or NICK URL, if you don't know the nickname of a SOURCE, you can make one up, or use the domain part of the URL (this can be easily changed later).",
		},
	}
	unfollowCommand command = command{
		name:        "unfollow",
		usage:       "[-chv] NICK|URL",
		description: "Remove an existing source form your list",
		flags: []flag{
			configFlag,
			helpFlag,
			verboseFlag,
			versionFlag,
		},
	}
	tweetCommand command = command{
		name:        "tweet",
		usage:       "[-cfhv] TWEET",
		description: "Send out a message into the void.",
		flags: []flag{
			configFlag,
			fileFlag,
			helpFlag,
			verboseFlag,
			versionFlag,
		},
	}
	viewCommand command = command{
		name:        "view",
		usage:       "[-chv] SOURCE [SOURCES...]",
		description: "View a source that you follow.",
		flags: []flag{
			configFlag,
			helpFlag,
			verboseFlag,
			versionFlag,
		},
		other: map[string]string{
			"Sources": "At least one SOURCE must be given (unless called with -h), each SOURCE consists of a NICK and a URL. Allowed formats are NICK@URL or NICK URL, if you don't know the nickname of a SOURCE, you can make one up, or use the domain part of the URL (this can be easily changed later).",
		},
	}
	configCommand command = command{
		name:        "config",
		usage:       "[-chv] [--edit]|[--remove KEY]|[KEY [VALUE]]",
		description: "Update your configuration.",
		flags: []flag{
			configFlag,
			editFlag,
			helpFlag,
			removeFlag,
			verboseFlag,
			versionFlag,
		},
	}
)

var commands []command = []command{
	quickstartCommand,
	timelineCommand,
	followingCommand,
	followCommand,
	unfollowCommand,
	tweetCommand,
	viewCommand,
	configCommand,
}
