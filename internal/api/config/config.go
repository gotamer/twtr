package config

import (
	"fmt"
	"io"
	"strings"

	"gopkg.in/ini.v1"
)

func setConfigString(section *ini.Section, key string, value *string) {
	if section.HasKey(key) {
		val := section.Key(key).String()
		*value = val
	}
}

func setConfigBool(section *ini.Section, key string, value *bool) (err error) {
	if section.HasKey(key) {
		val, e := section.Key(key).Bool()
		if e != nil {
			return e
		}

		*value = val
	}

	return
}

func setConfigInt(section *ini.Section, key string, value *int) (err error) {
	if section.HasKey(key) {
		val, e := section.Key(key).Int()
		if e != nil {
			return e
		}

		*value = val
	}

	return
}

func setConfigFloat64(section *ini.Section, key string, value *float64) (err error) {
	if section.HasKey(key) {
		val, e := section.Key(key).Float64()
		if e != nil {
			return e
		}

		*value = val
	}

	return
}

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
	// setup default values
	cfg := Config{
		CheckFollowing:         true,
		UseCache:               true,
		LimitTimeline:          20,
		TimelineUpdateInterval: 10,
		Timeout:                5.0,
		Following:              make(map[string]string),
	}

	file, err := ini.Load(source)
	if err != nil {
		return nil, err
	}

	// get twtxt config section, use defaults if not found
	if twtxt := file.Section("twtxt"); twtxt != nil {
		setConfigString(twtxt, "nick", &cfg.Nick)
		setConfigString(twtxt, "twtfile", &cfg.Twtfile)
		setConfigString(twtxt, "twturl", &cfg.Twturl)

		if err := setConfigBool(twtxt, "check_following", &cfg.CheckFollowing); err != nil {
			return nil, err
		}

		if err := setConfigBool(twtxt, "use_pager", &cfg.UsePager); err != nil {
			return nil, err
		}

		if err := setConfigBool(twtxt, "use_cache", &cfg.UseCache); err != nil {
			return nil, err
		}

		if err := setConfigBool(twtxt, "porcelain", &cfg.Porcelain); err != nil {
			return nil, err
		}

		if err := setConfigBool(twtxt, "disclose_identity", &cfg.DiscloseIdentity); err != nil {
			return nil, err
		}

		if err := setConfigInt(twtxt, "character_limit", &cfg.CharacterLimit); err != nil {
			return nil, err
		}

		if err := setConfigInt(twtxt, "character_warning", &cfg.CharacterWarning); err != nil {
			return nil, err
		}

		if err := setConfigInt(twtxt, "limit_timeline", &cfg.LimitTimeline); err != nil {
			return nil, err
		}

		if err := setConfigInt(twtxt, "timeline_update_interval", &cfg.TimelineUpdateInterval); err != nil {
			return nil, err
		}

		if err := setConfigFloat64(twtxt, "timeout", &cfg.Timeout); err != nil {
			return nil, err
		}

		var sorting string
		setConfigString(twtxt, "sorting", &sorting)
		switch strings.ToLower(sorting) {
		case "":
			// skip zero value
		case "descending":
			cfg.SortAscending = false
		case "ascending":
			cfg.SortAscending = true
		default:
			return nil, fmt.Errorf("invalid value for 'sorting': %q", sorting)
		}

		if err := setConfigBool(twtxt, "use_abs_time", &cfg.UseAbsoluteTime); err != nil {
			return nil, err
		}

		setConfigString(twtxt, "pre_tweet_hook", &cfg.PreTweetHook)
		setConfigString(twtxt, "post_tweet_hook", &cfg.PostTweetHook)
	}

	// get following config section, skip if not found
	if following := file.Section("following"); following != nil {
		for _, key := range following.Keys() {
			cfg.Following[key.Name()] = key.String()
		}
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
	fmtf := func(f float64) string { return fmt.Sprintf("%f", f) }
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

	if _, err = twtxt.NewKey("twtxt", c.Twtfile); err != nil {
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

		for nick, url := range c.Following {
			if _, err = following.NewKey(nick, url); err != nil {
				return
			}
		}
	}

	if _, err = file.WriteTo(to); err != nil {
		return
	}

	return
}
