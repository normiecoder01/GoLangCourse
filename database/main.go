package main

import (
	"fmt"
	"database/sql"
	// This is unnamed import becasuse this import functionality 
	// is going to be handled by the "database/sql" import
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"golang.org/x/crypto/bcrypt"

)

var schema =`
CREATE TABLE IF NOT EXISTS users(
id INTEGER PRIMARY KEY AUTOINCREMENT,
name TEXT NOT NULL,
email TEXT NOT NULL UNIQUE,
hashedPassword BLOB NOT NULL,
created at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`
func createUser(db *sql.DB, name , email , hashedPassword string ) (int64 , error) {
	stmt := `INSERT INTO users (name , email , hashedPassword) VALUES (? , ? , ?)`

	hp , err := bcrypt.GenerateFromPassword([]byte(hashedPassword) , bcrypt.DefaultCost )
	if err!=nil{
		log.Fatal(err)
		return 0 ,err
	}

	result , err := db.Exec(stmt , "Anupal", "anumail@gmail.com",string(hp))
	if err!=nil{
		log.Fatal(err)
		return 0 , err
	}

	return	result.LastInsertId()

}
func main(){
	// Here we are removing the file generated when the 
	// program was previously run because we want to start 
	// afresh with every new run
	dbName := "data.db"
	_ = os.Remove(dbName)

	db , error := sql.Open("sqlite3", dbName)
	if error != nil{
		log.Fatal(error)
	}
	defer func() {
		fmt.Println("Closing database conncetion.")
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	error = db.Ping()
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Database connection eastablished successfully!!")

	_ , err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
	
		fmt.Println("Database table was created successfully!!")

}