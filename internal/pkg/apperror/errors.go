package apperror

import (
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("Not found")
var ErrTokenExpired = errors.New("Not found")
var ErrUserNotFound = errors.New("User not found")
var ErrInternal = errors.New("Internal error")
