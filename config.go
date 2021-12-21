package main

import (
	"fmt"
	"os"

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

func getConfig() (config, error) {
	// setup default values
	cfg := config{
		checkFollowing:         true,
		useCache:               true,
		limitTimeline:          20,
		timelineUpdateInterval: 10,
		timeout:                5.0,
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
		if nick := twtxt.Key("nick"); nick != nil {
			cfg.nick = nick.String()
		}

		if twtfile := twtxt.Key("twtfile"); twtfile != nil {
			cfg.twtfile = twtfile.String()
		}

		if twturl := twtxt.Key("twturl"); twturl != nil {
			cfg.twturl = twturl.String()
		}

		if checkFollowing := twtxt.Key("checkFollowing"); checkFollowing != nil {
			b, err := checkFollowing.Bool()
			if err != nil {
				return config{}, err
			}

			cfg.checkFollowing = b
		}

		if usePager := twtxt.Key("usePager"); usePager != nil {
			b, err := usePager.Bool()
			if err != nil {
				return config{}, err
			}

			cfg.usePager = b
		}

		if useCache := twtxt.Key("useCache"); useCache != nil {
			b, err := useCache.Bool()
			if err != nil {
				return config{}, err
			}

			cfg.useCache = b
		}

		if porcelain := twtxt.Key("porcelain"); porcelain != nil {
			b, err := porcelain.Bool()
			if err != nil {
				return config{}, err
			}

			cfg.porcelain = b
		}

		if discloseIdentity := twtxt.Key("discloseIdentity"); discloseIdentity != nil {
			b, err := discloseIdentity.Bool()
			if err != nil {
				return config{}, err
			}

			cfg.discloseIdentity = b
		}

		if characterLimit := twtxt.Key("characterLimit"); characterLimit != nil {
			num, err := characterLimit.Int()
			if err != nil {
				return config{}, err
			}

			cfg.characterLimit = num
		}

		if characterWarning := twtxt.Key("characterWarning"); characterWarning != nil {
			num, err := characterWarning.Int()
			if err != nil {
				return config{}, err
			}

			cfg.characterLimit = num
		}

		if limitTimeline := twtxt.Key("limitTimeline"); limitTimeline != nil {
			num, err := limitTimeline.Int()
			if err != nil {
				return config{}, err
			}

			cfg.limitTimeline = num
		}

		if timelineUpdateInterval := twtxt.Key("timelineUpdateInterval"); timelineUpdateInterval != nil {
			num, err := timelineUpdateInterval.Int()
			if err != nil {
				return config{}, err
			}

			cfg.timelineUpdateInterval = num
		}

		if timeout := twtxt.Key("timeout"); timeout != nil {
			num, err := timeout.Float64()
			if err != nil {
				return config{}, err
			}

			cfg.timeout = num
		}

		if sortAscending := twtxt.Key("sortAscending"); sortAscending != nil {
			b, err := sortAscending.Bool()
			if err != nil {
				return config{}, err
			}

			cfg.sortAscending = b
		}

		if useAbsoluteTime := twtxt.Key("useAbsoluteTime"); useAbsoluteTime != nil {
			b, err := useAbsoluteTime.Bool()
			if err != nil {
				return config{}, err
			}

			cfg.useAbsoluteTime = b
		}

		if preTweetHook := twtxt.Key("preTweetHook"); preTweetHook != nil {
			cfg.preTweetHook = preTweetHook.String()
		}

		if postTweetHook := twtxt.Key("postTweetHook"); postTweetHook != nil {
			cfg.postTweetHook = postTweetHook.String()
		}
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
