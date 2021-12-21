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
//     twtr quickstart
//     twtr timeline
//     twtr following
//     twtr follow     [[<nickname>] <feed>]
//     twtr unfollow   [[<nickname>] <feed>]
//     twtr tweet      [<message>]
//     twtr view       [<nickname>|<feed>]
//     twtr config     [<key> [<value>]]
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
//     --version            Show the version and exit.
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
// Quickstart wizard for setting up your twtxt config.
//
//     twtr quickstart
//
// Retrieve your personal timeline.
//
//     twtr timeline
//
// Return the list of sources you're following.
//
//     twtr following
//
// Add a new source to your followings.
//
//     twtr follow [[<nickname>] <feed>]
//
// Remove an existing source from your following list.
//
//     twtr unfollow [[<nickname>] <feed>]
//
// Appends a new tweet to you twtxt file.
//
//     twtr tweet [<message>]
//
// View a specific feed.
//
//     twtr view [<nickname>|<feed>]
//
// Edit your twtxt configuration.
//
//     twtr config [<key> [<value>]]
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
//     pre_tweet_hook           = "spc nickname@example.com/twtxt.txt {twtxt}"
//     post_tweet_hook          = "spc {twtxt} nickname@example.com/twtxt.txt"
//
//     [following]
//     alice = https://example.com/user/alice/path/to/twtxt.txt
//     bob   = https://example.com/user/bob/path/to/twtxt.txt
//
// The [twtxt] section contains the settings for your "account", what your
// nickname is, where your twtxt file located, etc...
//
//     Option:                   Type:  Default:    Description:
//     nick                      TEXT               Your nickname, will be
//                                                  displayed in your timeline.
//
//     twtfile                   TEXT               Path to your local twtxt
//                                                  files, it should be user
//                                                  writeable and preferably
//                                                  in your user directory.
//
//     twturl                    TEXT               URL to your public twtxt
//                                                  file, this is the same URL
//                                                  that people will follow you
//                                                  by.
//
//     check_following           BOOL   true        Should twtr try to resolve
//                                                  URLs when listing your
//                                                  followings?
//
//     use_pager                 BOOL   false       Should twtr use a pager
//                                                  (i.e. less) to display your
//                                                  timeline?
//
//     use_cache                 BOOL   true        Should twtr cache remote
//                                                  twtxt files locally?
//
//     porcelain                 BOOL   false       Should twtr format output in
//                                                  an easy to parse format?
//
//     disclose_identity         BOOL   false       Should twtr include your
//                                                  nickname and twturl in the
//                                                  user-agent.
//
//     character_limit           INT                Shorten incoming tweets with
//                                                  more characters than this
//                                                  limit. If set to 0 (zero) or
//                                                  left unset, this will
//                                                  default to not shortening
//                                                  tweets at all.
//
//     character_warning         INT                Warn when you outgoing
//                                                  tweets exceed this length.
//                                                  Set to 0 (zero) or leave
//                                                  unset to disable the
//                                                  warning completely.
//
//     limit_timeline            INT    20          Limit the timeline history,
//                                                  set to 0 (zero) to always
//                                                  show the full history.
//
//     timeline_update_interval  INT    10          Time in seconds until a
//                                                  cache is considered out of
//                                                  date.
//
//     timeout                   FLOAT  5.0         Maximum time a http request
//                                                  is allows to take.
//
//     sorting                   TEXT   descending  How to sort the timeline,
//                                                  descending or ascending.
//
//     use_abs_time              BOOL   false       Use absolute date times in
//                                                  your timeline, defaults to
//                                                  relative, i.e. X minutes ago
//                                                  or Y hours ago.
//
//     pre_tweet_hook            TEXT               Command to be executed
//                                                  before tweeting.
//
//     post_tweet_hook           TEXT               Command to be executed after
//                                                  tweeting.
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

import (
	"fmt"
	"os"
	"path"
)

var (
	self    string = "twtr"
	conf    string
	verbose bool
	version string = "v0.0.0"
	usage   string = `Usage: %s COMMAND [OPTIONS] [ARGS...]

Decentralized, minimalist microblogging for hackers.

Options:
    -c, --config PATH  Specify a custom config file location.
    -v, --version      Enable verbose output for debugging.
    --version          Show the version and exit.
    -h, --help         Show this message and exit.

Command:
    timeline   Retrieve your personal timeline.
    following  View the sources that you are following.
    follow     Add a new source to your followings.
    unfollow   Remove an existing source from your list.
    tweet      Send out a message into the void.
    view       View a source that you follow.
    config     Update your configuration.
`
)

var commands map[string]func(...string) error = map[string]func(...string) error{}

func main() {
	self = path.Base(os.Args[0])
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		switch arg := args[i]; arg {
		case "-c", "--config":
			if i++; i >= len(args) {
				panic("config path not given")
			}

			conf = args[i]
		case "-v", "--verbose":
			verbose = true
		case "--version":
			fmt.Println(version)
			os.Exit(0)
		case "-h", "--help":
			fmt.Fprintf(os.Stderr, usage, self)
			os.Exit(0)
		default:
			if cmd, ok := commands[arg]; ok {
				if err := cmd(args[i:]...); err != nil {
					panic(err)
				}

				os.Exit(0)
			}

			panic("unknown command or flag: '" + arg + "'")
		}
	}
}
