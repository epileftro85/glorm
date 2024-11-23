package glorm

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/epileftro85/glorm/internal/consts"
	"github.com/epileftro85/glorm/internal/factory"
	"github.com/epileftro85/glorm/internal/models"
	"github.com/epileftro85/glorm/pkg/utils"
)

type queryBuilderFunc func(data *models.QueryStructure) (string, []interface{}, error)

// Glorm package is a query builder which is mean to just do that
// nothing fancy, just queries. Glorm stands for Go Light ORM
type Glorm struct {
	config         models.Config
	QueryStructure *models.QueryStructure
	factory        factory.ClientsFactory
}

func Builder(db *sql.DB) *Glorm {
	useTimestamps, created, updated := utils.ConfigTimestamps()
	factory.NewClientBuilderFactory()
	return &Glorm{
		config: models.Config{
			Db:             db,
			UseTimestamps:  useTimestamps,
			DbType:         utils.SetDBType(),
			TimestampsName: []string{created, updated},
		},
		QueryStructure: &models.QueryStructure{
			InsertData:     make(map[string]interface{}),
			WhereClauses:   []string{},
			Fields:         []string{},
			ReturnedValues: []string{},
			Joins:          []string{},
		},
	}
}

// Method for set the main table of query
func (s *Glorm) Table(table string) *Glorm {
	s.QueryStructure.Table = table
	return s
}

// Method used to indicate which fields wants to return in update or create query
func (s *Glorm) Returning(fields []string) *Glorm {
	s.QueryStructure.ReturnedValues = append(s.QueryStructure.ReturnedValues, fields...)
	return s
}

// Set the fields to be selected, if empty all (*) will be used
func (s *Glorm) Select(fields []string) *Glorm {
	s.QueryStructure.QueryType = consts.SelectQuery
	s.QueryStructure.Fields = fields
	return s
}

// Method used to indicate which data will be updated
func (s *Glorm) Insert(data map[string]interface{}) *Glorm {
	s.QueryStructure.QueryType = consts.InsertQuery
	s.setInsertData(data)
	s.setTimestamps(true)
	return s
}

func (s *Glorm) Update(data map[string]interface{}) *Glorm {
	s.QueryStructure.QueryType = consts.UpdateQuery
	s.setInsertData(data)
	s.setTimestamps(false)
	return s
}

func (s *Glorm) Delete() *Glorm {
	s.QueryStructure.QueryType = consts.DeleteQuery
	return s
}

func (s *Glorm) Count() *Glorm {
	s.QueryStructure.QueryType = consts.CountQuery
	return s
}

func (s *Glorm) Where(condition string, args ...interface{}) *Glorm {
	s.QueryStructure.WhereClauses = append(s.QueryStructure.WhereClauses, condition)
	s.QueryStructure.Values = append(s.QueryStructure.Values, args...)
	return s
}

func (s *Glorm) Limit(limit int) *Glorm {
	s.QueryStructure.Limit = limit
	return s
}

func (s *Glorm) Offset(offset int) *Glorm {
	s.QueryStructure.Offset = offset
	return s
}

func (s *Glorm) OrderBy(order string) *Glorm {
	s.QueryStructure.OrderBy = order
	return s
}

func (s *Glorm) Join(table string, key1 string, key2 string) *Glorm {
	s.QueryStructure.Joins = append(s.QueryStructure.Joins, fmt.Sprintf("JOIN %s ON %s = %s", table, key1, key2))
	return s
}

func (s *Glorm) setInsertData(data map[string]interface{}) {
	for key, value := range data {
		s.QueryStructure.InsertData[key] = value
	}
}

func (s *Glorm) setTimestamps(bothTimestamps bool) {
	if bothTimestamps {
		s.QueryStructure.InsertData[s.config.TimestampsName[0]] = "NOW()"
		s.QueryStructure.InsertData[s.config.TimestampsName[1]] = "NOW()"
		return
	}

	s.QueryStructure.InsertData[s.config.TimestampsName[1]] = "NOW()"
}

func (s *Glorm) getBuilders() (map[consts.QueryType]queryBuilderFunc, error) {
	builder, err := s.factory.Build(s.config.DbType)
	if err != nil {
		return nil, err
	}

	return map[consts.QueryType]queryBuilderFunc{
		consts.SelectQuery: builder.CreateSelect,
		consts.InsertQuery: builder.CreateUpdate,
	}, nil
}

func (s *Glorm) getQueryAndParams() (string, []interface{}, error) {
	builder, err := s.getBuilders()
	if err != nil {
		log.Fatalf("Error getting builder: %v", err)
		return "fail", nil, err

	}
	query, args, err := builder[s.QueryStructure.QueryType](s.QueryStructure)
	if err != nil {
		log.Fatalf("Error getting the query builder with Error: %v", err)
		return "fail", nil, err
	}

	return query, args, nil
}

func (s *Glorm) Exec() (sql.Result, error) {
	query, args, err := s.getQueryAndParams()
	if err != nil {
		log.Fatalf("Error getting the query: %v", err)
		return nil, err
	}

	return s.config.Db.Exec(query, args...)
}
