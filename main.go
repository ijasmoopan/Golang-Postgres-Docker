package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connectDB()

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hi", hiHandler)
	http.ListenAndServe(":8080", nil)
}

func connectDB(){
	fmt.Println("Connecting..")
	// these details match the docker-compose.yml file.
	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "postgres", 5432, "postgres", "mypassword", "test")
	
	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	start := time.Now()
	for db.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			fmt.Println("Failed to connect after 10 secs.")
			break
		}
	}
	fmt.Println("Connected:", db.Ping() == nil)
	_, err = db.Exec(`DROP TABLE IF EXISTS COMPANY;`)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`CREATE TABLE COMPANY (ID INT PRIMARY KEY, NAME text);`)
	if err != nil {
		fmt.Println("Error while creating table:", err)
		panic(err)
	}
	fmt.Println("Table company is created..")
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Hello Handler")
	fmt.Fprintf(w, "Hello World!")
}

func hiHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Hiii Handler")
	fmt.Fprintf(w, "Hiii World!")
}