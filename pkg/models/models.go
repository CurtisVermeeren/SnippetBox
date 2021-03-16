package models

import (
	"errors"
	"time"
)

// ErrNoRecord is a no records found error that can be used for consistency across different databases
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet models the fields of a snippet in the database
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
