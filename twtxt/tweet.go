package twtxt

import (
	"strings"
	"time"
)

// Tweet represents a single twtxt post, containing the timestamp and the main
// body of the post.
type Tweet struct {
	time time.Time
	post string
}

// NewTweet creates a new Tweet instance with the given post message and the
// current local time.
func NewTweet(post string) *Tweet {
	twt := &Tweet{post: post}

	// trigger time init
	twt.Time()

	return twt
}

// Time gets the time that the Tweet was posted.
func (twt *Tweet) Time() time.Time {
	if twt.time.IsZero() {
		twt.time = time.Now()
	}

	return twt.time
}

// Post gets the posted message of the Tweet.
func (twt *Tweet) Post() string {
	return twt.post
}

// Before determines if one Tweet was posted before the other.
func (twt *Tweet) Before(other *Tweet) bool {
	return twt.Time().Before(other.Time())
}

// After determines if one Tweet was posted after the other.
func (twt *Tweet) After(other *Tweet) bool {
	return twt.Time().After(other.Time())
}

// String formats the Tweet as an entry into a twtxt.txt file, returns the
// timestamp followed by a tab character and the post message. Any tabs or new
// line characters are escaped to prevent invalid formatting of the twtxt.txt
// file.
//
//     <yyyy>-<mm>-<dd>T<HH>:<MM>:<SS><+/-><XX>:<ZZ>\t<POST>
//
// See the format specification for more details on the file format:
// https://twtxt.readthedocs.io/en/latest/user/twtxtfile.html
func (twt *Tweet) String() string {
	r := strings.NewReplacer(
		"\n", "\\n",
		"\t", "\\t",
	)

	return twt.Time().Format(time.RFC3339 + "\t" + r.Replace(twt.Post()))
}
