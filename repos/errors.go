package repos

import (
	"errors"

	"github.com/jackc/pgx"
)

type errorCode string

const (
	NotFoundErrorCode     errorCode = "resource_not_found"
	DoesNotExistErrorCode errorCode = "resource_does_not_exist"

	ConflictErrorCode errorCode = "resource_already_exist"

	InvalidCredentialsErrorCode errorCode = "invalid_authentication_credentials"
	MissingCredentialsErrorCode errorCode = "missing_authentication_credentials"
)

// NotFoundError must be returned when a resource may exist but after
// searching found nothing.
type NotFoundError struct {
	err error
}

func (enf NotFoundError) Error() string {
	return "not found: this resource not found"
}

func (enf NotFoundError) Unwrap() error {
	return enf.err
}

func (enf NotFoundError) Code() errorCode {
	return NotFoundErrorCode
}

func IsNotFoundError(err error) bool {
	var enf *NotFoundError
	return errors.As(err, &enf)
}

// if the err is compatible, it gets wrapped into a NotFoundError
// and returns true, else do nothig and returns false
func AsNotFoundError(err *error) bool {
	if err == nil || *err == nil {
		return false
	}

	if errors.Is(*err, pgx.ErrNoRows) {
		*err = NotFoundError{*err}
		return true
	}

	return false
}

// DoesNotExistError must be returned when a resource such as a file or endpoint
// does not exist or has not been created. In can also be used when an user
// does not have permission to access a file or endpoint.
type DoesNotExistError struct {
	err error
}

func (dne DoesNotExistError) Error() string {
	return "does not exist: this resource does not exist"
}

func (dne DoesNotExistError) Unwrap() error {
	return dne.err
}

func (dne DoesNotExistError) Code() errorCode {
	return DoesNotExistErrorCode
}

func IsDoesNotExistError(err error) bool {
	var dne *DoesNotExistError
	return errors.As(err, &dne)
}

type ConflictError struct {
	err error
}

func (ce ConflictError) Error() string {
	return "conflict: this resource already exist"
}

func (ce ConflictError) Unwrap() error {
	return ce.err
}

func (ce ConflictError) Code() errorCode {
	return ConflictErrorCode
}

func IsConflictError(err error) bool {
	var ce *ConflictError
	return errors.As(err, &ce)
}

// InvalidCredentialsError must be returned when the provided credentials
// are invalid
type InvalidCredentialsError struct {
	err error
}

func (ice InvalidCredentialsError) Error() string {
	return "invalid credentials: the provided credentials are invalid"
}

func (ice InvalidCredentialsError) Unwrap() error {
	return ice.err
}

func (ice InvalidCredentialsError) Code() errorCode {
	return InvalidCredentialsErrorCode
}

// MissingCredentialsError must be returned when the credentials are not provided
type MissingCredentialsError struct {
	err error
}

func (ice MissingCredentialsError) Error() string {
	return "missing credentials: the credentials were not provided"
}

func (ice MissingCredentialsError) Unwrap() error {
	return ice.err
}

func (ice MissingCredentialsError) Code() errorCode {
	return MissingCredentialsErrorCode
}
