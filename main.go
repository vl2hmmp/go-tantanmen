package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Menu struct {
	ID           int64  `json:"id"`
}

func getDB() (*sqlx.DB, error) {
	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		port = "3306"
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("failed to read DB port number from an environment variable MYSQL_PORT.\nError: %s", err.Error())
	}
	user := os.Getenv("MYSQL_USER")
	if user == "" {
		user = "tantanmen"
	}
	dbname := os.Getenv("MYSQL_DBNAME")
	if dbname == "" {
		dbname = "tantanmen"
	}
	password := os.Getenv("MYSQL_PASS")
	if password == "" {
		password = "tantanmen"
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	return sqlx.Open("mysql", dsn)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//menu := Menu{}
	//dbx, err := getDB()
	//if err != nil {
	//	log.Fatalf("failed to connect to DB: %s.", err.Error())
	//}
	//menuId := 1
	//err = dbx.Get(&menu,"SELECT * FROM `menus` WHERE `id` = ?", menuId)
	//if err != nil {
	//	log.Fatalf("failed to connect to DB: %s.", err.Error())
	//}

	fmt.Fprint(w, "test")
}

func main() {
	fmt.Println("run")
	http.HandleFunc("/", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
