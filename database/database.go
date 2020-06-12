package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "hackathon_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		" dbname=%s sslmode=disable",
		host, port, user, dbname)
	_, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Postgres Database has been setup successfully")
}
