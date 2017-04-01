package main

import (
	"net/http"
	"database/sql"
	"html/template"
	"github.com/julienschmidt/httprouter"
)

var db *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("templates/index.tmpl")
	q := r.URL.Query()
	st := q.Get("searchTerm")
	data := ReadTweets(db, st)
	t.Execute(w, map[string]interface{}{"data": data})
}
