package api

import (
	"errors"
)

var (
	ERR_INVALID_REQUEST = errors.New("invalid request")
	ERR_INVALID_ID      = errors.New("invalid id")
	ERR_INVALID_PAGE    = errors.New("invalid page number")
)
