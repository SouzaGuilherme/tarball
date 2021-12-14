package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"net/http"
)

func conectDB()(db *sql.DB){
	db, err := sql.Open("mysql",
		"MainUser:MainPassword@tcp(127.0.0.1:3306)/Worldcities")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func AllCities(w http.ResponseWriter, r *http.Request) {
	var db = conectDB();

	var name string

	rows, err := db.Query("SELECT city FROM mytable;")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
	}
	
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func main() {
	log.Println("Server started on: http://localhost:9000")

	http.HandleFunc("/", AllCities)

	http.ListenAndServe(":9000", nil)
}