package error

import "errors"

// ErrNotFound is used when a beer could not be found.
var ErrNotFound = errors.New("beer not found")

// ErrDuplicate is used when a beer already existed.
var ErrDuplicate = errors.New("beer already exists")
