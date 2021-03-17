package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord is a no records found error that can be used for consistency across different databases
	ErrNoRecord = errors.New("models: no matching record found")

	// ErrInvalidCredentials used when a user fails a login attempt
	ErrInvalidCredentials = errors.New("models: invalid credentials")

	// ErrDuplicateEmail used when a new user attemps to signup with a duplicate email
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

// Snippet models the fields of a snippet in the database
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User models an entry in the users table of the database
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}
