package dialect

import "reflect"

type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistsSQL(tableName string) (string, []interface{})
}

var dialectMap = map[string]Dialect{}

func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
