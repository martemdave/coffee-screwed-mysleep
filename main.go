package main

import (
	// "RealDeal/ConnectingDB/ExecuteSelectQuery.go"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	// "time"
)

func main() {
	log.Println("starting program")
	// get the database connection URL.
	// usually, this is taken as an environment variable as in below commented out code
	// databaseUrl = os.Getenv("DATABASE_URL")
	// for the time being, let's hard code it as follows. change the values as needed.
	// databaseUrl := "postgres://postgres:mypassword@localhost:5432/postgres"
	databaseUrl := "postgres://nftart:secret@localhost:5432/postgres"
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//to close DB pool
	defer dbPool.Close()

	// ExecuteSelectQuery(dbPool)
	// ExecuteFunction(dbPool)
	log.Println("stopping program")
}

// func ExecuteSelectQuery(dbPool *pgxpool.Pool) {
// 	log.Println("starting execution of select query")
// 	//execute the query and get result rows
// 	rows, err := dbPool.Query(context.Background(), "select * from public.person")
// 	if err != nil {
// 		log.Fatal("error while executing query")
// 	}

// 	log.Println("result:")
// 	//iterate through the rows
// 	for rows.Next() {
// 		values, err := rows.Values()
// 		if err != nil {
// 			log.Fatal("error while iterating dataset")
// 		}
// 		//convert DB types to Go types
// 		id := values[0].(int32)
// 		firstName := values[1].(string)
// 		lastName := values[2].(string)
// 		dateOfBirth := values[3].(time.Time)
// 		log.Println("[id:", id, ", first_name:", firstName, ", last_name:", lastName, ", date_of_birth:", dateOfBirth, "]")
// 	}

// }

// func ExecuteFunction(dbPool *pgxpool.Pool) {
// 	log.Println("starting execution of databse function")
// 	// id can be taken as a user input
// 	// for the time being, let's hard code it
// 	id := 1

// 	//execute the query and get result rows
// 	rows, err := dbPool.Query(context.Background(), "select * from public.get_person_details($1)", id)
// 	log.Println("input id: ", id)
// 	if err != nil {
// 		log.Fatal("error while executing query")
// 	}

// 	log.Println("result:")
// 	//iterate through the rows
// 	for rows.Next() {
// 		values, err := rows.Values()
// 		if err != nil {
// 			log.Fatal("error while iterating dataset")
// 		}

// 		//convert DB types to Go types
// 		firstName := values[0].(string)
// 		lastName := values[1].(string)
// 		dateOfBirth := values[2].(time.Time)

// 		log.Println("[first_name:", firstName, ", last_name:", lastName, ", date_of_birth:", dateOfBirth, "]")
// 	}

// }

// package main

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"net/http"
// 	"os"

// 	// "database/sql"
// 	// _ "github.com/go-sql-driver/mysql"
// 	"github.com/gorilla/mux"
// 	"github.com/jackc/pgx/v4"
// 	"github.com/lib/pq"
// )

// func CreateApplication(w http.ResponseWriter, r *http.Request) {}

// func main() {
// 	conn, err := pgx.Connect(context.Background(), os.Getenv("http://localhost:5432/"))
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close(context.Background())

// 	var greeting string
// 	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(greeting)

// 	router := mux.NewRouter()

// 	router.HandleFunc("api/v0/application/add", CreateApplication).Methods("POST")
// 	router.HandleFunc("/api/user/login").Methods("GET")
// 	// router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
// 	// router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts

// 	router.Use(app.JwtAuthentication) //attach JWT auth middleware

// 	//router.NotFoundHandler = app.NotFoundHandler

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8000" //localhost
// 	}

// 	fmt.Println(port)

// 	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	db.Exec(fmt.Sprintf("CREATE TABLE %s", pq.QuoteIdentifier(table)))
// 	sql.DB

// }
