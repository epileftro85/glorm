package tests

import (
	"database/sql"
	"testing"

	"github.com/epileftro85/glorm"
	"github.com/epileftro85/glorm/internal/consts"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB

func TestGlormBuilderTable(t *testing.T) {
	builder := glorm.Builder(db)
	tableName := "some_table"
	builder.Table(tableName)

	assert.Equal(t, tableName, builder.QueryStructure.Table)
}

func TestGlormBuilderReturning(t *testing.T) {
	builder := glorm.Builder(db)
	returning := []string{"id", "name"}
	builder.Returning(returning)

	assert.Equal(t, returning, builder.QueryStructure.ReturnedValues)
}

func TestGlormBuilderSelect(t *testing.T) {
	builder := glorm.Builder(db)
	selectItems := []string{"id", "name"}
	builder.Select(selectItems)

	assert.Equal(t, selectItems, builder.QueryStructure.Fields)
	assert.Equal(t, consts.QueryType("SELECT"), builder.QueryStructure.QueryType)
}

func TestGlormBuilderDelete(t *testing.T) {
	builder := glorm.Builder(db)
	builder.Delete()

	assert.Equal(t, consts.QueryType("DELETE"), builder.QueryStructure.QueryType)
}

func TestGlormBuilderCount(t *testing.T) {
	builder := glorm.Builder(db)
	builder.Count()

	assert.Equal(t, consts.QueryType("COUNT"), builder.QueryStructure.QueryType)
}

func TestGlormBuilderOrder(t *testing.T) {
	builder := glorm.Builder(db)
	orderBy := "id"
	builder.OrderBy(orderBy)

	assert.Equal(t, orderBy, builder.QueryStructure.OrderBy)
}

func TestGlormBuilderLimit(t *testing.T) {
	builder := glorm.Builder(db)
	limit := 10
	builder.Limit(limit)

	assert.Equal(t, limit, builder.QueryStructure.Limit)
}

func TestGlormBuilderOffset(t *testing.T) {
	builder := glorm.Builder(db)
	offset := 100
	builder.Offset(offset)

	assert.Equal(t, offset, builder.QueryStructure.Offset)
}

func TestGlormBuilderInsert(t *testing.T) {
	builder := glorm.Builder(db)
	fields := map[string]interface{}{
		"name": "Andres",
		"last": "Clavijo",
	}
	builder.Insert(fields)
	expected := map[string]interface{}{
		"name":       "Andres",
		"last":       "Clavijo",
		"created_at": "NOW()",
		"updated_at": "NOW()",
	}
	assert.Equal(t, expected, builder.QueryStructure.InsertData)
}

func TestGlormBuilderUpdate(t *testing.T) {
	builder := glorm.Builder(db)
	fields := map[string]interface{}{
		"name": "Andrew",
		"last": "Clavijo",
	}
	builder.Update(fields)
	expected := map[string]interface{}{
		"name":       "Andrew",
		"last":       "Clavijo",
		"updated_at": "NOW()",
	}
	assert.Equal(t, expected, builder.QueryStructure.InsertData)
}

func TestGlormWhere(t *testing.T) {
	builder := glorm.Builder(db)
	condition := "id ="
	values := []interface{}{
		"10",
	}
	builder.Where(condition, values...)
	assert.Equal(t, []string([]string{condition}), builder.QueryStructure.WhereClauses)
	assert.Equal(t, values, builder.QueryStructure.Values)
}

func TestGlormJoin(t *testing.T) {
	builder := glorm.Builder(db)
	table, key1, key2 := "some_table", "table_one.one", "table_two.two"
	builder.Join(table, key1, key2)

	assert.Equal(t, []string([]string{"JOIN some_table ON table_one.one = table_two.two"}), builder.QueryStructure.Joins)
}
