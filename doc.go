// Decentralized, minimalist microblogging for hackers.
//
// NAME
//
// twtr - a decentralized microblogging client.
//
// SYNOPSIS
//
// The general syntax for twtr is:
//
//     twtr COMMAND [OPTIONS] [ARGS ...]
//
// Where the command specified the action to preform, modified by optional
// flags. Each command is different, see the SUBCOMMANDS section for a more
// detailed breakdown of each command.
//
// If you want to use twtr but don't want to read the entire documentation, try
// the quickstart subcommand to just start using twtr (and twtxt) as quickly as
// possible.
//
//     twtr quickstart
//
// This is ask you some questions and generate a config file for you, so you can
// start tweeting within three minutes.
//
// DESCRIPTION
//
// twtr is a drop in replacement for the original twtxt client.
//
// You want to get some thoughts out on the internet in a convenient and slick
// way, while also following the gibberish of others? Instead of signing up to a
// closed source and regulated platform, getting your status updates out is as
// easy as adding a line to a publicly accessible text file. The URL pointing to
// this file is your identity, your account. twtr then tracks these text files,
// like a feedreader, and builds your unique timeline from the text files you
// follow. The format is simple, human readable, and integrates well with UNIX
// command line tools.
//
// OPTIONS
//
// twtr is configured with a config file, to manage multiple twtxt feeds, it is
// possible to specify a different config file than the default. These options
// are special because they can be given before a command.
//
//     -c, --config PATH    Specify a custom configuration file location.
//     -v, --verbose        Enable verbose output for debugging purposes.
//         --version        Show the version and exit.
//     -h, --help           Show a help message and exit.
//
// Each subcommand also has its own options, see the SUBCOMMAND section for more
// fine grained control, for example, the help flag will show a general usage
// message for the whole of twtr if called as:
//
//     twtr --help
//
// To see more details surrounding a specific subcommand, the help flag can be
// passed to that subcommand, for example, calling config as:
//
//     twtr config --help
//
// SUBCOMMANDS
//
// The following are the subcommands of twtr, with their command line syntax,
// see the respective section of each subcommand for further information
//
//     twtr quickstart [-cfhnuv] [--disclose-identity] [--follow-news]
//     twtr timeline   [-chv] [--limit COUNT] [--sort ascending|descending]
//     twtr following  [-chv]
//     twtr follow     [-chv] [--replace] SOURCE [SOURCES...]
//     twtr unfollow   [-chv] SOURCE [SOURCES...]
//     twtr tweet      [-cfhv] TWEET
//     twtr view       [-chv] SOURCE [SOURCES...]
//     twtr config     [-chv] [--edit]|[--remove KEY]|[KEY [VALUE]]
//
// Note that the -c, -h, and -v flags are universal.
//
// QUICKSTART SYNOPSIS
//
// Quickstart wizard for setting up twtxt.
//
// Usage:
//
//     twtr quickstart [-cfhnuv] [--disclose-identity] [--follow-news]
//
// Options:
//
//     -c, --config PATH        Specify a custom configuration file location.
//         --disclose-identity  Show your nickname and url in the User Agent.
//     -f, --file PATH          Specify a custom twtxt file location.
//         --follow-news        Follow the official twtxt and twtr news feeds.
//     -h, --help               Show this message and exit.
//     -n, --nick NICK          Specify the nickname for your feed.
//     -u, --url URL            Specify the url that your feed will be hosted at.
//     -v, --verbose            Enable verbose output for debugging.
//         --version            Show the version and exit.
//
// TIMELINE SYNOPSIS
//
// Retrieve your personal timeline.
//
// Usage:
//
//     twtr timeline [-chv] [--limit COUNT] [--sort ascending|descending]
//
// Options:
//
//     -c, --config PATH     Specify a custom configuration file location.
//     -h, --help            Show this message and exit.
//         --limit COUNT     Limit the amount of tweets shown.
//         --sort DIRECTION  Sort tweets ascending or descending by timestamp.
//     -v, --verbose         Enable verbose output for debugging.
//         --version         Show the version and exit.
//
// FOLLOWING SYNOPSIS
//
// View the sources that you are following.
//
// Usage:
//
//     twtr following [-chv]
//
// Options:
//
//     -c, --config PATH  Specify a custom configuration file location.
//     -h, --help         Show this message and exit.
//     -v, --verbose      Enable verbose output for debugging.
//         --version      Show the version and exit.
//
// FOLLOW SYNOPSIS
//
// Add a new source to your following.
//
// Usage:
//
//     twtr follow [-chv] [--replace] SOURCE [SOURCES...]
//
// Options:
//     -c, --config PATH  Specify a custom configuration file location.
//     -h, --help         Show this message and exit.
//         --replace      Replace duplicates instead of returning an error.
//     -v, --verbose      Enable verbose output for debugging.
//         --version      Show the version and exit.
//
// Sources:
//
// At least one SOURCE must be given (unless called with -h), each SOURCE
// consists of a NICK and a URL. Allowed formats are NICK@URL or NICK URL, if
// you don't know the nickname of a SOURCE, you can make one up, or use the
// domain part of the URL (this can be easily changed later).
//
// UNFOLLOW SYNOPSIS
//
// Remove an existing source from your list.
//
// Usage:
//
//     twtr unfollow [-chv] SOURCE [SOURCES...]
//
// Options:
//
//     -c, --config PATH  Specify a custom configuration file location.
//     -h, --help         Show this message and exit.
//     -v, --verbose      Enable verbose output for debugging.
//         --version      Show the version and exit.
//
// Sources:
//
// At least one SOURCE must be given (unless called with -h), each SOURCE
// consists of a NICK and a URL. Allowed formats are NICK@URL or NICK URL, if
// you don't know the nickname of a SOURCE, you can make one up, or use the
// domain part of the URL (this can be easily changed later).
//
// TWEET SYNOPSIS
//
// Send out a message into the void.
//
// Usage:
//
//     twtr tweet [-cfhv] TWEET
//
// Options:
//
//     -c, --config PATH  Specify a custom configuration file location.
//     -f, --file PATH    Specify a custom twtxt file location.
//     -h, --help         Show this message and exit.
//     -v, --verbose      Enable verbose output for debugging.
//         --version      Show the version and exit.
//
// VIEW SYNOPSIS
//
// View a source that you follow.
//
// Usage:
//
//     twtr view [-chv] SOURCE [SOURCES...]
//
// Options:
//     -c, --config PATH  Specify a custom configuration file location.
//     -h, --help         Show this message and exit.
//     -v, --verbose      Enable verbose output for debugging.
//         --version      Show the version and exit.
//
// Sources:
//
// At least one SOURCE must be given (unless called with -h), each SOURCE
// consists of a NICK and a URL. Allowed formats are NICK@URL or NICK URL, if
// you don't know the nickname of a SOURCE, you can make one up, or use the
// domain part of the URL (this can be easily changed later).
//
// CONFIG SYNOPSIS
//
// Update your configuration.
//
// Usage:
//
//  twtr config [-chv] [--edit]|[--remove KEY]|[KEY [VALUE]]
//
// Options:
//
//     -c, --config PATH  Specify a custom configuration file location.
//         --edit         Edit the configuration file manually.
//     -h, --help         Show this message and exit.
//         --remove KEY   Remove a configuration by its KEY, e.g. twtxt.nick.
//     -v, --verbose      Enable verbose output for debugging.
//         --version      Show the version and exit.
//
// See the CONFIGURATION section for a list of available options.
//
// EXIT STATUS
//
// The twtr command exits 0 on success, and >0 if an error occurs.
//
// FILES
//
// The twtxt configuration file is located in the user's configuration
// directory. On most UNIX systems this defaults to:
//
//     ~/.config/twtxt/config
//
// Except on macOS, where the default is:
//
//     ~/Library/Application Support/twtxt/config
//
// More generally this the default is defined by the XDG standard, see also the
// ENVIRONMENT section.
//
//     $XDG_CONFIG_HOME/twtxt/config
//
// See the CONFIGURATION section for more details on the config file.
//
// CONFIGURATION
//
// The configuration file is a simply INI file with two main sections, [twtxt]
// and [following]. You can set up a basic config file using the quickstart
// wizard. See the config subcommand for more information on how to modify the
// configuration via the twtr tool, or you can edit the file with your
// preferred text editor.
//
//     [twtxt]
//     nick                     = nickname
//     twtfile                  = path/to/twtxt.txt
//     twturl                   = https://example.com/twtxt.txt
//     check_following          = true
//     use_pager                = false
//     use_cache                = true
//     porcelain                = false
//     disclose_identity        = false
//     character_limit          = 140
//     character_warning        = 140
//     limit_timeline           = 20
//     timeline_update_interval = 10
//     timeout                  = 5.0
//     sorting                  = descending
//     use_abs_time             = false
//     pre_tweet_hook           = "scp nickname@example.com/twtxt.txt {twtxt}"
//     post_tweet_hook          = "scp {twtxt} nickname@example.com/twtxt.txt"
//
//     [following]
//     alice = https://example.com/user/alice/path/to/twtxt.txt
//     bob   = https://example.com/user/bob/path/to/twtxt.txt
//
// The [twtxt] section contains the settings for your "account", what your
// nickname is, where your twtxt file located, etc...
//
// Your nickname, will be displayed in your timeline.
//
//     nick
//
// Path to your local twtxt files, it should be writeable and preferably in your
// user directory.
//
//     twtfile
//
// Url to your public twtxt files, this is the same URL that people will follow
// you by.
//
//     twturl
//
// Should twtr try to resolve URLs when listing your followings?
//
//     check_following
//
// Should twtr use a pages (i.e. less) to display your timeline.
//
//     use_pager
//
// Should twtr cache remote twtxt files locally?
//
//     use_cache
//
// Should twtr format output in an easy to parse format?
//
//     porcelain
//
// Should twtr include include your nickname and twturl in the user-agent.
//
//     disclose_identity
//
// Shorten incoming tweets with more characters that this limit. If set to 0
// (zero), or left unset, this will default to not shortening tweets at all.
//
//     character_limit
//
// Warn when your outgoing tweets exceed this length. Set to 0 (zero) or leave
// unset to disable the warning complete.
//
//     character_warning
//
// Limit the timeline history, set to 0 (zero) to always show the full history.
//
//     limit_timeline
//
// Time in seconds until a cached file is considered out of date.
//
//     timeline_update_interval
//
// Maximum time a http request is allowed to take.
//
//     timeout
//
// How to sort the timeline, descending or ascending.
//
//     sorting
//
// Use absolute date times in your timeline, defaults to relative, i.e. X
// minutes ago or Y hours ago.
//
//     use_abs_time
//
// Command to be executed before tweeting.
//
//     pre_tweet_hook
//
// Command to be executed after tweeting.
//
//     post_tweet_hook
//
// The pre/post tweet hooks are executed as system commands, any occurrences of
// "{foo}" will be replaced with the value of that configuration. For example,
// "{twtfile}" will be replaced with the path to your local file.
//
// The [following] section contains all the sources you follow, the keys in
// this section are the nicknames, and the values of those keys are the urls of
// the twtxt files. You can update this section using the (un)follow commands.
//
// ENVIRONMENT
//
// This is the user configuration directory used for the twtxt config file, it
// varies from system to system, to see where this defaults to on your local
// operating system, see the os.UserConfigDir() notes in the "os" package.
//
//     XDG_CONFIG_HOME
//
// CONFORMING TO
//
// twtr conforms to the twtxt file specification, traditionally the file is
// located at https://example.com/path/to/twtxt.txt, however, as not everyone
// has access to a personal website to host their feeds, twtr also supports
// specialised hosting options, such as a GitHub gist.
//
// See https://twtxt.readthedocs.io/en/latest/user/twtxtfile.html for more
// information on the file structure.
//
// NOTES
//
// The original client was written is Python around 2016, and a small user
// base has been built around the twtxt format. Since the format is human
// readable and can be easily used with just shell commands, in addition to
// the original client, many users have written their own or just use the
// echo command.
//
// This client aims to be a complete drop-in replacement for the original
// client, not only to replicate the original feature set, but also to support
// many additions that the community of users have requested. There have also
// been a number of issues with the original client breaking because of
// backwards compatibility issues with the Python language. twtr aims to be a
// permanently supported tool, the Go language protects its backwards
// compatibility, so twtr will work forever!
//
// COPYRIGHT
//
// All rites reversed, use, distribute, and modify freely.
//
// AUTHOR
//
// ~duriny <duriny@envs.net>
//
// BUGS
//
// Probably. Let me know if you find any.
//
// SEE ALSO
//
// twtxt(1) - https://github.com/buckket/twtxt
//
package main
