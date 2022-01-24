package twtxt

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFields(t *testing.T) {
	tests := []struct {
		name   string
		fields Fields
	}{
		{
			name:   "Empty",
			fields: Fields{},
		},
		{
			name: "SingleField",
			fields: Fields{
				&Field{
					key: "nick",
					val: "buckket",
				},
			},
		},
		{
			name: "MultipleFields",
			fields: Fields{
				&Field{
					key: "nick",
					val: "buckket",
				},
				&Field{
					key: "url",
					val: "https://example.org/buckket/twtxt.txt",
				},
				&Field{
					key: "description",
					val: "Author of twtxt",
				},
				&Field{
					key: "someNonStandardField",
					val: "someArbitraryValue",
				},
			},
		},
		{
			name: "MultipleFieldsWithDuplicateKeys",
			fields: Fields{
				&Field{
					key: "nick",
					val: "buckket",
				},
				&Field{
					key: "url",
					val: "https://example.org/buckket/twtxt.txt",
				},
				&Field{
					key: "url",
					val: "https://buckket.example.org/twtxt.txt",
				},
				&Field{
					key: "description",
					val: "Author of twtxt",
				},
				&Field{
					key: "someNonStandardField",
					val: "someArbitraryValue",
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			name := "Search(%s)"

			// default test cases
			searches := []string{
				"url",
				"nick",
				"avatar",
				"description",
				"follow",
				"following",
				"followers",
				"link",
				"prev",
				"refresh",
			}

			// populate test cases
			for _, field := range test.fields {
				var dup bool

				// skip invalid test cases
				if field == nil {
					continue
				}

				// search for previous duplicate test case
				for _, search := range searches {
					if search == field.key {
						dup = true
						break
					}
				}

				// add test case if not a duplicate
				if !dup {
					searches = append(searches, field.key)
				}
			}

			// run test cases
			for _, search := range searches {
				search := search

				t.Run(fmt.Sprintf(name, search), func(t *testing.T) {
					have := test.fields.Search(search)
					want := make(Fields, 0)

					for _, field := range test.fields {
						if field.key == search {
							want = append(want, field)
						}
					}

					if diff := cmp.Diff(have, want, cmp.AllowUnexported(Field{})); diff != "" {
						t.Errorf("diff:\n%s", diff)
					}
				})
			}
		})
	}
}
