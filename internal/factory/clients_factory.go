package factory

import (
	"errors"
	"fmt"

	"github.com/epileftro85/glorm/internal/consts"
	"github.com/epileftro85/glorm/internal/query_creator"
	"github.com/epileftro85/glorm/internal/strategies"
)

type ClientsFactory struct {
	builders map[consts.DatabaseType]query_creator.QueryCreator
}

func NewClientBuilderFactory() *ClientsFactory {
	return &ClientsFactory{
		builders: map[consts.DatabaseType]query_creator.QueryCreator{
			consts.PostgresSQL: &strategies.PostgresStrategy{},
		},
	}
}

func (f *ClientsFactory) Build(dbType consts.DatabaseType) (query_creator.QueryCreator, error) {
	builder, ok := f.builders[dbType]
	if !ok {
		message := fmt.Sprintf("database type %s is not supported", dbType)
		return nil, errors.New(message)
	}

	return builder, nil
}
