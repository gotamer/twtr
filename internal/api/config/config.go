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

// Save writes an existing config to the given writer, allowing the config to be
// saved to a file.
func (c *Config) Save(to io.Writer) (err error) {
	var twtxt *ini.Section

	file := ini.Empty()
	fmti := func(i int) string { return fmt.Sprintf("%d", i) }
	fmtf := func(f float64) string { return fmt.Sprintf("%.1f", f) }
	fmtt := func(b bool, t string, f string) string {
		if b {
			return t
		} else {
			return f
		}
	}
	fmtb := func(b bool) string { return fmtt(b, "True", "False") }

	if twtxt, err = file.NewSection("twtxt"); err != nil {
		return
	}

	if _, err = twtxt.NewKey("nick", c.Nick); err != nil {
		return
	}

	if _, err = twtxt.NewKey("twtfile", c.Twtfile); err != nil {
		return
	}

	if _, err = twtxt.NewKey("twturl", c.Twturl); err != nil {
		return
	}

	if _, err = twtxt.NewKey("check_following", fmtb(c.CheckFollowing)); err != nil {
		return
	}

	if _, err = twtxt.NewKey("use_pager", fmtb(c.UsePager)); err != nil {
		return
	}

	if _, err = twtxt.NewKey("use_cache", fmtb(c.UseCache)); err != nil {
		return
	}

	if _, err = twtxt.NewKey("porcelain", fmtb(c.Porcelain)); err != nil {
		return
	}

	if _, err = twtxt.NewKey("disclose_identity", fmtb(c.DiscloseIdentity)); err != nil {
		return
	}

	if _, err = twtxt.NewKey("character_limit", fmti(c.CharacterLimit)); err != nil {
		return
	}

	if _, err = twtxt.NewKey("character_warning", fmti(c.CharacterWarning)); err != nil {
		return
	}

	if _, err = twtxt.NewKey("limit_timeline", fmti(c.LimitTimeline)); err != nil {
		return
	}

	if _, err = twtxt.NewKey("timeline_update_interval", fmti(c.TimelineUpdateInterval)); err != nil {
		return
	}

	if _, err = twtxt.NewKey("timeout", fmtf(c.Timeout)); err != nil {
		return
	}

	if _, err = twtxt.NewKey("sorting", fmtt(c.SortAscending, "ascending", "descending")); err != nil {
		return
	}

	if _, err = twtxt.NewKey("use_abs_time", fmtb(c.UseAbsoluteTime)); err != nil {
		return
	}

	if _, err = twtxt.NewKey("pre_tweet_hook", c.PreTweetHook); err != nil {
		return
	}

	if _, err = twtxt.NewKey("post_tweet_hook", c.PostTweetHook); err != nil {
		return
	}

	if len(c.Following) > 0 {
		var following *ini.Section

		if following, err = file.NewSection("following"); err != nil {
			return
		}

		i, nicks := 0, make([]string, len(c.Following))
		for nick := range c.Following {
			nicks[i] = nick
			i++
		}

		sort.Strings(nicks)

		for _, nick := range nicks {
			if _, err = following.NewKey(nick, c.Following[nick]); err != nil {
				return
			}
		}
	}

	if _, err = file.WriteTo(to); err != nil {
		return
	}

	return
}
