package models

import (
	"database/sql"
	"github.com/epileftro85/glorm/internal/consts"
)

type Config struct {
	Db             *sql.DB
	DbType         consts.DatabaseType
	UseTimestamps  bool
	TimestampsName []string
}
