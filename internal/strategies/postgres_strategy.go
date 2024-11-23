package strategies

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/epileftro85/glorm/internal/models"
	"github.com/epileftro85/glorm/pkg/utils"
)

type PostgresStrategy struct{}

func (s PostgresStrategy) CreateSelect(data *models.QueryStructure) (string, []interface{}, error) {
	query := "SELECT "
	var values []interface{}

	if len(data.Fields) > 0 {
		query += strings.Join(data.Fields, ", ")
	}
	query += " FROM " + data.Table

	return query, values, nil
}

func (s PostgresStrategy) CreateUpdate(data *models.QueryStructure) (string, []interface{}, error) {
	setClauses := []string{}
	values := []interface{}{}
	paramIndex := 1 // PostgreSQL placeholders start at $1

	for column, value := range data.InsertData {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, paramIndex))
		values = append(values, value)
		paramIndex++
	}

	query := fmt.Sprintf("UPDATE %s SET %s", data.Table, strings.Join(setClauses, ", "))

	if len(data.WhereClauses) > 0 {
		query += " WHERE " + utils.BuildWhereWithPlaceholders(data, paramIndex, &values)
	}

	return query, values, nil
}

func (s PostgresStrategy) Execute(config *models.Config, query string, args ...interface{}) (sql.Result, error) {
	return config.Db.Exec(query, args)
}
