package models

import (
	"errors"
	"fmt"
)

var (
	// ErrInternalServerError to store internal server error message
	ErrInternalServerError = errors.New("internal Server Error")

	// ErrSomethingWrong is default error message
	ErrSomethingWrong = errors.New("Terjadi kesalahan")

	// ErrNotFound to store not found error message
	ErrNotFound = errors.New("item tidak ditemukan")

	// ErrConflict to store conflicted error message
	ErrConflict = errors.New("item sudah ada")

	// ErrUnauthorized to validate pds api error message
	ErrUnauthorized = errors.New("unauthorized")

	// ErrPassword to store password error message
	ErrPassword = errors.New("username atau Password yang digunakan tidak valid")
)

// DynamicErr to return parameterize errors
func DynamicErr(message string, args []interface{}) error {
	return fmt.Errorf(message, args...)
}

type ErrorValidationData struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
