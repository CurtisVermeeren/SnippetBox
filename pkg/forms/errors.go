package forms

// errors type will be used to hold form validation error messages
type errors map[string][]string

// add an error message for the given field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// retrieve an errofr message for the given field
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
