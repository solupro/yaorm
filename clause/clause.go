package clause

import "strings"

type Type int

const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDER_BY
	UPDATE
	DELETE
	COUNT
)

type Clause struct {
	sql     map[Type]string
	sqlVars map[Type][]interface{}
}

func (c *Clause) Set(name Type, vars ...interface{}) {
	if nil == c.sql {
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]interface{})
	}
	sql, vs := generators[name](vars...)

	c.sql[name] = sql
	c.sqlVars[name] = vs
}

func (c *Clause) Build(orders ...Type) (string, []interface{}) {
	var sqls []string
	var vars []interface{}
	for _, order := range orders {
		if sql, ok := c.sql[order]; ok {
			sqls = append(sqls, sql)
			vars = append(vars, c.sqlVars[order]...)
		}
	}
	return strings.Join(sqls, " "), vars
}
