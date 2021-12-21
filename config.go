package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

type config struct {
	nick                   string
	twtfile                string
	twturl                 string
	checkFollowing         bool
	usePager               bool
	useCache               bool
	porcelain              bool
	discloseIdentity       bool
	characterLimit         int
	characterWarning       int
	limitTimeline          int
	timelineUpdateInterval int
	timeout                float64
	sortAscending          bool
	useAbsoluteTime        bool
	preTweetHook           string
	postTweetHook          string
	following              map[string]string
}

func setConfigString(section *ini.Section, key string, value *string) {
	if section.HasKey(key) {
		val := section.Key(key).String()
		value = &val
	}
}

func setConfigBool(section *ini.Section, key string, value *bool) (err error) {
	if section.HasKey(key) {
		val, e := section.Key(key).Bool()
		if e != nil {
			return e
		}

		value = &val
	}

	return
}

func setConfigInt(section *ini.Section, key string, value *int) (err error) {
	if section.HasKey(key) {
		val, e := section.Key(key).Int()
		if e != nil {
			return e
		}

		value = &val
	}

	return
}

func setConfigFloat64(section *ini.Section, key string, value *float64) (err error) {
	if section.HasKey(key) {
		val, e := section.Key(key).Float64()
		if e != nil {
			return e
		}

		value = &val
	}

	return
}

func getConfig() (config, error) {
	// setup default values
	cfg := config{
		checkFollowing:         true,
		useCache:               true,
		limitTimeline:          20,
		timelineUpdateInterval: 10,
		timeout:                5.0,
		following:              make(map[string]string),
	}

	// use default path if non set
	if conf == "" {
		dir, err := os.UserConfigDir()
		if err != nil {
			return config{}, err
		}

		conf = dir + "/twtxt/config"
	}

	// load configuration file
	file, err := ini.Load(conf)
	if err != nil {
		return config{}, err
	}

	// get twtxt config section, use defaults if not found
	if twtxt := file.Section("twtxt"); twtxt != nil {
		setConfigString(twtxt, "nick", &cfg.nick)
		setConfigString(twtxt, "twtfile", &cfg.twtfile)
		setConfigString(twtxt, "twturl", &cfg.twturl)

		if err := setConfigBool(twtxt, "check_following", &cfg.checkFollowing); err != nil {
			return config{}, err
		}

		if err := setConfigBool(twtxt, "use_pager", &cfg.usePager); err != nil {
			return config{}, err
		}

		if err := setConfigBool(twtxt, "use_cache", &cfg.useCache); err != nil {
			return config{}, err
		}

		if err := setConfigBool(twtxt, "porcelain", &cfg.porcelain); err != nil {
			return config{}, err
		}

		if err := setConfigBool(twtxt, "disclose_identity", &cfg.discloseIdentity); err != nil {
			return config{}, err
		}

		if err := setConfigInt(twtxt, "character_limit", &cfg.characterLimit); err != nil {
			return config{}, err
		}

		if err := setConfigInt(twtxt, "character_warning", &cfg.characterWarning); err != nil {
			return config{}, err
		}

		if err := setConfigInt(twtxt, "limit_timeline", &cfg.limitTimeline); err != nil {
			return config{}, err
		}

		if err := setConfigInt(twtxt, "timeline_update_interval", &cfg.timelineUpdateInterval); err != nil {
			return config{}, err
		}

		if err := setConfigFloat64(twtxt, "timeout", &cfg.timeout); err != nil {
			return config{}, err
		}

		var sorting string
		setConfigString(twtxt, "sorting", &sorting)
		switch strings.ToLower(sorting) {
		case "":
			// skip zero value
		case "descending":
			cfg.sortAscending = false
		case "ascending":
			cfg.sortAscending = true
		default:
			return config{}, fmt.Errorf("Invalid value for 'sorting': %q", sorting)
		}

		if err := setConfigBool(twtxt, "use_abs_time", &cfg.useAbsoluteTime); err != nil {
			return config{}, err
		}

		setConfigString(twtxt, "pre_tweet_hook", &cfg.preTweetHook)
		setConfigString(twtxt, "post_tweet_hook", &cfg.postTweetHook)
	}

	// get following config section, skip if not found
	if following := file.Section("following"); following != nil {
		for _, key := range following.Keys() {
			cfg.following[key.Name()] = key.String()
		}
	}

	// return config
	return cfg, nil
}

func main_config(args ...string) error {
	cfg, err := getConfig()
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", cfg)

	return nil
}
