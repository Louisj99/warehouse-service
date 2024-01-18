package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"warehouse-service/pkg/adapters"
	"warehouse-service/pkg/drivers"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	postgresAdapter, err := adapters.NewPostgresAdapter(psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	r := drivers.SetupRouter(postgresAdapter)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
