package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type users struct {
	id    int
	name  string
	email string
	phone string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result []users

	rows, err := db.Query("SELECT * FROM users WHERE name like 'rhe%'")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var each = users{}
		var err = rows.Scan(&each.id, &each.name, &each.email, &each.phone)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	json.NewEncoder(w).Encode(result)
}

// our main function
func main() {
	//usersJson, err := json.Marshal(users)
	router := mux.NewRouter()
	router.HandleFunc("/users", GetUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
