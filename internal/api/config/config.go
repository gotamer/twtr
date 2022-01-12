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
