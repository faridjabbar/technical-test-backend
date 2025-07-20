package helper

import "technical-test-backend/exception"

var DatabaseError = []exception.DatabaseError{
	{
		Contains: "for key 'idx_user'",
		Message:  "user already exists",
	},
}
