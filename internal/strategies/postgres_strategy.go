package strategies

import (
	"database/sql"
	"strings"

	"github.com/epileftro85/glorm/internal/models"
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

func (s PostgresStrategy) Execute(config *models.Config, query string, args ...interface{}) (sql.Result, error) {
	return config.Db.Exec(query, args)
}
