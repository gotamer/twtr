package twtxt

// Tweets are a collection of Tweet instances, these are not necessarily from
// the same source, or in any particular order. However, a feed can be sorted by
// the timestamp (in a timezone aware manner).
type Tweets []*Tweet

// Len reports the number of Tweets.
func (twts Tweets) Len() int { return len(twts) }

// Less reports if the Tweet at i was posted before the Tweet at j.
func (twts Tweets) Less(i, j int) bool { return twts[i].Time().Before(twts[j].Time()) }

// Swap exchanges the Tweet at i with the Tweet at j.
func (twts Tweets) Swap(i, j int) { twts[i], twts[j] = twts[j], twts[i] }
