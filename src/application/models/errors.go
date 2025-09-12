package models

import "errors"

var (
	ErrModelInUse        = errors.New("model is in use")
	ErrModelNotFound     = errors.New("model not found")
	ErrCategoryNotFound  = errors.New("category not found")
	ErrFieldValidation   = errors.New("field validation failed")
	ErrDataNotFound      = errors.New("data not found")
	ErrDeleteFailed      = errors.New("delete failed")
	ErrUpdateFailed      = errors.New("update failed")
	ErrInsertFailed      = errors.New("insert failed")
	ErrInvalidParameter  = errors.New("invalid parameter")
	ErrPermissionDenied  = errors.New("permission denied")
	ErrInternal          = errors.New("internal error")
)
