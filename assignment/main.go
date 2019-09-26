package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Shared struct {
	mu    sync.RWMutex
	cache map[int64]bool
}

var (
	shared    Shared
	db        *sql.DB
	id        int64
	workersno = flag.Int("n", 30, "The number of workers to start") //no of workers
	queuesize = flag.Int("s", 15000, "The size of Workerqueue ")    //size of queue
)

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func initDB() { //connecting to database
	config := dbConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string { //configuring database
	conf := make(map[string]string)
	conf[dbhost] = "localhost"
	conf[dbport] = "5432"
	conf[dbuser] = "postgres"
	conf[dbpass] = "password"
	conf[dbname] = "mydb"
	return conf
}

func getcount(count *int) { //getting count of rows from database
	sqlStatement := `SELECT COUNT(*) from locale1`
	row, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("error at count no of rows", err.Error())
	}
	if row.Next() {
		err1 := row.Scan(count)
		if err1 != nil {
			fmt.Println("error at inserting count", err1)
		}
	}
}
func getid() { //getting max id from database
	sqlStatement := `SELECT MAX(id) from locale1`
	row, err := db.Query(sqlStatement)

	var count int
	getcount(&count)

	if count == 0 {
		id = 0
		fmt.Println("Empty result")
		return
	}

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		for row.Next() {
			err1 := row.Scan(&id)
			if err1 != nil {
				fmt.Println(err1)
			}
		}
	default:
		fmt.Println(err)
	}

}
func initmem() { //initializing shared memory

	shared.cache = make(map[int64]bool)
}
func Index(w http.ResponseWriter, r *http.Request) { //home api

	fmt.Fprintf(w, "Home")
}

func main() {

	flag.Parse()                //parse flags
	initDB()                    //create DB connections
	getid()                     //getting the lastinserted ID from database
	initmem()                   //initialize shared memnory
	StartDispatcher(*workersno) //Starting dispatcher with workers = NWorkers

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index).Methods("GET")          //Home page
	router.HandleFunc("/send", Collector).Methods("POST") //API to insert data in postgres
	err := http.ListenAndServe(":8080", router)           //Listener
	log.Fatal(err)
}
