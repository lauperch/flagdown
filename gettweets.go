package main

import (
  "fmt"
  "github.com/cdipaolo/sentiment"
  "github.com/ChimeraCoder/anaconda"
)

func GetTweets() {
  anaconda.SetConsumerKey("")
  anaconda.SetConsumerSecret("")
  api := anaconda.NewTwitterApi("", "")

  model, _ := sentiment.Train()

  for true {
    fmt.Println("next iter")
    searchTerms := [...]string{"knope", "location", "place", "on the", "under", "address", "down"}
    for _, searchTerm := range searchTerms {
      searchResult, err := api.GetSearch(searchTerm, nil)
      if err != nil {
        panic(err)
      }
      for _, tweet := range searchResult.Statuses {
        if !tweet.HasCoordinates() {
          continue
        }
        fmt.Println("got coords")
        _, err := tweet.CreatedAtTime()
        if err != nil {
          panic(err)
        }
        a := model.SentimentAnalysis(tweet.Text, sentiment.English)
        sent := a.Score
        fmt.Println(sent)
        StoreTweet(db, tweet, sent)
      }
    }
  }
}
