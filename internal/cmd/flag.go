package cmd

type flag struct {
	short       string
	long        string
	option      string
	description string
}

var (
	configFlag           flag = flag{"-c", "--config", "PATH", "Specify a custom config file location."}
	helpFlag             flag = flag{"-h", "--help", "", "Show this message and exit."}
	verboseFlag          flag = flag{"-v", "--verbose", "", "Enable verbose output for debugging."}
	fileFlag             flag = flag{"-f", "--file", "PATH", "Specify a custom twtxt file location."}
	nickFlag             flag = flag{"-n", "--nick", "NICK", "Specify the nickname for your feed."}
	urlFlag              flag = flag{"-u", "--url", "URL", "Specify the url that your feed will be hosted at."}
	versionFlag          flag = flag{"", "--version", "", "Show the version and exit."}
	discloseIdentityFlag flag = flag{"", "--disclose-identity", "", "Show your nickname and url in the User Agent."}
	followNewsFlag       flag = flag{"", "--follow-news", "", "Follow the official twtxt and twtr news feeds."}
	limitFlag            flag = flag{"", "--limit", "COUNT", "Limit the amount of tweets shown."}
	sortFlag             flag = flag{"", "--sort", "DIRECTION", "Sort tweets ascending or descending by timestamp."}
	replaceFlag          flag = flag{"", "--replace", "", "Replace duplicates instead of returning an error."}
	editFlag             flag = flag{"", "--edit", "", "Edit the configuration file manually."}
	removeFlag           flag = flag{"", "--remove", "KEY", "Remove a configuration by its KEY, e.g. twtxt.nick."}
)
