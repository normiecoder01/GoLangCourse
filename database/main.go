package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	// This is unnamed import becasuse this import functionality
	// is going to be handled by the "database/sql" import
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
}

type UserWithoutPassword struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var schema = `
CREATE TABLE IF NOT EXISTS users(
id INTEGER PRIMARY KEY AUTOINCREMENT,
name TEXT NOT NULL,
email TEXT NOT NULL UNIQUE,
hashedPassword BLOB NOT NULL,
created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

func createSchema(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database table was created successfully!!")
}
func createUser(db *sql.DB, name, email, hashedPassword string) (int64, error) {
	// Writing 'VALUES (? , ? , ?)' instead of adding the variables using printf is very crucial to avoid SQL Injection.
	stmt := `INSERT INTO users (name , email , hashedPassword) VALUES (? , ? , ?)`

	hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	result, err := db.Exec(stmt, name, email, string(hp))
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return result.LastInsertId()

}

func createUserWithPrepStmt(db *sql.DB, name, email, hashedPassword string) (int64, error) {
	// Replace the string with db.Prepare(<string>) and save into a variable called stmt
	stmt, err :=db.Prepare(`INSERT INTO users (name , email , hashedPassword) VALUES (? , ? , ?)`)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}	

	hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	result, err := stmt.Exec(name, email, string(hp))
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return result.LastInsertId()

}
func fetchUserByEmail(db *sql.DB, email string) (*User, error) {
	stmt := `SELECT id , name,  email , hashedPassword, created_at FROM users WHERE email = ?`

	row := db.QueryRow(stmt, email)

	var user User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
func fetchUsers(db *sql.DB) ([]UserWithoutPassword, error) {

	stmt := `SELECT id, name, email, created_at FROM users`

	rows, err := db.Query(stmt)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()

	var users []UserWithoutPassword
	var user UserWithoutPassword		
	for rows.Next() {
		
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}

func main() {
	// Here we are removing the file generated when the
	// program was previously run because we want to start
	// afresh with every new run
	dbName := "data.db"
	_ = os.Remove(dbName)

	db, error := sql.Open("sqlite3", dbName)
	if error != nil {
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

	createSchema(db)

	result, error := createUser(db, "Anupal Umale", "normiecoder@gmail.com", "Anupal@123")
	error = db.Ping()
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(result)

	result, error = createUserWithPrepStmt(db, "Ashutosh Paliwal", "ashupal@gmail.com", "ashupal@123")
	error = db.Ping()
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(result)

	anupal, err := fetchUserByEmail(db, "normiecoder@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(anupal)

	bs, err := json.MarshalIndent(anupal, "", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))

	users, err := fetchUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	bs, err = json.MarshalIndent(users, "", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))

}
