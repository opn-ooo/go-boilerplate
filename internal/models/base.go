package models

// Base ... Base structure for models
type Base struct {
}

// TableName ... get table name
func (base *Base) TableName() string {
	return "table_name"
}
