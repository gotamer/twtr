package twtxt

import (
	"strings"
	"time"
)

// Tweet represents a single twtxt post, containing the timestamp and the main
// body of the post.
type Tweet struct {
	Time time.Time
	Post string
}

// String formats the Tweet as an entry into a twtxt.txt file, returns the
// timestamp followed by a tab character and the post message. Any tabs or new
// line characters are escaped to prevent invalid formating of the twtxt.txt
// file.
//
//     <yyyy>-<mm>-<dd>T<HH>:<MM>:<SS>+<XX>:<ZZ>\t<POST>
//
// See the format specification for more details on the file format:
// https://twtxt.readthedocs.io/en/latest/user/twtxtfile.html
func (twt *Tweet) String() string {
	r := strings.NewReplacer(
		"\n", "\\n",
		"\t", "\\t",
	)

	return twt.Time.Format(time.RFC3339 + "\t" + r.Replace(twt.Post))
}
