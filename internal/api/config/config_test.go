package config_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"internal/api/config"

	"gopkg.in/ini.v1"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		file          string
		want          config.Config
		expectedError func(error) bool
	}{
		{ // All fields with example value
			file: "_test/example1.ini",
			want: config.Config{
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
		{ // No following section
			file: "_test/example2.ini",
			want: config.Config{
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
				Following:              make(map[string]string),
			},
		},
		{ // No twtxt section (default values) + following section
			file: "_test/example3.ini",
			want: config.Config{
				CheckFollowing:         true,
				UseCache:               true,
				LimitTimeline:          20,
				TimelineUpdateInterval: 10,
				Timeout:                5.0,
				Following: map[string]string{
					"alice": "https://example.org/alice.txt",
					"bob":   "https://example.org/bob.txt",
				},
			},
		},
		{ // empty file (default values)
			file: "_test/example4.ini",
			want: config.Config{
				CheckFollowing:         true,
				UseCache:               true,
				LimitTimeline:          20,
				TimelineUpdateInterval: 10,
				Timeout:                5.0,
				Following:              make(map[string]string),
			},
		},
		{ // valid INI file but with no relevant sections (default values)
			file: "_test/example5.ini",
			want: config.Config{
				CheckFollowing:         true,
				UseCache:               true,
				LimitTimeline:          20,
				TimelineUpdateInterval: 10,
				Timeout:                5.0,
				Following:              make(map[string]string),
			},
		},
		{ // valid INI file with relevant sections, but no real values (default values)
			file: "_test/example6.ini",
			want: config.Config{
				CheckFollowing:         true,
				UseCache:               true,
				LimitTimeline:          20,
				TimelineUpdateInterval: 10,
				Timeout:                5.0,
				Following: map[string]string{
					"meaningOfLife":        "42",
					"notActuallyANickname": "Not actually a url",
				},
			},
		},
		{ // invalid INI file (parse error)
			file: "_test/example7.ini",
			expectedError: func(err error) bool {
				return ini.IsErrDelimiterNotFound(err)
			},
		},
		{ // valid INI file but wrong types for values (parse error)
			file: "_test/example8.ini",
			expectedError: func(err error) bool {
				return strings.Contains(err.Error(), "invalid syntax")
			},
		},
	}

	for _, test := range tests {
		test := test
		name := test.file
		want := test.want

		t.Run(name, func(t *testing.T) {
			file, err := os.Open(test.file)
			defer file.Close()
			if err != nil {
				t.Fatalf("could not open file '%s': %q", test.file, err)
			}

			have, err := config.New(file)
			if err != nil {
				if test.expectedError != nil && test.expectedError(err) {
					t.Logf("expected error: %q", err)
					t.Skip()
				} else {
					t.Fatalf("unexpected error: %q", err)
				}
			} else if test.expectedError != nil {
				t.Fatal("an error was expected but none was received")
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
