package twtxt

import (
	"fmt"
	"testing"
)

func TestField(t *testing.T) {
	tests := []struct {
		name  string
		field Field
	}{
		{
			name:  "Empty",
			field: Field{},
		},
		{
			name: "NameOnly",
			field: Field{
				key: "Foo",
			},
		},
		{
			name: "ValueOnly",
			field: Field{
				val: "Bar",
			},
		},
		{
			name: "BothSet",
			field: Field{
				key: "Baz",
				val: "Qux",
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			field := test.field

			t.Run("Name()", func(t *testing.T) {
				if have, want := field.Name(), field.key; have != want {
					t.Errorf("have %q, want %q", have, want)
				}
			})

			t.Run("Value()", func(t *testing.T) {
				if have, want := field.Value(), field.val; have != want {
					t.Errorf("have %q, want %q", have, want)
				}
			})

			t.Run("String()", func(t *testing.T) {
				have := field.String()
				want := fmt.Sprintf("# %s = %s", field.key, field.val)

				if have != want {
					t.Errorf("have %q, want %q", have, want)
				}
			})
		})
	}
}
