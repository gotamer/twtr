package twtxt

import (
	"errors"
	"testing"
)

//type ParseError struct {
//	line  uint64
//	inner error
//	msg   string
//}

func TestParseError(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		err  *ParseError
	}{
		{
			name: "ZeroValue",
			msg:  "parse error: could not parse twtxt feed",
			err:  &ParseError{},
		},
		{
			name: "LineNumberOnly",
			msg:  "parse error on line 23452: could not parse twtxt feed",
			err: &ParseError{
				line: 23452,
			},
		},
		{
			name: "MessageOnly",
			msg:  "parse error: test message",
			err: &ParseError{
				msg: "test message",
			},
		},
		{
			name: "InnerErrorOnly",
			msg:  "parse error: test error",
			err: &ParseError{
				inner: errors.New("test error"),
			},
		},
		{
			name: "LineNumberAndMessage",
			msg:  "parse error on line 1234: test message",
			err: &ParseError{
				line: 1234,
				msg:  "test message",
			},
		},
		{
			name: "LineNumberAndInnerError",
			msg:  "parse error on line 1: test error",
			err: &ParseError{
				line:  1,
				inner: errors.New("test error"),
			},
		},
		{
			name: "LineNumberAndMessageAndInnerError",
			msg:  "parse error on line 10023142: test message: test error",
			err: &ParseError{
				line:  10023142,
				msg:   "test message",
				inner: errors.New("test error"),
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Run("Error()", func(t *testing.T) {
				if have, want := test.err.Error(), test.msg; have != want {
					t.Errorf("have %q, want %q", have, want)
				}
			})

			t.Run("Unwrap()", func(t *testing.T) {
				if have, want := test.err.Unwrap(), test.err.inner; have != want {
					t.Errorf("have %q, want %q", have, want)
				}
			})
		})
	}
}
