package errors

import "errors"

var (
	ErrBadRequest     = errors.New("bad request")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrNotFound       = errors.New("not found")
	ErrConflict       = errors.New("conflict")
	ErrInternalServer = errors.New("internal server error")
)

func Is(err, target error) bool {
	return errors.Is(err, target)
}
