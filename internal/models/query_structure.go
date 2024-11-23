package models

import (
	"github.com/epileftro85/glorm/internal/consts"
)

type QueryStructure struct {
	QueryType      consts.QueryType
	Table          string
	Fields         []string
	WhereClauses   []string
	Joins          []string
	Values         []interface{}
	InsertData     map[string]interface{}
	Limit          int
	Offset         int
	OrderBy        string
	ReturnedValues []string
	Placeholder    string
	CreatedAt      string
	UpdatedAt      string
}
