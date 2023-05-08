package vars

import "errors"

var (
	ErrParamEmpty         = errors.New("input parameters cannot be empty")
	ErrMethodNotSupported = errors.New("method not supported")
	ErrTokenEmpty         = errors.New("token cannot be empty")
)
