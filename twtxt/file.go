package twtxt

import (
	"bufio"
	"io"
	"strings"
	"time"
)

// File represents a twtxt.txt file, it doesn't need to be called that, but
// twtxt.txt is traditional.
//
// File contains both Tweets that the twtxt.txt file contained, as well as any
// metadata Fields that were defined.
type File struct {
	Fields
	Tweets
}

// Parse reads a twtxt file from a twtxt.txt file or other source.
//
// See the twtxt file format specification for more information:
// https://twtxt.readthedocs.io/en/latest/user/twtxtfile.html
//
// Parse also supports community metadata fields, including the Yarn.Social
// metadata extensions: https://dev.twtxt.net/doc/metadataextension.html
func Parse(source io.Reader) (*File, error) {
	file := &File{
		make(Fields, 0),
		make(Tweets, 0),
	}

	// if the source is nil, then there aren't any Tweets
	if source == nil {
		return file, nil
	}

	// create a scanner to read the source
	scanner := bufio.NewScanner(source)

	// parse each line into Tweet
	var lineNumber uint64
	for scanner.Scan() {
		// increment line count
		lineNumber++

		// read next line
		line := scanner.Text()

		// skip comment lines
		if strings.HasPrefix(line, "#") {
			// TODO: try to parse fields here
			continue
		}

		// split the line by the tab delimiter
		parts := strings.SplitN(line, "\t", 2)

		// there has to be a timestamp
		if len(parts) < 1 || parts[0] == "" {
			return nil, &ParseError{
				line: lineNumber,
				msg:  "missing timestamp",
			}
		}

		// there has to be a tab delimiter
		if len(parts) < 2 || !strings.Contains(line, "\t") {
			return nil, &ParseError{
				line: lineNumber,
				msg:  "missing tab delimiter",
			}
		}

		// parse the timestamp
		t, err := time.Parse(time.RFC3339, parts[0])
		if err != nil {
			return nil, &ParseError{
				line:  lineNumber,
				inner: err,
			}
		}

		// add the parsed tweet
		file.Tweets = append(file.Tweets, &Tweet{
			time: t,
			post: parts[1],
		})
	}

	// bail early if there is a reading error
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return file, nil
}
