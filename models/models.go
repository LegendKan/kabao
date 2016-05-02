package models

import (
	"errors"
)

var (
	ErrNoRows = errors.New("No row found")
	ErrArgs   = errors.New("args error may be empty")
	ErrExist  = errors.New("Have already exist")
)
