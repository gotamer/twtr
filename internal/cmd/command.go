package cmd

import "fmt"

type command struct {
	name        string
	usage       string
	description string
	flags       []flag
	other       map[string]string
}

func (c command) help(ctx *Context) string {
	// TODO format options section
	// TODO format other sections

	return fmt.Sprintf(
		"Usage %s %s %s\n\n%s\n\nOptions:\n\tTODO read c.flags\n",
		ctx.Self,
		c.name,
		c.usage,
		c.description,
	)
}

var commands []command = []command{
	{
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
	},
	{
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
	},
	{
		name:        "following",
		usage:       "[-chv]",
		description: "View the sources that you are following.",
		flags: []flag{
			configFlag,
			helpFlag,
			versionFlag,
			verboseFlag,
		},
	},
	{
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
	},
	{
		name:        "unfollow",
		usage:       "[-chv] NICK|URL",
		description: "Remove an existing source form your list",
		flags: []flag{
			configFlag,
			helpFlag,
			verboseFlag,
			versionFlag,
		},
	},
	{
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
	},
	{
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
	},
	{
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
	},
}
