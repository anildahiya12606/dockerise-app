package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	startServer()
}

func startServer() {
	http.HandleFunc("/app", handler)
	http.ListenAndServe(":5001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	handleMysqlDrivers()
	connectToRedis()
	fmt.Fprintf(w, "Welcome to my world!!")
	fmt.Println("Hit home page")
}

func handleMysqlDrivers() *sql.DB {

	//sql.Open("mysql","User:Pass@tcp(127.0.0.0:3306)/test")
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:whatever@tcp(mysql-server:3306)/test")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("Failed to ping")
		panic(err.Error())
	}

	// _, err = db.Query("Create table test( id int primary key,value varchar(200))")
	// if err != nil {
	// 	fmt.Println("Error in table creation")
	// 	panic(err.Error())
	// }

	_, err = db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")
	if err != nil {
		fmt.Println("Error in insert")
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	return db
}

func connectToRedis() {

	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "redis-server:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()

	fmt.Println(pong, err)
}
