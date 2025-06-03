package repository

import "errors"

var (
	ErrConstraintViolation = errors.New("constraint violation")
	ErrNotFound            = errors.New("not found")
	ErrInternal            = errors.New("internal error")
)
