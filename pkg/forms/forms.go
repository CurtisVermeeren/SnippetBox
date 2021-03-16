package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

// The Form struct hold the form data in url.Values and and validation errors
type Form struct {
	url.Values
	Errors errors
}

// Initialize a new from struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Check for required fields in the form and create errors for missing data
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Check that a field fits within the maximum allowed length
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum is %d characters)", d))
	}
}

// Check that a field contains petmitted values
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}

	for _, opt := range opts {
		if value == opt {
			return
		}
	}

	f.Errors.Add(field, "This field is invalid")
}

// Valid returns true if there are no errors in the form
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
