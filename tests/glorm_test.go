package tests

import (
	"database/sql"
	"reflect"
	"strings"
	"testing"

	"github.com/epileftro85/glorm"
	"github.com/epileftro85/glorm/pkg/utils"
)

var db *sql.DB

func getQueryStructureItem(g *glorm.Glorm, item string) reflect.Value {
	v := reflect.ValueOf(g).Elem()
	return v.FieldByName("queryStructure").Elem().FieldByName(item)
}

func TestGlormBuilderTable(t *testing.T) {
	builder := glorm.Builder(db)
	tableName := "some_table"
	builder.Table(tableName)

	tableField := getQueryStructureItem(builder, "Table")

	if tableName != tableField.String() {
		t.Errorf("expected %s, got %s", tableName, tableField)
	}
}

func TestGlormBuilderReturning(t *testing.T) {
	builder := glorm.Builder(db)
	returning := []string{"id", "name"}
	builder.Returning(returning)

	returnedField := getQueryStructureItem(builder, "ReturnedValues")

	var actualValues []string
	for i := 0; i < returnedField.Len(); i++ {
		actualValues = append(actualValues, returnedField.Index(i).String())
	}

	got := strings.Join(actualValues, ",")
	expected := strings.Join(returning, ",")

	if expected != got {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func TestGlormBuilderSelect(t *testing.T) {
	builder := glorm.Builder(db)
	selectItems := []string{"id", "name"}
	builder.Select(selectItems)
	selectField := getQueryStructureItem(builder, "Fields")
	queryType := getQueryStructureItem(builder, "QueryType")

	var actualValues []string
	for i := 0; i < selectField.Len(); i++ {
		actualValues = append(actualValues, selectField.Index(i).String())
	}

	got := strings.Join(actualValues, ",")
	expected := strings.Join(selectItems, ",")

	if expected != got {
		t.Errorf("expected %s, got %s", expected, got)
	}
	if queryType.String() != "SELECT" {
		t.Errorf("expected query type %s, got %s", "SELECT", queryType.String())
	}
}

func TestGlormBuilderInsert(t *testing.T) {
	builder := glorm.Builder(db)
	fields := map[string]interface{}{
		"name": "Andres",
		"last": "Clavijo",
	}
	builder.Insert(fields)
	selectField := getQueryStructureItem(builder, "InsertData")
	queryType := getQueryStructureItem(builder, "QueryType")

	if queryType.String() != "INSERT" {
		t.Errorf("expected query type %s, got %s", "INSERT", queryType.String())
	}

	comparables, _ := utils.ConvertMap(selectField)
	if !utils.CompareMaps(fields, comparables) {
		t.Errorf("CreateSelect() = %v; args want %v", fields, comparables)
	}
}
