package main

import (
    "flag"
    "fmt"
    "github.com/darkhelmet/twanalyzer/tweets"
)

var (
    input = flag.String("input", "tweets.csv", "The location of your tweets.csv file")
)

func main() {
    flag.Parse()
    t, err := tweets.ParseTweets(*input)
    if err != nil {
        panic(err)
    }
    fmt.Println(t.Stats())
}
