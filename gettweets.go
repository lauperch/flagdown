package main

import (
  "fmt"
  "github.com/cdipaolo/sentiment"
  "github.com/ChimeraCoder/anaconda"
)

func GetTweets() {
  anaconda.SetConsumerKey("86xjjpfHNFinjo2FcUF0sP0SH")
  anaconda.SetConsumerSecret("k1mxnOinS9CANchVRPt3sHZkNRajjh2bfMGFwDOS4pjOPRNlzY")
  api := anaconda.NewTwitterApi("62061676-OuYLpxWpmuRlLooOWMDdmh4Tai4Nu14xyBY3qFLEs", "RTMLo3f2f7uhT1rwziKjywJPabAlAUL4N6uBNa3yirG18")

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
