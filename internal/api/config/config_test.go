package config_test

import (
	"fmt"
	"testing"

	"internal/api/config"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		file string
		want config.Config
	}{
		{
			"_test/example1.ini",
			config.Config{
				Nick:                   "buckket",
				Twtfile:                "~/twtxt.txt",
				Twturl:                 "http://example.org/twtxt.txt",
				CheckFollowing:         true,
				UsePager:               false,
				UseCache:               true,
				Porcelain:              false,
				DiscloseIdentity:       false,
				CharacterLimit:         140,
				CharacterWarning:       140,
				LimitTimeline:          20,
				TimelineUpdateInterval: 10,
				Timeout:                5.0,
				SortAscending:          false,
				PreTweetHook:           "scp buckket@example.org:~/public_html/twtxt.txt {twtfile}",
				PostTweetHook:          "scp {twtfile} buckket@example.org:~/public_html/twtxt.txt",
				Following: map[string]string{
					"alice": "https://example.org/alice.txt",
					"bob":   "https://example.org/bob.txt",
				},
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
