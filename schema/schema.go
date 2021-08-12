package schema

import (
	"go/ast"
	"reflect"
	"yaorm/dialect"
)

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func (s *Schema) GetField(name string) *Field {
	return s.fieldMap[name]
}

func Parse(obj interface{}, d dialect.Dialect) *Schema {
	typ := reflect.Indirect(reflect.ValueOf(obj)).Type()
	schema := &Schema{
		Model:    obj,
		Name:     typ.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if !f.Anonymous && ast.IsExported(f.Name) {
			field := &Field{
				Name: f.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(f.Type))),
			}

			if v, ok := f.Tag.Lookup("yaorm"); ok {
				field.Tag = v
			}

			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, f.Name)
			schema.fieldMap[f.Name] = field
		}
	}

	return schema
}

func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldVars []interface{}
	for _, field := range schema.Fields {
		fieldVars = append(fieldVars, destValue.FieldByName(field.Name).Interface())
	}

	return fieldVars
}
