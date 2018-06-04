package generator

import (
	"database/sql"
	"fmt"
)

type DatabaseNumberGenerator struct {
	Db *sql.DB
}

func (generator *DatabaseNumberGenerator) Next() string {
	rows, _ := generator.Db.Query("SELECT MAX(number) FROM account;")
	var id int

	for rows.Next() {
		rows.Scan(&id)
		id++
	}
	if string(id) == "" {
		id = 1
	}
	return fmt.Sprintf("%026d", id)
}
