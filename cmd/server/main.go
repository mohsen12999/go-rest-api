package main

import (
	"fmt"

	"github.com/mohsen12999/go-rest-api/internal/db"
)

// Run - is going to be responsible for the instantiation and startup of our go application
func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("successfully connect and pinged database")

	return nil
}

func main() {
	fmt.Println("Go REST API!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
