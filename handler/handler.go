package helper

import (
	"fmt"
	"runtime/debug"
	"technical-test-backend/exception"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(databaseErrors []exception.DatabaseError) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
				exception.ErrorHandler(c, err, databaseErrors)
			}
		}()
		c.Next()
	}
}
