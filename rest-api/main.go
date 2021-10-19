package main

// https://www.udemy.com/course/curso-go/learn/lecture/8776196#overview
// https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connected!")

	/*
		CREATE TABLE IF NOT EXISTS users (
		  id SERIAL PRIMARY KEY,
		  age INT,
		  first_name TEXT,
		  last_name TEXT,
		  email TEXT UNIQUE NOT NULL
		);
	*/

	// CreateUser(db)
	// UpdateUser(db, 4, 29, "yuritest", "caldeiratest")
	// DeleteUser(db, 4)

}
