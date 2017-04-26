package main

import(
  "database/sql"
  "time"
  "strings"
  //"fmt"
  m "github.com/ChimeraCoder/anaconda"
  _ "github.com/mattn/go-sqlite3"
)

type TweetAbstr struct {
  Id        int64
  Tid       int64
  Author    string
  Body      string
  Sent      uint8
  Lng       float64
  Lat       float64
  Timestamp time.Time
}

func initDB(filepath string) *sql.DB {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if database == nil {
		panic("db nil")
	}
	return database
}

func CreateTable(db *sql.DB) {
  tweetsTable := `
  CREATE TABLE IF NOT EXISTS tweets(
    id INTEGER NOT NULL PRIMARY KEY,
    tid INTEGER,
    author  TEXT,
    body TEXT,
    sent INTEGER,
    lng FLOAT,
    lat FLOAT,
    timestamp DATETIME, UNIQUE(tid)
    );
    `
	_, err := db.Exec(tweetsTable)
	if err != nil {
		panic(err)
	}
}

func StoreTweet(db *sql.DB, tweet m.Tweet, sent uint8) {
  dbAddTweet := `
    INSERT OR IGNORE INTO tweets(
      tid,
      author,
      body,
      sent,
      lng,
      lat,
      timestamp
    ) VALUES(?, ?, ?, ?, ?, ?, ?)
  `
  stmt, err := db.Prepare(dbAddTweet)
  if err != nil {
    panic(err)
  }
  defer stmt.Close()

  lng, _ := tweet.Longitude()
  lat, _ := tweet.Latitude()

  createtime, err2 := tweet.CreatedAtTime()
  if err2 != nil {
    panic(err2)
  }

  _, err3 := stmt.Exec(tweet.Id, tweet.User.Name, tweet.Text, sent, lng, lat, createtime)
  if err3 != nil {
    panic(err3)
  }
}

func ReadTweets(db *sql.DB, st string) []TweetAbstr {
	var result []TweetAbstr
  sqlRead := `
  SELECT * FROM tweets WHERE datetime(timestamp) >= DATE('now') INTERVAL 1 DAY;
  `

  rows, err := db.Query(sqlRead)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		t := TweetAbstr{}
		_ = rows.Scan(&t.Id, &t.Tid, &t.Author, &t.Body, &t.Sent, &t.Lng, &t.Lat, &t.Timestamp)
    if (strings.Contains(t.Body, st) || st == "") {
		    result = append(result, t)
    }
	}
	return result
}
