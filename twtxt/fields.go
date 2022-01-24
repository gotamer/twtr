package twtxt

// Fields is a collection of Field instances, and allows for searching and
// accessing the field's values.
type Fields []*Field

// Search retrieves all Fields with a given name, if no field exists with that
// name, zero Fields are returned. Search never returns nil.
func (fields Fields) Search(name string) Fields {
	found := make(Fields, 0)

	for _, field := range fields {
		if field.Name() == name {
			found = append(found, field)
		}
	}

	return found
}
