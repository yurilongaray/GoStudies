package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid) // Convert string to int

	switch {
	case r.Method == "GET" && id > 0:
		getUserById(w, r, id)
	case r.Method == "GET":
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found!")
	}
}

func getUserById(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open()
}
