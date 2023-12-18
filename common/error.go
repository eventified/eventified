package common

import "net/http"

type Error struct {
	msg    string
	Status int
}

func newError(msg string, status int) *Error {
	return &Error{
		msg:    msg,
		Status: status,
	}
}

func NotFoundError(msg string) *Error {
	return newError(msg, http.StatusNotFound)
}

func InternalError(err error) *Error {
	return newError(err.Error(), http.StatusInternalServerError)
}

func BadRequestError(msg string) *Error {
	return newError(msg, http.StatusBadRequest)
}

func (err Error) Error() string {
	return err.msg
}
