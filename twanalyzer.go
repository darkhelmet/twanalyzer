package main

import (
    "flag"
    "github.com/darkhelmet/twanalyzer/tweets"
    "log"
)

var (
    input = flag.String("input", "tweets.csv", "The location of your tweets.csv file")
)

func main() {
    flag.Parse()
    log.Printf("loading from %s", *input)
    t, err := tweets.ParseTweets(*input)
    if err != nil {
        log.Fatalf("parsing failed: %s", err)
    }
    log.Println(t.Stats())
}
