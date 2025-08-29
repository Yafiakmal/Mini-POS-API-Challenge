package service

import "errors"

var (
	ErrNotFound  = errors.New("resource not found")
	ErrDuplicate = errors.New("resource already exists")
	ErrInternal  = errors.New("internal server error")
)
