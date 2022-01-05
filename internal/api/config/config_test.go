package config_test

import (
	"fmt"
	"testing"

	"internal/api/config"
)

/*
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
*/

func TestNewConfig(t *testing.T) {
	tests := []struct {
		file string
		want config.Config
	}{
		{
			"_test/example1.ini",
			config.Config{
				CheckFollowing:         true,
				UseCache:               true,
				LimitTimeline:          20,
				TimelineUpdateInterval: 10,
				Timeout:                5.0,
				Following:              make(map[string]string),
			},
		},
	}

	for _, test := range tests {
		test := test
		name := test.file
		want := test.want

		t.Run(name, func(t *testing.T) {
			have, err := config.NewConfig(test.file)
			if err != nil {
				t.Fatalf("could not open file '%s': %q", test.file, err)
			}

			if have.Nick != want.Nick {
				t.Errorf("Nick = %q, want %q", have.Nick, want.Nick)
			}

			if have.Twtfile != want.Twtfile {
				t.Errorf("Twtfile = %q, want %q", have.Twtfile, want.Twtfile)
			}

			if have.Twturl != want.Twturl {
				t.Errorf("Twturl = %q, want %q", have.Twturl, want.Twturl)
			}

			if have.CheckFollowing != want.CheckFollowing {
				t.Errorf("CheckFollowing = %t, want %t", have.CheckFollowing, want.CheckFollowing)
			}

			if have.UsePager != want.UsePager {
				t.Errorf("UsePager = %t, want %t", have.UsePager, want.UsePager)
			}

			if have.UseCache != want.UseCache {
				t.Errorf("UseCache = %t, want %t", have.UseCache, want.UseCache)
			}

			if have.Porcelain != want.Porcelain {
				t.Errorf("Porcelain = %t, want %t", have.Porcelain, want.Porcelain)
			}

			if have.DiscloseIdentity != want.DiscloseIdentity {
				t.Errorf("DiscloseIdentity = %t, want %t", have.DiscloseIdentity, want.DiscloseIdentity)
			}

			if have.CharacterLimit != want.CharacterLimit {
				t.Errorf("CharacterLimit = %q, want %q", have.CharacterLimit, want.CharacterLimit)
			}

			if have.CharacterWarning != want.CharacterWarning {
				t.Errorf("CharacterWarning = %q, want %q", have.CharacterWarning, want.CharacterWarning)
			}

			if have.LimitTimeline != want.LimitTimeline {
				t.Errorf("LimitTimeline = %q, want %q", have.LimitTimeline, want.LimitTimeline)
			}

			if have.TimelineUpdateInterval != want.TimelineUpdateInterval {
				t.Errorf("TimelineUpdateInterval = %q, want %q", have.TimelineUpdateInterval, want.TimelineUpdateInterval)
			}

			if have.Timeout != want.Timeout {
				t.Errorf("Timeout = %f, want %f", have.Timeout, want.Timeout)
			}

			if have.SortAscending != want.SortAscending {
				t.Errorf("SortAscending = %t, want %t", have.SortAscending, want.SortAscending)
			}

			if have.UseAbsoluteTime != want.UseAbsoluteTime {
				t.Errorf("UseAbsoluteTime = %t, want %t", have.UseAbsoluteTime, want.UseAbsoluteTime)
			}

			if have.PreTweetHook != want.PreTweetHook {
				t.Errorf("PreTweetHook = %q, want %q", have.PreTweetHook, want.PreTweetHook)
			}

			if have.PostTweetHook != want.PostTweetHook {
				t.Errorf("PostTweetHook = %q, want %q", have.PostTweetHook, want.PostTweetHook)
			}

			if fmt.Sprint(have.Following) != fmt.Sprint(want.Following) {
				t.Errorf("Following = %#v, want %#v", have.Following, want.Following)
			}
		})
	}
}

/*
func NewConfig(path string) (config, error) {
	// setup default values
	cfg := config{
		CheckFollowing:         true,
		UseCache:               true,
		LimitTimeline:          20,
		TimelineUpdateInterval: 10,
		Timeout:                5.0,
		Following:              make(map[string]string),
	}

	// use default path if non set
	if path == "" {
		dir, err := os.UserConfigDir()
		if err != nil {
			return config{}, err
		}

		conf = dir + "/twtxt/config"
	}

	// load configuration file
	file, err := ini.Load(path)
	if err != nil {
		return config{}, err
	}

	// get twtxt config section, use defaults if not found
	if twtxt := file.Section("twtxt"); twtxt != nil {
		setConfigString(twtxt, "nick", &cfg.Nick)
		setConfigString(twtxt, "twtfile", &cfg.Twtfile)
		setConfigString(twtxt, "twturl", &cfg.Twturl)

		if err := setConfigBool(twtxt, "check_following", &cfg.CheckFollowing); err != nil {
			return config{}, err
		}

		if err := setConfigBool(twtxt, "use_pager", &cfg.UsePager); err != nil {
			return config{}, err
		}

		if err := setConfigBool(twtxt, "use_cache", &cfg.UseCache); err != nil {
			return config{}, err
		}

		if err := setConfigBool(twtxt, "porcelain", &cfg.Porcelain); err != nil {
			return config{}, err
		}

		if err := setConfigBool(twtxt, "disclose_identity", &cfg.DiscloseIdentity); err != nil {
			return config{}, err
		}

		if err := setConfigInt(twtxt, "character_limit", &cfg.CharacterLimit); err != nil {
			return config{}, err
		}

		if err := setConfigInt(twtxt, "character_warning", &cfg.CharacterWarning); err != nil {
			return config{}, err
		}

		if err := setConfigInt(twtxt, "limit_timeline", &cfg.LimitTimeline); err != nil {
			return config{}, err
		}

		if err := setConfigInt(twtxt, "timeline_update_interval", &cfg.TimelineUpdateInterval); err != nil {
			return config{}, err
		}

		if err := setConfigFloat64(twtxt, "timeout", &cfg.Timeout); err != nil {
			return config{}, err
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
			return config{}, fmt.Errorf("Invalid value for 'sorting': %q", sorting)
		}

		if err := setConfigBool(twtxt, "use_abs_time", &cfg.UseAbsoluteTime); err != nil {
			return config{}, err
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
	return cfg, nil
}
*/
