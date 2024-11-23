package query_creator

import (
	"database/sql"

	"github.com/epileftro85/glorm/internal/models"
)

type QueryCreator interface {
	CreateSelect(structure *models.QueryStructure) (string, []interface{}, error)
	Execute(config *models.Config, query string, args ...interface{}) (sql.Result, error)
}
