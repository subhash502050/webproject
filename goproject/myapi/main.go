package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database connection string (username:password@protocol(address)/dbname)
	// Change these values to your MySQL credentials and database name
	dsn := "root:Hello12345@tcp(localhost:3306)/mydb"

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	defer db.Close() // Ensure the database connection is closed when main exits

	// Verify the connection is valid by pinging the database
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	fmt.Println("Successfully connected to the database!")

	// Example: Fetching data from a table (replace with your actual table and columns)
	rows, err := db.Query("SELECT id, name FROM blog_articleuser")
	if err != nil {
		log.Fatal("Error querying database:", err)
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal("Error scanning row:", err)
		}
		fmt.Printf("User ID: %d, Name: %s\n", id, name)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		log.Fatal("Error during row iteration:", err)
	}
}
