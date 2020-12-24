package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GoodWatchesMovie struct {
	Title  string `json:"Title"`
	Rating int    `json:5`
}

func loadDatabase() *sql.DB {
	// godotenv.Load(".env")
	// username := os.Getenv("db_username")
	// password := os.Getenv("db_password")
	// address := os.Getenv("db_hostname")
	// dbName := os.Getenv("db_name")
	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
	// 	username,
	// 	password,
	// 	address,
	// 	dbName)
	// fmt.Println(username, password, address, dataSourceName)
	// fmt.Println("attempting to load database...")
	// database, err := sql.Open("mysql", dataSourceName)
	// // if err != nil {
	// // 	panic(err)
	// // }
	// fmt.Println("database successfully loaded!")
	return nil
}

func handler(writer http.ResponseWriter, request *http.Request) {
	sample := GoodWatchesMovie{Title: "Hello World", Rating: 5}
	fmt.Println("movies endpoint hit")
	json.NewEncoder(writer).Encode(sample)
}

func main() {
	// database := loadDatabase()
	http.HandleFunc("/movies", handler)
	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
