package formats

import "errors"

const (
	ErrCode int = -10000 - iota
	ErrCodeBinding
	// append here, do not insert
)

var (
	ErrBinding = errors.New("binding error")
)
