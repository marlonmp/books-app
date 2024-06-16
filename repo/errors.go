package repo

import "errors"

const (
	NotFoundErrorCode     = "resource_not_found"
	DoesNotExistErrorCode = "resource_does_not_exist"
)

// NotFoundError must be returned when a resource may exist but after
// searching found nothing.
type NotFoundError struct {
	err error
}

func (enf NotFoundError) Error() string {
	return "not found err: Resource not found"
}

func (enf NotFoundError) Unwrap() error {
	return enf.err
}

func (enf NotFoundError) Code() string {
	return NotFoundErrorCode
}

func (enf NotFoundError) Message() string {
	return ""
}

func IsNotFoundError(err error) bool {
	var enf *NotFoundError
	return errors.As(err, &enf)
}

// DoesNotExistError must be returned when a resource such as a file or endpoint
// does not exist or has not been created. In can also be used when an user
// does not have permission to access a file or endpoint.
type DoesNotExistError struct {
	err error
}

func (dne DoesNotExistError) Error() string {
	return "does not exist err: Resource does not exist"
}

func (dne DoesNotExistError) Unwrap() error {
	return dne.err
}

func (dne DoesNotExistError) Code() string {
	return DoesNotExistErrorCode
}

func (dne DoesNotExistError) Message() string {
	return ""
}

func IsDoesNotExistError(err error) bool {
	var dne *DoesNotExistError
	return errors.As(err, &dne)
}
