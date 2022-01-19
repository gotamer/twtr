package config

import (
	"fmt"
	"io"
	"sort"

	"gopkg.in/ini.v1"
)

// Config holds the configured values from the ~/.config/twtxt/config file,
// which define how to read and post tweets, as well as who to follow.
type Config struct {
	Nick                   string
	Twtfile                string
	Twturl                 string
	CheckFollowing         bool
	UsePager               bool
	UseCache               bool
	Porcelain              bool
	DiscloseIdentity       bool
	CharacterLimit         int
	CharacterWarning       int
	LimitTimeline          int
	TimelineUpdateInterval int
	Timeout                float64
	SortAscending          bool
	UseAbsoluteTime        bool
	PreTweetHook           string
	PostTweetHook          string
	Following              map[string]string
}

// New parses a Config for the given reader source. Returns any parsing error
// that occur.
func New(source io.Reader) (*Config, error) {
	file, err := ini.Load(source)
	if err != nil {
		return nil, err
	}

	cfg := Config{
		Nick:                   file.Section("twtxt").Key("nick").String(),
		Twtfile:                file.Section("twtxt").Key("twtfile").String(),
		Twturl:                 file.Section("twtxt").Key("twturl").String(),
		CheckFollowing:         file.Section("twtxt").Key("check_following").MustBool(true),
		UsePager:               file.Section("twtxt").Key("use_pager").MustBool(false),
		UseCache:               file.Section("twtxt").Key("use_cache").MustBool(true),
		Porcelain:              file.Section("twtxt").Key("porcelain").MustBool(false),
		DiscloseIdentity:       file.Section("twtxt").Key("disclose_identity").MustBool(false),
		CharacterLimit:         file.Section("twtxt").Key("character_limit").MustInt(0),
		CharacterWarning:       file.Section("twtxt").Key("character_warn").MustInt(0),
		LimitTimeline:          file.Section("twtxt").Key("limit_timeline").MustInt(20),
		TimelineUpdateInterval: file.Section("twtxt").Key("timeline_update_interval").MustInt(10),
		Timeout:                file.Section("twtxt").Key("timeout").MustFloat64(5.0),
		UseAbsoluteTime:        file.Section("twtxt").Key("use_abs_time").MustBool(false),
		PreTweetHook:           file.Section("twtxt").Key("pre_tweet_hook").String(),
		PostTweetHook:          file.Section("twtxt").Key("post_tweet_hook").String(),
		Following:              make(map[string]string),
	}

	// get twtxt config section
	switch sorting := file.Section("twtxt").Key("sorting").String(); sorting {
	case "descending":
		cfg.SortAscending = false
	case "ascending":
		cfg.SortAscending = true
	}

	// get following config section
	for _, key := range file.Section("following").Keys() {
		cfg.Following[key.Name()] = key.String()
	}

	// return config
	return &cfg, nil
}

// WriteTo writes an existing config to the given writer, allowing the config to be
// saved to a file.
func (c *Config) WriteTo(w io.Writer) (n int64, err error) {
	file := ini.Empty()

	file.Section("twtxt").Key("nick").SetValue(c.Nick)
	file.Section("twtxt").Key("twtfile").SetValue(c.Twtfile)
	file.Section("twtxt").Key("twturl").SetValue(c.Twturl)
	file.Section("twtxt").Key("check_following").SetValue(fmt.Sprintf("%v", c.CheckFollowing))
	file.Section("twtxt").Key("use_pager").SetValue(fmt.Sprintf("%v", c.UsePager))
	file.Section("twtxt").Key("use_cache").SetValue(fmt.Sprintf("%v", c.UseCache))
	file.Section("twtxt").Key("porcelain").SetValue(fmt.Sprintf("%v", c.Porcelain))
	file.Section("twtxt").Key("disclose_identity").SetValue(fmt.Sprintf("%v", c.DiscloseIdentity))
	file.Section("twtxt").Key("character_limit").SetValue(fmt.Sprintf("%v", c.CharacterLimit))
	file.Section("twtxt").Key("character_warning").SetValue(fmt.Sprintf("%v", c.CharacterWarning))
	file.Section("twtxt").Key("limit_timeline").SetValue(fmt.Sprintf("%v", c.LimitTimeline))
	file.Section("twtxt").Key("timeline_update_interval").SetValue(fmt.Sprintf("%v", c.TimelineUpdateInterval))
	file.Section("twtxt").Key("timeout").SetValue(fmt.Sprintf("%.1f", c.Timeout))
	file.Section("twtxt").Key("use_abs_time").SetValue(fmt.Sprintf("%v", c.UseAbsoluteTime))
	file.Section("twtxt").Key("pre_tweet_hook").SetValue(c.PreTweetHook)
	file.Section("twtxt").Key("post_tweet_hook").SetValue(c.PostTweetHook)

	if c.SortAscending {
		file.Section("twtxt").Key("sorting").SetValue("ascending")
	} else {
		file.Section("twtxt").Key("sorting").SetValue("descending")
	}

	i, nicks := 0, make([]string, len(c.Following))
	for nick := range c.Following {
		nicks[i] = nick
		i++
	}

	sort.Strings(nicks)

	for _, nick := range nicks {
		file.Section("following").Key(nick).SetValue(c.Following[nick])
	}

	if n, err = file.WriteTo(w); err != nil {
		return
	}

	return
}
