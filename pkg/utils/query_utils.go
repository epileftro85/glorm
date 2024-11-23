package utils

import (
	"fmt"
	"strings"

	"github.com/epileftro85/glorm/internal/models"
)

func BuildWhereWithPlaceholders(data *models.QueryStructure, startIndex int, values *[]interface{}) string {
	whereClauses := []string{}
	paramIndex := startIndex

	for _, clause := range data.WhereClauses {
		whereClauses = append(whereClauses, fmt.Sprintf("%s %s%d", clause, data.Placeholder, paramIndex))
		paramIndex++
	}

	*values = append(*values, data.Values...)
	return strings.Join(whereClauses, " AND ")
}
