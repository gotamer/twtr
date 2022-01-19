package twtxt

// Feed is a collection of Tweets, these are not necessarily from the same
// source, or in any particular order. However, a feed can be sorted by the
// timestamp (in a timezone aware manner).
type Feed []Tweet

// Len reports the number of Tweets in the Feed.
func (feed Feed) Len() int { return len(feed) }

// Less reports if the Tweet at i was posted before the Tweet at j.
func (feed Feed) Less(i, j int) bool { return feed[i].Time.Before(feed[j].Time) }

// Swap exchanges the Tweet at i with the Tweet at j.
func (feed Feed) Swap(i, j int) { feed[i], feed[j] = feed[j], feed[i] }
