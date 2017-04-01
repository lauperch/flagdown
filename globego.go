package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func main() {
	const dbpath = "globego.db"
	db = initDB(dbpath)
	defer db.Close()
	CreateTable(db)


	go GetTweets()

	router := httprouter.New()
	http.Handle("/img", http.FileServer(http.Dir("./templates/img")))
	router.GET("/", IndexHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
