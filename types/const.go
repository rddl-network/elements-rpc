package types

import (
	"errors"
)

var (
	ErrMissingSignatures = errors.New("transaction does not have the complete set of signatures")
)
