package twtxt

import (
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	// "# this is a comment",
	// "# this = is a field",
	tests := []struct {
		name   string
		source io.Reader
		fields Fields
		tweets Tweets
		err    error
	}{
		{
			name:   "Nil",
			source: nil,
			fields: Fields{},
			tweets: Tweets{},
		},
		{
			name:   "Empty",
			source: strings.NewReader(""),
			fields: Fields{},
			tweets: Tweets{},
		},
		{
			name:   "SingleTweet",
			source: strings.NewReader("2022-01-19T14:14:00+13:00\tThis post contains newlines\\n\\n\\n"),
			fields: Fields{},
			tweets: Tweets{
				&Tweet{
					time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
					post: "This post contains newlines\\n\\n\\n",
				},
			},
		},
		{
			name: "MultipleTweets",
			source: strings.NewReader(strings.Join([]string{
				"2022-01-19T14:11:00+13:00\tThis post contains tabs\\t\\t\\t",
				"2016-02-03T23:05:00+01:00\t@<example http://example.org/twtxt.txt> welcome to twtxt!",
				"2022-01-19T14:14:00+13:00\tThis post contains newlines\\n\\n\\n",
				"2016-02-01T11:00:00+01:00\tThis is just another example.",
				"2015-12-12T12:00:00+01:00\tFiat lux!",
				"2016-02-04T13:30:00+01:00\tYou can really go crazy here! ┐(ﾟ∀ﾟ)┌",
			}, "\n")),
			fields: Fields{},
			tweets: Tweets{
				&Tweet{
					time: time.Date(2022, 1, 19, 14, 11, 0, 0, loc(+13)),
					post: "This post contains tabs\\t\\t\\t",
				},
				&Tweet{
					time: time.Date(2016, 2, 3, 23, 5, 0, 0, loc(+1)),
					post: "@<example http://example.org/twtxt.txt> welcome to twtxt!",
				},
				&Tweet{
					time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
					post: "This post contains newlines\\n\\n\\n",
				},
				&Tweet{
					time: time.Date(2016, 2, 1, 11, 0, 0, 0, loc(+1)),
					post: "This is just another example.",
				},
				&Tweet{
					time: time.Date(2015, 12, 12, 12, 0, 0, 0, loc(+1)),
					post: "Fiat lux!",
				},
				&Tweet{
					time: time.Date(2016, 2, 4, 13, 30, 0, 0, loc(+1)),
					post: "You can really go crazy here! ┐(ﾟ∀ﾟ)┌",
				},
			},
		},
		{
			name:   "SingleComment",
			source: strings.NewReader("# this is a comment"),
			fields: Fields{},
			tweets: Tweets{},
		},
		{
			name: "MultipleComments",
			source: strings.NewReader(strings.Join([]string{
				"# this is a comment",
				"# this is also a comment",
				"# huh, no tweets yet, just another comment",
			}, "\n")),
			fields: Fields{},
			tweets: Tweets{},
		},
		{
			name:   "SingleField",
			source: strings.NewReader("# this = is a field"),
			fields: Fields{
				&Field{
					key: "this",
					val: "is a field",
				},
			},
			tweets: Tweets{},
		},
		{
			name: "MultipleFields",
			source: strings.NewReader(strings.Join([]string{
				"# this = is a field",
				"# this = is also a field",
				"# huh = no tweets yet, just another field",
			}, "\n")),
			fields: Fields{
				&Field{
					key: "this",
					val: "is a field",
				},
				&Field{
					key: "this",
					val: "is also a field",
				},
				&Field{
					key: "huh",
					val: "no tweets yet, just another field",
				},
			},
			tweets: Tweets{},
		},
		{
			name: "MixedFieldsCommentsAndTweets",
			source: strings.NewReader(strings.Join([]string{
				"# this is a comment",
				"# this is also a comment",
				"# huh, no tweets yet, just another comment",
				"# this = is a field",
				"# this = is also a field",
				"# huh = no tweets yet, just another field",
				"2022-01-19T14:11:00+13:00\tThis post contains tabs\\t\\t\\t",
				"2016-02-03T23:05:00+01:00\t@<example http://example.org/twtxt.txt> welcome to twtxt!",
				"2022-01-19T14:14:00+13:00\tThis post contains newlines\\n\\n\\n",
				"2016-02-01T11:00:00+01:00\tThis is just another example.",
				"2015-12-12T12:00:00+01:00\tFiat lux!",
				"2016-02-04T13:30:00+01:00\tYou can really go crazy here! ┐(ﾟ∀ﾟ)┌",
			}, "\n")),
			fields: Fields{
				&Field{
					key: "this",
					val: "is a field",
				},
				&Field{
					key: "this",
					val: "is also a field",
				},
				&Field{
					key: "huh",
					val: "no tweets yet, just another field",
				},
			},
			tweets: Tweets{
				&Tweet{
					time: time.Date(2022, 1, 19, 14, 11, 0, 0, loc(+13)),
					post: "This post contains tabs\\t\\t\\t",
				},
				&Tweet{
					time: time.Date(2016, 2, 3, 23, 5, 0, 0, loc(+1)),
					post: "@<example http://example.org/twtxt.txt> welcome to twtxt!",
				},
				&Tweet{
					time: time.Date(2022, 1, 19, 14, 14, 0, 0, loc(+13)),
					post: "This post contains newlines\\n\\n\\n",
				},
				&Tweet{
					time: time.Date(2016, 2, 1, 11, 0, 0, 0, loc(+1)),
					post: "This is just another example.",
				},
				&Tweet{
					time: time.Date(2015, 12, 12, 12, 0, 0, 0, loc(+1)),
					post: "Fiat lux!",
				},
				&Tweet{
					time: time.Date(2016, 2, 4, 13, 30, 0, 0, loc(+1)),
					post: "You can really go crazy here! ┐(ﾟ∀ﾟ)┌",
				},
			},
		},
		{
			name: "MissingTabDelimiter",
			err: &ParseError{
				line: 3,
				msg:  "missing tab delimiter",
			},
			source: strings.NewReader(strings.Join([]string{
				"# this is a comment",
				"# this = is a field",
				"2022-01-19T14:11:00+13:00 This post contains tabs\\t\\t\\t",
				"2016-02-03T23:05:00+01:00 @<example http://example.org/twtxt.txt> welcome to twtxt!",
				"2022-01-19T14:14:00+13:00 This post contains newlines\\n\\n\\n",
				"2016-02-01T11:00:00+01:00 This is just another example.",
				"2015-12-12T12:00:00+01:00 Fiat lux!",
				"2016-02-04T13:30:00+01:00 You can really go crazy here! ┐(ﾟ∀ﾟ)┌",
			}, "\n")),
		},
		{
			name: "MissingTimestamp",
			err: &ParseError{
				line: 3,
				msg:  "missing timestamp",
			},
			source: strings.NewReader(strings.Join([]string{
				"# this is a comment",
				"# this = is a field",
				"\tThis post contains tabs\\t\\t\\t",
				"\t@<example http://example.org/twtxt.txt> welcome to twtxt!",
				"\tThis post contains newlines\\n\\n\\n",
				"\tThis is just another example.",
				"\tFiat lux!",
				"\tYou can really go crazy here! ┐(ﾟ∀ﾟ)┌",
			}, "\n")),
		},
		{
			name: "InvalidTimestamp",
			err: &ParseError{
				line:  4,
				inner: errors.New("parsing time \"2016-02-74T23:05:00+01:00\": day out of range"),
			},
			source: strings.NewReader(strings.Join([]string{
				"# this is a comment",
				"# this = is a field",
				"2022-01-19T14:11:00+13:00\tThis post contains tabs\\t\\t\\t",
				"2016-02-74T23:05:00+01:00\t@<example http://example.org/twtxt.txt> welcome to twtxt!",
				"2022-01-19T14:14:00+13:00\tThis post contains newlines\\n\\n\\n",
				"2016-13-01T11:00:60+01:00\tThis is just another example.",
				"2015-12-12T12:00:00+99:00\tFiat lux!",
				"2016-02-04T60:30:00+01:00\tYou can really go crazy here! ┐(ﾟ∀ﾟ)┌",
			}, "\n")),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			file, err := Parse(test.source)

			if test.err == nil {
				// want a file
				if file == nil {
					t.Error("want file but got nil")
				}

				// want no error
				if err != nil {
					t.Fatalf("want nil error but got %q", err)
				}

				// file should have expected fields
				if diff := cmp.Diff(file.Fields, test.fields, cmp.AllowUnexported(Field{})); diff != "" {
					t.Errorf("diff:\n%s", diff)
				}

				// file should have expected tweets
				if diff := cmp.Diff(file.Tweets, test.tweets, cmp.AllowUnexported(Tweet{})); diff != "" {
					t.Errorf("diff:\n%s", diff)
				}
			} else {
				// want no file
				if file != nil {
					t.Errorf("want nil file but got %v", file)
				}

				// want error
				if err == nil {
					t.Error("want error but got nil")
				}

				// error should match expected err
				if _, wantParseError := test.err.(*ParseError); wantParseError {
					if _, haveParseError := err.(*ParseError); !haveParseError {
						t.Errorf("have %T, want %T", err, test.err)
					}
				}

				if have, want := err, test.err; have.Error() != want.Error() {
					t.Errorf("have %q, want %q", have, want)
				}
			}
		})
	}
}
