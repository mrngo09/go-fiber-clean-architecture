package common

import (
	"errors"
	"fmt"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"root_err"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: 400,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(500, err,
		"something went wrong in the server", err.Error(), "ErrInternal")
}

func ErrParseJson(err error) *AppError {
	return NewErrorResponse(err, "something went wrong with parse json", err.Error(), "ErrParseJson")
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrDataNotFound(entity string) *AppError {
	return NewFullErrorResponse(404,
		nil,
		fmt.Sprintf("%s Not Found", strings.ToLower(entity)),
		fmt.Sprintf("ErrDataNotFound%s", entity),
		"ErrDataNotFound",
	)
}

func ErrDataAlreadyExist(entity string, field string) *AppError {
	return NewCustomError(
		nil,
		fmt.Sprintf("%s %s already exist", strings.ToLower(field), strings.ToLower(entity)),
		fmt.Sprintf("Err%s%sAlreadyExist", field, entity),
	)
}

var (
	ErrUnAuthorization  = NewFullErrorResponse(401, nil, "UnAuthorization", "UnAuthorization", "ErrUnAuthorization")
	ErrPermissionDenied = NewFullErrorResponse(403, nil, "You don't have permission to do this action", "You don't have permission to do this action", "ErrPermissionDenied")
)
