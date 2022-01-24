package twtxt

import "fmt"

// ParseError represents an error that occurred while parsing a twtxt feed.
//
// A ParseError may occur when the feed is invalid according to the twtxt file
// specifications, or if a validly formatted feed has an invalid fields, e.g. if
// the timestamp is correctly formatted but doesn't represent a valid time
// according to RFC3339.
type ParseError struct {
	line  uint64
	inner error
	msg   string
}

// Error returns the error message of the parse error.
func (err *ParseError) Error() string {
	msg := "parse error"

	if err.line != 0 {
		msg += fmt.Sprintf(" on line %d", err.line)
	}

	msg += ": "

	switch {
	case err.msg == "" && err.inner == nil:
		msg += "could not parse twtxt feed"
	case err.msg == "" && err.inner != nil:
		msg += err.inner.Error()
	case err.msg != "" && err.inner == nil:
		msg += err.msg
	case err.msg != "" && err.inner != nil:
		msg += fmt.Sprintf("%s: %s", err.msg, err.inner)
	}

	return msg
}

// Unwrap returns the inner error of err, if there is one.
func (err *ParseError) Unwrap() error {
	return err.inner
}
