package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")


//created  todos for collecting information from the form.
type Todos struct {
	ID      int
	Name    string
	Created time.Time
	Expires time.Time
}

// created the user for collecting information in login/sginup.
type User struct {
	ID int
	Name string
	Email string
	HashedPassword []byte
	Created time.Time
	}

// create another 'special' for collecting information from home that contains
// 'special:' keyword.

type Special struct{
	ID int 
	Name string
	Created time.Time
	Expires time.Time
}