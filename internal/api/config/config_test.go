package config_test

import (
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
	file := "_test/twtxt.config"
	_, err := config.NewConfig(file)
	if err != nil {
		t.Fatalf("could not open file '%s': %q", file, err)
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
