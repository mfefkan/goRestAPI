package main

import (
	"fmt"
	"test/db"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		fmt.Println("Database connection error:", err)
	}

	fmt.Println(database.Config)

}
