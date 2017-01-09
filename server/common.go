package server

import "database/sql"

func ParseRowsToCreditList(row *sql.Rows, c_list *[]Credit) {
	defer row.Close()
	for row.Next() {
		var c = &Credit{}
		c.ConventFormRow(row)
		*c_list = append(*c_list, *c)
	}
}