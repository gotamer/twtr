package cmd

import "fmt"

func help(ctx *Context) string {
	usage := `Usage: %s COMMAND [OPTIONS] [ARGS...]

Decentralized, minimalist microblogging for hackers.

Options:
	-c, --config PATH  Specify a custom config file location.
	-h, --help         Show this message and exit.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.

Commands:
	quickstart  Quickstart wizard for setting up twtxt.
	timeline    Retrieve your personal timeline.
	following   View the sources that you are following.
	follow      Add a new source to your followings.
	unfollow    Remove an existing source from your list.
	tweet       Send out a message into the void.
	view        View a source that you follow.
	config      Update your configuration.
`

	return fmt.Sprintf(usage, ctx.Self)
}

func helpQuickstart(ctx *Context) string {
	usage := `Usage: %s quickstart [-cfhnuv] [--disclose-identity] [--follow-news]

Quickstart wizard for setting up twtxt.

Options:
	-c, --config PATH        Specify a custom configuration file location.
	    --disclose-identity  Show your nickname and url in the User Agent.
	-f, --file PATH          Specify a custom twtxt file location.
	    --follow-news        Follow the official twtxt and twtr news feeds.
	-h, --help               Show this message and exit.
	-n, --nick NICK          Specify the nickname for your feed.
	-u, --url URL            Specify the url that your feed will be hosted at.
	-v, --verbose            Enable verbose output for debugging.
	    --version            Show the version and exit.
`

	return fmt.Sprintf(usage, ctx.Self)
}

func helpTimeline(ctx *Context) string {
	usage := `Usage: %s timeline [-chv] [--limit COUNT] [--sort ascending | descending]

Retrieve your personal timeline.

Options:
	-c, --config PATH     Specify a custom configuration file location.
	-h, --help            Show this message and exit.
	    --limit COUNT     Limit the amount of tweets shown.
	    --sort DIRECTION  Sort tweets ascending or descending by timestamp.
	-v, --verbose         Enable verbose output for debugging.
	    --version         Show the version and exit.
`

	return fmt.Sprintf(usage, ctx.Self)
}

func helpFollowing(ctx *Context) string {
	usage := `Usage: %s following [-chv]

View the sources that you are following.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	-h, --help         Show this message and exit.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.
`

	return fmt.Sprintf(usage, ctx.Self)
}

func helpFollow(ctx *Context) string {
	usage := `Usage: %s follow [-chv] [--replace] SOURCE [SOURCES...]

Add a new source to your following.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	-h, --help         Show this message and exit.
	    --replace      Replace duplicates instead of returning error.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.

Sources:
	At least one SOURCE must be given (unless called with -h), each SOURCE
	consists of a NICK and a URL. Allowed formats are NICK@URL or NICK URL, if
	you don't know the nickname of a SOURCE, you can make one up, or use the
	domain part of the URL (this can be easily changed later).
`

	return fmt.Sprint(usage, ctx.Self)
}

func helpUnfollow(ctx *Context) string {
	usage := `Usage: twtr unfollow [-chv] NICK|URL

Remove an existing source from your list.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	-h, --help         Show this message and exit.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.
`

	return fmt.Sprint(usage, ctx.Self)
}

func helpTweet(ctx *Context) string {
	usage := `Usage: twtr tweet [-cfhv] TWEET

Send out a message into the void.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	-f, --file PATH    Specify a custom twtxt file location.
	-h, --help         Show this message and exit.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.
`

	return fmt.Sprint(usage, ctx.Self)
}

func helpView(ctx *Context) string {
	usage := `Usage: twtr view [-chv] SOURCE [SOURCES]

View a source that you follow.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	-h, --help         Show this message and exit.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.

Sources:
	At least one SOURCE must be given (unless called with -h), each SOURCE
	consists of a NICK and a URL. Allowed formats are NICK@URL or NICK URL, if
	you don't know the nickname of a SOURCE, you can make one up, or use the
	domain part of the URL (this can be easily changed later).
`

	return fmt.Sprint(usage, ctx.Self)
}

func helpConfig(ctx *Context) string {
	usage := `Usage: twtr config [-chv] [--edit] [--remove KEY]|[KEY [VALUE]]

Update your configuration.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	    --edit         Edit the configuration file manually.
	-h, --help         Show this message and exit.
	    --remove KEY   Remove a configuration by its KEY, e.g. twtxt.nick.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.

Examples:
	Get config value:
		$ twtr config twtxt.nick
		> YourOldNickName

	Set config value:
		$ twtr config twtxt.nick YourNewNickName

	Remove config value:
		$ twtr config --remove twtxt.nick

	Edit config file manually:
		$ twtr config --edit
`

	return fmt.Sprint(usage, ctx.Self)
}
