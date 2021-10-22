package main

// https://www.udemy.com/course/curso-go/learn/lecture/8776196#overview
// https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

/*
	CREATE TABLE IF NOT EXISTS employee (
	  id SERIAL PRIMARY KEY,
	  name TEXT,
	  email TEXT UNIQUE NOT NULL
	);
*/
type employee struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	createHTTPServer()

	// CreateEmployee(db)
	// UpdateEmployee(db, 4, 29, "yuritest", "caldeiratest")
	// DeleteEmployee(db, 4)
	// getEmployeeById(1)
}

func createHTTPServer() {
	http.HandleFunc("/employees/", EmployeeHandler)
	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

/* Database Topic */
func createPGConnection() *sql.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "dbtest"
	)

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

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connected!")

	return db
}

func CreateEmployee(db *sql.DB) {

	sqlStatement := `INSERT INTO employee (name, email) VALUES ($1, $2) RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, "yuri", "yurix@email.com").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}

func UpdateEmployee(db *sql.DB, id int, name, email string) {

	sqlStatement := `UPDATE employee SET name =$2, email = $3 WHERE id = $1`
	_, err := db.Exec(sqlStatement, id, name, email)
	if err != nil {
		panic(err)
	}
	fmt.Printf(`UPDATE employee SET name =%v, email = %v WHERE id = %v`, name, email, id)
}

func DeleteEmployee(db *sql.DB, id int) {

	sqlStatement := `DELETE FROM employee WHERE id = $1`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	fmt.Printf(`DELETE FROM employee WHERE id = %v`, id)
}

/* End of Database Topic */

func EmployeeHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/employees/")
	id, _ := strconv.Atoi(sid) // Convert string to int

	switch {
	case r.Method == "GET" && id > 0:
		getEmployeeById(w, r, id)
	case r.Method == "GET":
		getEmployees(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found!")
	}
}

func getEmployeeById(w http.ResponseWriter, r *http.Request, id int) {
	db := createPGConnection()
	defer db.Close()
	var e employee
	db.QueryRow("SELECT id, name, email FROM employee WHERE id = $1", id).Scan(&e.ID, &e.Name, &e.Email)

	employeeJson, _ := json.Marshal(e)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(employeeJson))
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	db := createPGConnection()
	defer db.Close()
	rows, _ := db.Query("SELECT id, name, email FROM employee")

	var employees []employee
	for rows.Next() {
		var e employee
		rows.Scan(&e.ID, &e.Name, &e.Email)
		employees = append(employees, e)
	}

	employeesJson, _ := json.Marshal(employees)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(employeesJson))

	// db, err := sql.Open()
}
