package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ExecuteFunction(dbPool *pgxpool.Pool) {
	log.Println("starting execution of databse function")
	// id can be taken as a user input
	// for the time being, let's hard code it
	id := 1

	//execute the query and get result rows
	rows, err := dbPool.Query(context.Background(), "select * from public.get_person_details($1)", id)
	log.Println("input id: ", id)
	if err != nil {
		log.Fatal("error while executing query")
	}

	log.Println("result:")
	//iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		//convert DB types to Go types
		firstName := values[0].(string)
		lastName := values[1].(string)
		dateOfBirth := values[2].(time.Time)

		log.Println("[first_name:", firstName, ", last_name:", lastName, ", date_of_birth:", dateOfBirth, "]")
	}

}
