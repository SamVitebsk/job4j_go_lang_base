package tracker

import "errors"

var ErrNotFound = errors.New("not found")
var ErrItemAlreadyExists = errors.New("already exists")
