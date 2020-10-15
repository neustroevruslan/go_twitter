package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port :8111")

	dbConnect()
	loadRoutes()
}

func loadRoutes() {
	router := MakeRoutes()

	err := http.ListenAndServe(":8111", router)

	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

func dbConnect() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_twitter")

	if err != nil {
		fmt.Println("db connect error: ", err)
	}

	defer func() {
		err = db.Close()

		if err != nil {
			fmt.Println("db close error: ", err)
		}
	}()
}