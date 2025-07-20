package helper

import (
	"strings"
)

type DuplicateError struct {
	contains string
	message  string
}

func ErrorDuplicateMessage(err error) string {
	errors := []DuplicateError{
		{
			contains: "for key 'idx_user'",
			message:  "user already exists",
		},
	}

	for _, duplicateError := range errors {
		if strings.Contains(err.Error(), duplicateError.contains) {
			return duplicateError.message
		}
	}

	return "record already exists"
}
