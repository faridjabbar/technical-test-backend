package helper

import (
	"fmt"
	"strings"
)

func SearchCondition(search string, columns []string) string {
	var conditions []string

	for _, column := range columns {
		conditions = append(conditions, fmt.Sprintf("LOWER(%s) LIKE LOWER('%%%s%%')", column, search))
	}

	return "(" + strings.Join(conditions, " OR ") + ")"
}
