package twtxt

import (
	"bufio"
	"io"
	"strings"
	"time"
)

// Tweets are a collection of Tweet instances, these are not necessarily from
// the same source, or in any particular order. However, a feed can be sorted by
// the timestamp (in a timezone aware manner).
type Tweets []*Tweet

// ParseTweets reads Tweets from a file or other source.
//
// See the twtxt file format specification for more information:
// https://twtxt.readthedocs.io/en/latest/user/twtxtfile.html
func ParseTweets(source io.Reader) (Tweets, error) {
	// if the source is nil, then there aren't any Tweets
	if source == nil {
		return Tweets{}, nil
	}

	// create a scanner to read the source
	scanner := bufio.NewScanner(source)

	// read the lines from the reader
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		// skip comment lines
		if strings.HasPrefix(line, "#") {
			continue
		}

		lines = append(lines, line)
	}

	// bail early if there is a reading error
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// parse each line into Tweet
	tweets := make(Tweets, len(lines))
	for i, line := range lines {
		// split the line by the tab delimiter
		parts := strings.SplitN(line, "\t", 2)

		// there has to be a timestamp
		if len(parts) < 1 || parts[0] == "" {
			return nil, &ParseError{
				line: uint64(i + 1),
				msg:  "missing timestamp",
			}
		}

		// there has to be a tab delimiter
		if len(parts) < 2 || !strings.Contains(line, "\t") {
			return nil, &ParseError{
				line: uint64(i + 1),
				msg:  "missing tab delimiter",
			}
		}

		// parse the timestamp
		t, err := time.Parse(time.RFC3339, parts[0])
		if err != nil {
			return nil, &ParseError{
				line:  uint64(i + 1),
				inner: err,
			}
		}

		// add the parsed tweet
		tweets[i] = &Tweet{
			time: t,
			post: parts[1],
		}
	}

	return tweets, nil
}

// Len reports the number of Tweets.
func (twts Tweets) Len() int { return len(twts) }

// Less reports if the Tweet at i was posted before the Tweet at j.
func (twts Tweets) Less(i, j int) bool { return twts[i].Before(twts[j]) }

// Swap exchanges the Tweet at i with the Tweet at j.
func (twts Tweets) Swap(i, j int) { twts[i], twts[j] = twts[j], twts[i] }
