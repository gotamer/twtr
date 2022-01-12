package config_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"duriny.envs.sh/twtr/internal/api/config"

	"gopkg.in/ini.v1"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name          string
		source        io.Reader
		want          config.Config
		expectedError func(error) bool
	}{
		{
			name: "AllFieldsSetWithExampleValue",
			source: strings.NewReader(`
[twtxt]
nick = buckket
twtfile = ~/twtxt.txt
twturl = http://example.org/twtxt.txt
check_following = True
use_pager = False
use_cache = True
porcelain = False
disclose_identity = False
character_limit = 140
character_warning = 140
limit_timeline = 20
timeline_update_interval = 10
timeout = 5.0
sorting = descending
use_abs_time = false
pre_tweet_hook = "scp buckket@example.org:~/public_html/twtxt.txt {twtfile}"
post_tweet_hook = "scp {twtfile} buckket@example.org:~/public_html/twtxt.txt"

[following]
bob = https://example.org/bob.txt
alice = https://example.org/alice.txt
`),
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
				UseAbsoluteTime:        false,
				PreTweetHook:           "scp buckket@example.org:~/public_html/twtxt.txt {twtfile}",
				PostTweetHook:          "scp {twtfile} buckket@example.org:~/public_html/twtxt.txt",
				Following: map[string]string{
					"alice": "https://example.org/alice.txt",
					"bob":   "https://example.org/bob.txt",
				},
			},
		},
		{
			name: "NoFollowingSection",
			source: strings.NewReader(`
[twtxt]
nick = buckket
twtfile = ~/twtxt.txt
twturl = http://example.org/twtxt.txt
check_following = True
use_pager = False
use_cache = True
porcelain = False
disclose_identity = False
character_limit = 140
character_warning = 140
limit_timeline = 20
timeline_update_interval = 10
timeout = 5.0
sorting = descending
use_abs_time = false
pre_tweet_hook = "scp buckket@example.org:~/public_html/twtxt.txt {twtfile}"
post_tweet_hook = "scp {twtfile} buckket@example.org:~/public_html/twtxt.txt"
`),
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
				UseAbsoluteTime:        false,
				PreTweetHook:           "scp buckket@example.org:~/public_html/twtxt.txt {twtfile}",
				PostTweetHook:          "scp {twtfile} buckket@example.org:~/public_html/twtxt.txt",
				Following:              make(map[string]string),
			},
		},
		{
			name: "NoTwtxtSection",
			source: strings.NewReader(`
[following]
bob = https://example.org/bob.txt
alice = https://example.org/alice.txt
`),
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
		{
			name:   "EmptyFile",
			source: strings.NewReader(""),
			want: config.Config{
				CheckFollowing:         true,
				UseCache:               true,
				LimitTimeline:          20,
				TimelineUpdateInterval: 10,
				Timeout:                5.0,
				Following:              make(map[string]string),
			},
		},
		{
			name: "ValidIniFileWithUnrelatedSections",
			source: strings.NewReader(`
[Section A]
magic = on
deepMagic = off

[Section B]
samuraiWarrior = "foolish"
`),
			want: config.Config{
				CheckFollowing:         true,
				UseCache:               true,
				LimitTimeline:          20,
				TimelineUpdateInterval: 10,
				Timeout:                5.0,
				Following:              make(map[string]string),
			},
		},
		{
			name: "ValidFileWithCorrectSectionsButUnrelatedSections",
			source: strings.NewReader(`
[twtxt]
magic = enable
deepMagic = DEAR_GOD_TURN_IT_OFF

[following]
notActuallyANickname = "Not actually a url"
meaningOfLife = 42
`),
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
		{
			name: "InvalidIniFile",
			source: strings.NewReader(`
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean elementum nisi
ac nisi vulputate, efficitur varius risus ultricies. Sed fringilla hendrerit
augue, a porta nunc placerat a. Donec ac nisi scelerisque, gravida diam iaculis,
molestie libero. Nam congue tristique ipsum, nec interdum metus maximus id.
Aenean nec sodales lectus, et eleifend nibh. Nulla sit amet pellentesque elit.
Vestibulum ultricies porta nibh eu convallis.

Aliquam ut dui vitae nisi faucibus fermentum faucibus vel tellus. Proin aliquam
nisi quis metus ornare, ullamcorper ultrices erat euismod. Nulla semper
consequat ipsum eget sagittis. Vestibulum egestas pellentesque nisl, eget
gravida magna finibus sit amet. Morbi tristique ac magna at rutrum. Nullam
efficitur sollicitudin vestibulum. Nullam non justo massa. Vivamus maximus, enim
in sodales euismod, turpis nunc egestas arcu, nec bibendum augue dolor ut augue.
Donec malesuada ut arcu sed mattis. Nulla consectetur felis at varius finibus.
Vivamus porta nibh justo, eu ultrices ante ullamcorper eget.

Phasellus maximus magna sed neque viverra, maximus rhoncus arcu accumsan. Nam
convallis metus id nunc aliquet, sed placerat ante placerat. Maecenas quis
sapien hendrerit, mattis turpis ut, convallis mauris. Lorem ipsum dolor sit
amet, consectetur adipiscing elit. Suspendisse mattis ultricies elit, commodo
lacinia leo. Aenean at dolor risus. Donec vel imperdiet ante. Praesent pharetra
ex tincidunt sapien interdum, ac accumsan lacus blandit. Donec risus augue,
ornare sed elit sed, porta varius ex. Proin sit amet eros at neque condimentum
laoreet ut vitae nisi.

In dignissim, sem sed viverra gravida, velit turpis sagittis lacus, non vehicula
magna mi quis tellus. Vestibulum ullamcorper justo eu est lacinia, id porttitor
ipsum tempor. Orci varius natoque penatibus et magnis dis parturient montes,
nascetur ridiculus mus. Pellentesque tristique nisi lacus, ut cursus dui
accumsan vel. Cras semper commodo magna, eu blandit mauris euismod id. Donec
justo mi, tempor sed tellus facilisis, lacinia imperdiet eros. Etiam in placerat
erat. Pellentesque nibh nibh, commodo a dolor quis, tempor efficitur orci.

Interdum et malesuada fames ac ante ipsum primis in faucibus. Sed luctus auctor
velit, in euismod magna interdum rutrum. Etiam sit amet mi et urna elementum
condimentum. Proin pellentesque lectus ac malesuada mollis. Nulla ornare
dignissim felis, vel lobortis dui malesuada sed. Sed tempor luctus libero.
Aenean convallis nec enim placerat lobortis. Praesent eleifend nisl tortor, a
imperdiet urna dictum in. Cras pellentesque eleifend sapien at elementum. Proin
aliquam placerat quam. Cras vel vestibulum mauris, nec viverra purus. Curabitur
dolor neque, commodo sed luctus ac, lacinia non mi.
`),
			expectedError: func(err error) bool {
				return ini.IsErrDelimiterNotFound(err)
			},
		},
		{
			name: "ValidIniFileWithWrongTypes",
			source: strings.NewReader(`
[twtxt]
nick = buckket
twtfile = ~/twtxt.txt
twturl = http://example.org/twtxt.txt
check_following = True
use_pager = "not a boolean string"
use_cache = "not a boolean string"
porcelain = False
disclose_identity = False
character_limit = 140.123
character_warning = 140.123
limit_timeline = 20
timeline_update_interval = 10
timeout = 5.0
sorting = neither_descending_nor_ascending
use_abs_time = "This is not a pipe, I mean bool"
pre_tweet_hook = "scp buckket@example.org:~/public_html/twtxt.txt {twtfile}"
post_tweet_hook = "scp {twtfile} buckket@example.org:~/public_html/twtxt.txt"

[following]
bob = https://example.org/bob.txt
alice = https://example.org/alice.txt
`),
			expectedError: func(err error) bool {
				return strings.Contains(err.Error(), "invalid syntax")
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			want := test.want

			have, err := config.New(test.source)
			if err != nil {
				if test.expectedError != nil && test.expectedError(err) {
					// t.Logf("expected error: %q", err)
					return
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

func TestConfigSave(t *testing.T) {
	tests := []struct {
		name string
		from config.Config
		want string
	}{
		{
			name: "AllFieldsSetWithExampleValue",
			from: config.Config{
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
				UseAbsoluteTime:        false,
				PreTweetHook:           "scp buckket@example.org:~/public_html/twtxt.txt {twtfile}",
				PostTweetHook:          "scp {twtfile} buckket@example.org:~/public_html/twtxt.txt",
				Following: map[string]string{
					"alice": "https://example.org/alice.txt",
					"bob":   "https://example.org/bob.txt",
				},
			},
			want: `[twtxt]
nick                     = buckket
twtfile                  = ~/twtxt.txt
twturl                   = http://example.org/twtxt.txt
check_following          = True
use_pager                = False
use_cache                = True
porcelain                = False
disclose_identity        = False
character_limit          = 140
character_warning        = 140
limit_timeline           = 20
timeline_update_interval = 10
timeout                  = 5.0
sorting                  = descending
use_abs_time             = False
pre_tweet_hook           = scp buckket@example.org:~/public_html/twtxt.txt {twtfile}
post_tweet_hook          = scp {twtfile} buckket@example.org:~/public_html/twtxt.txt

[following]
alice = https://example.org/alice.txt
bob   = https://example.org/bob.txt

`,
		},
		{
			name: "NoFollowingSection",
			from: config.Config{
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
				UseAbsoluteTime:        false,
				PreTweetHook:           "scp buckket@example.org:~/public_html/twtxt.txt {twtfile}",
				PostTweetHook:          "scp {twtfile} buckket@example.org:~/public_html/twtxt.txt",
				Following:              make(map[string]string),
			},
			want: `[twtxt]
nick                     = buckket
twtfile                  = ~/twtxt.txt
twturl                   = http://example.org/twtxt.txt
check_following          = True
use_pager                = False
use_cache                = True
porcelain                = False
disclose_identity        = False
character_limit          = 140
character_warning        = 140
limit_timeline           = 20
timeline_update_interval = 10
timeout                  = 5.0
sorting                  = descending
use_abs_time             = False
pre_tweet_hook           = scp buckket@example.org:~/public_html/twtxt.txt {twtfile}
post_tweet_hook          = scp {twtfile} buckket@example.org:~/public_html/twtxt.txt
`,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			var have bytes.Buffer

			if err := test.from.Save(&have); err != nil {
				t.Fatalf("err = %q, want nil", err)
			}

			haveLines := strings.Split(have.String(), "\n")
			wantLines := strings.Split(test.want, "\n")

			if len(haveLines) < len(wantLines) {
				for i, have := range haveLines {
					have := have
					want := wantLines[i]

					if have != want {
						t.Errorf("line %d:\nhave %q\nwant %q\n", i+1, have, want)
					}
				}

				t.Errorf("line %d:\nhave EOF\nwant %q", len(haveLines)+1, wantLines[len(haveLines)])
			} else if len(haveLines) > len(wantLines) {
				for i, want := range wantLines {
					have := haveLines[1]
					want := want

					if have != want {
						t.Errorf("line %d:\nhave %q\nwant %q\n", i+1, have, want)
					}
				}

				t.Errorf("line %d:\nhave %q\nwant EOF", len(wantLines)+1, haveLines[len(wantLines)])
			} else {
				for i, have := range haveLines {
					have := have
					want := wantLines[i]

					if have != want {
						t.Errorf("line %d:\nhave %q\nwant %q\n", i+1, have, want)
					}
				}
			}

			t.Run("Reversibly", func(t *testing.T) {
				want := test.from

				have, err := config.New(strings.NewReader(test.want))
				if err != nil {
					t.Fatalf("unexpected error: %q", err)
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
		})
	}
}
