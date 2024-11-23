package examples

import (
	"database/sql"
	"fmt"

	"github.com/epileftro85/glorm"
)

var pool *sql.DB

func SimpleSelect() sql.Result {
	glorm := glorm.Builder(pool)
	row, err := glorm.Exec()

	if err != nil {
		fmt.Printf("Error on getting result: %v", err)
	}

	return row
}
