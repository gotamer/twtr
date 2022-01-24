package twtxt

// Field is a key value pair as defined in the metadata comments of a twtxt
// file.
type Field struct {
	key, val string
}

// Name returns the name (or key) of the field.
func (field *Field) Name() string {
	return field.key
}

// Value returns the value of the field.
func (field *Field) Value() string {
	return field.val
}

// String returns the field formatted as a comment line of metadata.
func (field *Field) String() string {
	return "# " + field.Name() + " = " + field.Value()
}
