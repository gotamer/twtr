package twtxt

import (
	"fmt"
	"testing"
	"time"
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
			tweet := NewTweet(test)

			if have, want := tweet.post, test; have != want {
				t.Errorf("have %q, want %q", have, want)
			}
		})
	}
}

func TestTweet(t *testing.T) {
	tests := []struct {
		String  string
		Time time.Time
		Post string
		twt *Tweet
	}{
		{
			String: "2022-01-19T14:14:00+13:00\tThis post contains newlines\\n\\n\\n",
			Time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
			Post: "This post contains newlines\n\n\n",
			twt: &Tweet{
				time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
				post: "This post contains newlines\n\n\n",
			},
		},
		{
			String: "2022-01-19T14:11:00+13:00\tThis post contains tabs\\t\\t\\t",
			Time: time.Date(2022, 1, 19, 14, 11, 0, 0, loc(+13)),
			Post: "This post contains tabs\t\t\t",
			twt: &Tweet{
				time: time.Date(2022, 1, 19, 14, 11, 0, 0, loc(+13)),
				post: "This post contains tabs\t\t\t",
			},
		},
		{
			String: "2016-02-04T13:30:00+01:00\tYou can really go crazy here! ┐(ﾟ∀ﾟ)┌",
			Time: time.Date(2016, 2, 4, 13, 30, 0, 0, loc(+1)),
			Post: "You can really go crazy here! ┐(ﾟ∀ﾟ)┌",
			twt: &Tweet{
				time: time.Date(2016, 2, 4, 13, 30, 0, 0, loc(+1)),
				post: "You can really go crazy here! ┐(ﾟ∀ﾟ)┌",
			},
		},
		{
			String: "2016-02-03T23:05:00+01:00\t@<example http://example.org/twtxt.txt> welcome to twtxt!",
			Time: time.Date(2016, 2, 3, 23, 5, 0, 0, loc(+1)),
			Post: "@<example http://example.org/twtxt.txt> welcome to twtxt!",
			twt: &Tweet{
				time: time.Date(2016, 2, 3, 23, 5, 0, 0, loc(+1)),
				post: "@<example http://example.org/twtxt.txt> welcome to twtxt!",
			},
		},
		{
			String: "2016-02-01T11:00:00+01:00\tThis is just another example.",
			Time: time.Date(2016, 2, 1, 11, 0, 0, 0, loc(+1)),
			Post: "This is just another example.",
			twt: &Tweet{
				time: time.Date(2016, 2, 1, 11, 0, 0, 0, loc(+1)),
				post: "This is just another example.",
			},
		},
		{
			String: "2015-12-12T12:00:00+01:00\tFiat lux!",
			Time: time.Date(2015, 12, 12, 12, 0, 0, 0, loc(+1)),
			Post: "Fiat lux!",
			twt: &Tweet{
				time: time.Date(2015, 12, 12, 12, 0, 0, 0, loc(+1)),
				post: "Fiat lux!",
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.String, func(t *testing.T) {
			t.Run("String()", func(t *testing.T) {
				if have, want := test.twt.String(), test.String; have != want {
					t.Errorf("\nhave: %q\nwant: %q", have, want)
				}
			})

			t.Run("Time()", func(t *testing.T) {
				if have, want := test.twt.Time(), test.Time; have.Unix() != want.Unix() {
					t.Errorf("\nhave: %s\nwant: %s", have, want)
				}
			})

			t.Run("Post()", func(t *testing.T) {
				if have, want := test.twt.Post(), test.Post; have != want {
					t.Errorf("\nhave: %s\nwant: %s", have, want)
				}
			})
		})
	}
}
