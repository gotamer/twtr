package twtxt_test

import (
	"fmt"
	"testing"
	"time"

	"duriny.envs.sh/twtr/twtxt"
)

// loc is a helper to create arbitrary time.Locations.
func loc(offset int) *time.Location {
	return time.FixedZone(fmt.Sprintf("UTC%+d", offset), offset*60*60)
}

func TestNewTweet(t *testing.T) {
	tests := []string{
		"Some message or status update",
	}

	for _, test := range tests {
		test := test

		t.Run(test, func(t *testing.T) {
			tweet := twtxt.NewTweet(test)

			if have, want := tweet.Post, test; have != want {
				t.Errorf("have %q, want %q", have, want)
			}
		})
	}
}

func TestTweet(t *testing.T) {
	tests := []struct {
		want  string
		tweet twtxt.Tweet
	}{
		{
			want: "2022-01-19T14:14:00+13:00\tThis post contains newlines\\n\\n\\n",
			tweet: twtxt.Tweet{
				Time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
				Post: "This post contains newlines\n\n\n",
			},
		},
		{
			want: "2022-01-19T14:11:00+13:00\tThis post contains tabs\\t\\t\\t",
			tweet: twtxt.Tweet{
				Time: time.Date(2022, 1, 19, 14, 11, 0, 0, loc(+13)),
				Post: "This post contains tabs\t\t\t",
			},
		},
		{
			want: "2016-02-04T13:30:00+01:00\tYou can really go crazy here! ┐(ﾟ∀ﾟ)┌",
			tweet: twtxt.Tweet{
				Time: time.Date(2016, 2, 4, 13, 30, 0, 0, loc(+1)),
				Post: "You can really go crazy here! ┐(ﾟ∀ﾟ)┌",
			},
		},
		{
			want: "2016-02-03T23:05:00+01:00\t@<example http://example.org/twtxt.txt> welcome to twtxt!",
			tweet: twtxt.Tweet{
				Time: time.Date(2016, 2, 3, 23, 5, 0, 0, loc(+1)),
				Post: "@<example http://example.org/twtxt.txt> welcome to twtxt!",
			},
		},
		{
			want: "2016-02-01T11:00:00+01:00\tThis is just another example.",
			tweet: twtxt.Tweet{
				Time: time.Date(2016, 2, 1, 11, 0, 0, 0, loc(+1)),
				Post: "This is just another example.",
			},
		},
		{
			want: "2015-12-12T12:00:00+01:00\tFiat lux!",
			tweet: twtxt.Tweet{
				Time: time.Date(2015, 12, 12, 12, 0, 0, 0, loc(+1)),
				Post: "Fiat lux!",
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.want, func(t *testing.T) {
			t.Run("String()", func(t *testing.T) {
				if have, want := test.tweet.String(), test.want; have != want {
					t.Errorf("\nhave: %q\nwant: %q", have, want)
				}
			})
		})
	}
}
