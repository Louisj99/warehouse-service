package main

import (
	"fmt"
	"log"
	"warehouse-service/pkg/adapters"
	"warehouse-service/pkg/drivers"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "12345"
	dbname   = "postgres"
)

func main() {
	// setup postgresql adapter
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
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
