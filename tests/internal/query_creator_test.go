package internal

import (
	"testing"

	"github.com/epileftro85/glorm/internal/models"
	"github.com/epileftro85/glorm/internal/strategies"
	"github.com/epileftro85/glorm/pkg/utils"
)

func TestQueryCreator(t *testing.T) {
	pq := strategies.PostgresStrategy{}
	input := &models.QueryStructure{
		Table:  "users",
		Fields: []string{"id", "name", "email"},
	}

	expected := "SELECT id, name, email FROM users"
	var expectedArgs []interface{}

	query, args, err := pq.CreateSelect(input)

	if err != nil {
		t.Errorf("CreateSelect() error = %v", err)
	}

	if query != expected {
		t.Errorf("CreateSelect() = %v; query want %v", query, expected)
	}

	if !utils.CompareInterfaceSlices(args, expectedArgs) {
		t.Errorf("CreateSelect() = %v; args want %v", args, expectedArgs)
	}
}
