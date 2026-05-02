package notification

import (
	"errors"
)

var (
	ErrInvalidPagination = errors.New("both afterID and beforeID cannot be set")
	ErrInvalidLimit      = errors.New("limit is not valid")
)
