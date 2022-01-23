package twtxt

import "io"

// Tweets are a collection of Tweet instances, these are not necessarily from
// the same source, or in any particular order. However, a feed can be sorted by
// the timestamp (in a timezone aware manner).
type Tweets []*Tweet

// ParseTweets reads Tweets from a file or other source.
//
// See the twtxt file format specification for more information:
// https://twtxt.readthedocs.io/en/latest/user/twtxtfile.html
func ParseTweets(reader io.Reader) (Tweets, error) {
	// TODO: read from reader
	lines := []string{}

	// make room for enough tweets
	tweets := make(Tweets, len(lines))

	return tweets, nil
}

// Len reports the number of Tweets.
func (twts Tweets) Len() int { return len(twts) }

// Less reports if the Tweet at i was posted before the Tweet at j.
func (twts Tweets) Less(i, j int) bool { return twts[i].Before(twts[j]) }

// Swap exchanges the Tweet at i with the Tweet at j.
func (twts Tweets) Swap(i, j int) { twts[i], twts[j] = twts[j], twts[i] }
