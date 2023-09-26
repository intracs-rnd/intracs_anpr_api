package main

import (
	"fmt"
	"net/http"

	// Third party
	_ "github.com/go-sql-driver/mysql"

	// Project
	"intracs_anpr_api/api"
)

func main() {

	http.HandleFunc("/captured", api.InsertCapture)

	// handlers.SqlQuery()
	// handlers.SqlExec()

	fmt.Println("Starting API server at http://localhost:8881/")
	http.ListenAndServe(":8881", nil)
}
