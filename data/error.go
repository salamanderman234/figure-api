package data

import "errors"

var (
	ErrNotFound       = errors.New("data not found")
	ErrInternalServer = errors.New("something went wrong")
)
