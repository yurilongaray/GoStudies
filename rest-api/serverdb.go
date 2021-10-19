package rest_api

// https://www.udemy.com/course/curso-go/learn/lecture/8776196#overview
// https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "dbtest"
)

func CreateUser(db *sql.DB) {

	sqlStatement := `INSERT INTO users (age, first_name, last_name, email) VALUES ($1, $2, $3, $4) RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, 28, "yuri", "caldeira", "yurix@email.com").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}

func UpdateUser(db *sql.DB, id int, age int, first_name string, last_name string) {

	sqlStatement := `UPDATE users SET age = $2, first_name = $3, last_name = $4 WHERE id = $1`
	_, err := db.Exec(sqlStatement, id, age, first_name, last_name)
	if err != nil {
		panic(err)
	}
	fmt.Printf(`UPDATE users SET age = %v, first_name = %v, last_name = %v WHERE id = %v`, age, first_name, last_name, id)
}

func DeleteUser(db *sql.DB, id int) {

	sqlStatement := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	fmt.Printf(`DELETE FROM users WHERE id = %v`, id)
}
