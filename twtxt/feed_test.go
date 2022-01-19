package twtxt_test

import (
	"fmt"
	"sort"
	"testing"
	"time"

	"duriny.envs.sh/twtr/twtxt"
)

func copyFeed(feed twtxt.Feed) twtxt.Feed {
	tmp := make([]twtxt.Tweet, len(feed))

	copy(tmp, feed)

	return twtxt.Feed(tmp)
}

func TestFeed(t *testing.T) {
	tests := []struct {
		name string
		feed twtxt.Feed
	}{
		{
			name: "NilFeed",
			feed: nil,
		},
		{
			name: "FeedWithNoTweets",
			feed: twtxt.Feed{
				twtxt.Tweet{
					Time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
					Post: "This post contains newlines\n\n\n",
				},
			},
		},
		{
			name: "FeedWithSingleTweet",
			feed: twtxt.Feed{
				twtxt.Tweet{
					Time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
					Post: "This post contains newlines\n\n\n",
				},
			},
		},
		{
			name: "FeedWithMultipleTweets",
			feed: twtxt.Feed{
				twtxt.Tweet{
					Time: time.Date(2022, 1, 19, 14, 11, 0, 0, loc(+13)),
					Post: "This post contains tabs\t\t\t",
				},
				twtxt.Tweet{
					Time: time.Date(2016, 2, 3, 23, 5, 0, 0, loc(+1)),
					Post: "@<example http://example.org/twtxt.txt> welcome to twtxt!",
				},
				twtxt.Tweet{
					Time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
					Post: "This post contains newlines\n\n\n",
				},
				twtxt.Tweet{
					Time: time.Date(2016, 2, 1, 11, 0, 0, 0, loc(+1)),
					Post: "This is just another example.",
				},
				twtxt.Tweet{
					Time: time.Date(2015, 12, 12, 12, 0, 0, 0, loc(+1)),
					Post: "Fiat lux!",
				},
				twtxt.Tweet{
					Time: time.Date(2016, 2, 4, 13, 30, 0, 0, loc(+1)),
					Post: "You can really go crazy here! ┐(ﾟ∀ﾟ)┌",
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
						want := feed[i].Time.Before(feed[j].Time)

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
