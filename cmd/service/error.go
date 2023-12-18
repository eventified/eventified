package service

import "strings"

type ValidationError struct {
	errs []string
}

func (err ValidationError) Error() string {
	return strings.Join(err.errs, "\n")
}
