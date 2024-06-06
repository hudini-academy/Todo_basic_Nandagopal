package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")



type Todos struct {
	ID      int
	Name    string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID int
	Name string
	Email string
	HashedPassword []byte
	Created time.Time
	}