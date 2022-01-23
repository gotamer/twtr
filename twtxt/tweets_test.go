package twtxt

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func copyFeed(twts Tweets) Tweets {
	tmp := make([]*Tweet, len(twts))

	copy(tmp, twts)

	return Tweets(tmp)
}

func TestParseTweets(t *testing.T) {
	tests := []struct {
		name   string
		source io.Reader
		tweets Tweets
		err    error
	}{
		{
			name:   "Nil",
			source: nil,
			tweets: Tweets{},
		},
		{
			name:   "Empty",
			source: strings.NewReader(""),
			tweets: Tweets{},
		},
		{
			name:   "SingleTweet",
			source: strings.NewReader("2022-01-19T14:14:00+13:00\tThis post contains newlines\\n\\n\\n"),
			tweets: Tweets{
				&Tweet{
					time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
					post: "This post contains newlines\n\n\n",
				},
			},
		},
		{
			name: "MultipleTweets",
			source: strings.NewReader(strings.Join([]string{
				"2022-01-19T14:11:00+13:00\tThis post contains tabs\\t\\t\\t",
				"2016-02-03T23:05:00+01:00\t@<example http://example.org/twtxt.txt> welcome to twtxt!",
				"2022-01-19T14:14:00+13:00\tThis post contains newlines\\n\\n\\n",
				"2016-02-01T11:00:00+01:00\tThis is just another example.",
				"2015-12-12T12:00:00+01:00\tFiat lux!",
				"2016-02-04T13:30:00+01:00\tYou can really go crazy here! ┐(ﾟ∀ﾟ)┌",
			}, "\n")),
			tweets: Tweets{
				&Tweet{
					time: time.Date(2022, 1, 19, 14, 11, 0, 0, loc(+13)),
					post: "This post contains tabs\t\t\t",
				},
				&Tweet{
					time: time.Date(2016, 2, 3, 23, 5, 0, 0, loc(+1)),
					post: "@<example http://example.org/twtxt.txt> welcome to twtxt!",
				},
				&Tweet{
					time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
					post: "This post contains newlines\n\n\n",
				},
				&Tweet{
					time: time.Date(2016, 2, 1, 11, 0, 0, 0, loc(+1)),
					post: "This is just another example.",
				},
				&Tweet{
					time: time.Date(2015, 12, 12, 12, 0, 0, 0, loc(+1)),
					post: "Fiat lux!",
				},
				&Tweet{
					time: time.Date(2016, 2, 4, 13, 30, 0, 0, loc(+1)),
					post: "You can really go crazy here! ┐(ﾟ∀ﾟ)┌",
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			tweets, err := ParseTweets(test.source)

			if err != test.err {
				t.Fatalf("have %q, want %q", err, test.err)
			}

			if !cmp.Equal(tweets, test.tweets) {
				t.Errorf("diff:\n%s", cmp.Diff(tweets, test.tweets))
			}
		})
	}
}

func TestTweets(t *testing.T) {
	tests := []struct {
		name string
		twts Tweets
	}{
		{
			name: "Nil",
			twts: nil,
		},
		{
			name: "Empty",
			twts: Tweets{
				&Tweet{
					time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
					post: "This post contains newlines\n\n\n",
				},
			},
		},
		{
			name: "SingleTweet",
			twts: Tweets{
				&Tweet{
					time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
					post: "This post contains newlines\n\n\n",
				},
			},
		},
		{
			name: "MultipleTweets",
			twts: Tweets{
				&Tweet{
					time: time.Date(2022, 1, 19, 14, 11, 0, 0, loc(+13)),
					post: "This post contains tabs\t\t\t",
				},
				&Tweet{
					time: time.Date(2016, 2, 3, 23, 5, 0, 0, loc(+1)),
					post: "@<example http://example.org/twtxt.txt> welcome to twtxt!",
				},
				&Tweet{
					time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
					post: "This post contains newlines\n\n\n",
				},
				&Tweet{
					time: time.Date(2016, 2, 1, 11, 0, 0, 0, loc(+1)),
					post: "This is just another example.",
				},
				&Tweet{
					time: time.Date(2015, 12, 12, 12, 0, 0, 0, loc(+1)),
					post: "Fiat lux!",
				},
				&Tweet{
					time: time.Date(2016, 2, 4, 13, 30, 0, 0, loc(+1)),
					post: "You can really go crazy here! ┐(ﾟ∀ﾟ)┌",
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Run("Len", func(t *testing.T) {
				if have, want := test.twts.Len(), len(test.twts); have != want {
					t.Errorf("have %d, want %d", have, want)
				}
			})

			for i := range test.twts {
				for j := range test.twts {
					t.Run(fmt.Sprintf("Less(%d,%d)", i, j), func(t *testing.T) {
						twts := copyFeed(test.twts)
						have := twts.Less(i, j)
						want := twts[i].Time().Before(twts[j].Time())

						if have != want {
							t.Errorf("have %t, want %t", have, want)
						}
					})

					t.Run(fmt.Sprintf("Swap(%d,%d)", i, j), func(t *testing.T) {
						twts := copyFeed(test.twts)
						orig := copyFeed(test.twts)

						twts.Swap(i, j)

						if have, want := orig[j], twts[i]; have != want {
							t.Errorf("have %#v, want %#v", have, want)
						}

						if have, want := orig[i], twts[j]; have != want {
							t.Errorf("have %#v, want %#v", have, want)
						}
					})
				}
			}

			t.Run("Sort", func(t *testing.T) {
				twts := copyFeed(test.twts)

				sort.Sort(twts)

				if !sort.IsSorted(twts) {
					t.Error("sort failed")
				}

				t.Run("Reverse", func(t *testing.T) {
					twts := copyFeed(test.twts)

					sort.Sort(sort.Reverse(twts))

					if !sort.IsSorted(sort.Reverse(twts)) {
						t.Error("sort failed")
					}
				})
			})
		})
	}
}
