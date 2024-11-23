package consts

type QueryType string

const (
	SelectQuery QueryType = "SELECT"
	InsertQuery QueryType = "INSERT"
	UpdateQuery QueryType = "UPDATE"
	DeleteQuery QueryType = "DELETE"
	CountQuery  QueryType = "COUNT"
)
