package config_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"duriny.envs.sh/twtr/twtxt/config"
	"github.com/google/go-cmp/cmp"

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
			want: config.Config{
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

			if cmp.Equal(have, want) {
				t.Errorf("diff:\n%s", cmp.Diff(have, want))
			}
		})
	}
}

func TestConfigWriteTo(t *testing.T) {
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
check_following          = true
use_pager                = false
use_cache                = true
porcelain                = false
disclose_identity        = false
character_limit          = 140
character_warning        = 140
limit_timeline           = 20
timeline_update_interval = 10
timeout                  = 5.0
use_abs_time             = false
pre_tweet_hook           = scp buckket@example.org:~/public_html/twtxt.txt {twtfile}
post_tweet_hook          = scp {twtfile} buckket@example.org:~/public_html/twtxt.txt
sorting                  = descending

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
check_following          = true
use_pager                = false
use_cache                = true
porcelain                = false
disclose_identity        = false
character_limit          = 140
character_warning        = 140
limit_timeline           = 20
timeline_update_interval = 10
timeout                  = 5.0
use_abs_time             = false
pre_tweet_hook           = scp buckket@example.org:~/public_html/twtxt.txt {twtfile}
post_tweet_hook          = scp {twtfile} buckket@example.org:~/public_html/twtxt.txt
sorting                  = descending

`,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			var have bytes.Buffer

			n, err := test.from.WriteTo(&have)

			if n != int64(len(test.want)) {
				t.Fatalf("n = %d, want %d", n, int64(len(test.want)))
			}

			if err != nil {
				t.Fatalf("err = %q, want nil", err)
			}

			if have, want := have.String(), test.want; have != want {
				t.Errorf("diff:\n%s", cmp.Diff(have, want))
			}

			t.Run("Reversibly", func(t *testing.T) {
				want := test.from
				have, err := config.New(strings.NewReader(test.want))
				if err != nil {
					t.Fatalf("unexpected error: %q", err)
				}

				if cmp.Equal(have, want) {
					t.Errorf("diff:\n%s", cmp.Diff(have, want))
				}
			})
		})
	}
}
