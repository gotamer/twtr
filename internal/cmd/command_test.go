package cmd

import "testing"

func TestCommand(t *testing.T) {
	tests := []struct {
		command command
		help    string
	}{
		{
			command: quickstartCommand,
			help: `Usage: twtr quickstart [-cfhnuv] [--disclose-identity] [--follow-news]

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
`,
		},
		{
			command: timelineCommand,
			help: `Usage: twtr timeline [-chv] [--limit COUNT] [--sort ascending | descending]

Retrieve your personal timeline.

Options:
	-c, --config PATH     Specify a custom configuration file location.
	-h, --help            Show this message and exit.
	    --limit COUNT     Limit the amount of tweets shown.
	    --sort DIRECTION  Sort tweets ascending or descending by timestamp.
	-v, --verbose         Enable verbose output for debugging.
	    --version         Show the version and exit.
`,
		},
		{
			command: followingCommand,
			help: `Usage: twtr following [-chv]

View the sources that you are following.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	-h, --help         Show this message and exit.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.
`,
		},
		{
			command: followCommand,
			help: `Usage: twtr follow [-chv] [--replace] SOURCE [SOURCES...]

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
`,
		},
		{
			command: unfollowCommand,
			help: `Usage: twtr unfollow [-chv] NICK|URL

Remove an existing source from your list.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	-h, --help         Show this message and exit.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.
`,
		},
		{
			command: tweetCommand,
			help: `Usage: twtr tweet [-cfhv] TWEET

Send out a message into the void.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	-f, --file PATH    Specify a custom twtxt file location.
	-h, --help         Show this message and exit.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.
`,
		},
		{
			command: viewCommand,
			help: `Usage: twtr view [-chv] SOURCE [SOURCES]

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
`,
		},
		{
			command: configCommand,
			help: `Usage: twtr config [-chv] [--edit] [--remove KEY]|[KEY [VALUE]]

Update your configuration.

Options:
	-c, --config PATH  Specify a custom configuration file location.
	    --edit         Edit the configuration file manually.
	-h, --help         Show this message and exit.
	    --remove KEY   Remove a configuration by its KEY, e.g. twtxt.nick.
	-v, --verbose      Enable verbose output for debugging.
	    --version      Show the version and exit.
`,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.command.name, func(t *testing.T) {
			t.Run("help", func(t *testing.T) {
				ctx := Context{Self: "twtr"}

				if have, want := test.command.help(&ctx), test.help; have != want {
					t.Errorf("\nhave:\n%s\n\nwant:\n%s\n", have, want)
				}
			})
		})
	}
}
