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
		feed Tweets
	}{
		{
			name: "Nil",
			feed: nil,
		},
		{
			name: "Empty",
			feed: Tweets{
				&Tweet{
					time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
					post: "This post contains newlines\n\n\n",
				},
			},
		},
		{
			name: "SingleTweet",
			feed: Tweets{
				&Tweet{
					time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
					post: "This post contains newlines\n\n\n",
				},
			},
		},
		{
			name: "MultipleTweets",
			feed: Tweets{
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
				if have, want := test.feed.Len(), len(test.feed); have != want {
					t.Errorf("have %d, want %d", have, want)
				}
			})

			for i := range test.feed {
				for j := range test.feed {
					t.Run(fmt.Sprintf("Less(%d,%d)", i, j), func(t *testing.T) {
						feed := copyFeed(test.feed)
						have := feed.Less(i, j)
						want := feed[i].Time().Before(feed[j].Time())

						if have != want {
							t.Errorf("have %t, want %t", have, want)
						}
					})

					t.Run(fmt.Sprintf("Swap(%d,%d)", i, j), func(t *testing.T) {
						feed := copyFeed(test.feed)
						orig := copyFeed(test.feed)

						feed.Swap(i, j)

						if have, want := orig[j], feed[i]; have != want {
							t.Errorf("have %#v, want %#v", have, want)
						}

						if have, want := orig[i], feed[j]; have != want {
							t.Errorf("have %#v, want %#v", have, want)
						}
					})
				}
			}

			t.Run("Sort", func(t *testing.T) {
				feed := copyFeed(test.feed)

				sort.Sort(feed)

				if !sort.IsSorted(feed) {
					t.Error("sort failed")
				}

				t.Run("Reverse", func(t *testing.T) {
					feed := copyFeed(test.feed)

					sort.Sort(sort.Reverse(feed))

					if !sort.IsSorted(sort.Reverse(feed)) {
						t.Error("sort failed")
					}
				})
			})
		})
	}
}
