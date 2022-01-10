package cmd

import (
	"fmt"
	"sort"
	"strings"
)

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
		// TODO format options section
		messages = append(messages, "Options:\n\tTODO: Format options\n")
	}

	// start list of all help message sections
	sections := make(map[string]string, len(c.other))
	headings := make([]string, len(sections), len(sections))

	// include any other sections, e.g. "Examples", "Notes", "Bugs" etc
	if c.other != nil {
		for k, v := range c.other {
			// TODO format other sections
			headings = append(headings, k)
			sections[k] = v
		}
	}

	// sort these other headings
	sort.Strings(headings)

	// add the headings to the message in order
	for _, heading := range headings {
		messages = append(messages, fmt.Sprintf("%s:\n%s\n", heading, sections[heading]))
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
