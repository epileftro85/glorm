package utils

import (
	"github.com/epileftro85/glorm/internal/consts"
	"os"
	"strings"
)

func ConfigTimestamps() (bool, string, string) {
	timestamps := os.Getenv("USE_TIMESTAMP")
	if timestamps == "false" {
		return false, os.Getenv("CREATED"), os.Getenv("UPDATED")
	}

	return true, "created_at", "updated_at"
}

func SetDBType() consts.DatabaseType {
	dbType := os.Getenv("DB_TYPE")
	if strings.ToLower(dbType) == "postgres" {
		return consts.PostgresSQL
	}

	return consts.MySQL
}
