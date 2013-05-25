package tweets

import (
    "encoding/csv"
    "fmt"
    "io"
    "os"
    "time"
    "unicode/utf8"
)

const (
    TimestampFormat   = "2006-01-02 15:04:05 -0700"
    GoldenTweetLength = 140
)

type ID uint64
type Timestamp time.Time

type Tweet struct {
    Id                                   ID
    InReplyToStatusId, InReplyToUserId   ID
    RetweetStatusId, RetweetStatusUserId ID
    Timestamp                            Timestamp
    Source                               string
    Text                                 string
    URLs                                 []string
}

func (t *Tweet) IsReply() bool {
    return t.InReplyToStatusId > 0
}

func (t *Tweet) IsRetweet() bool {
    return t.RetweetStatusId > 0
}

func (t *Tweet) IsGolden() bool {
    return utf8.RuneCountInString(t.Text) == GoldenTweetLength
}

func (t *Tweet) NumberOfUrls() int {
    return len(t.URLs)
}

type Stats struct {
    Total     int
    Golden    int
    Replies   int
    Retweets  int
    MostUrls  int
    TotalUrls int
}

func (s *Stats) String() string {
    format := "total: %d, golden: %d (%.2f%%), replies: %d (%.2f%%), retweets: %d (%.2f%%), total urls: %d, most urls in one tweet: %d"
    return fmt.Sprintf(format, s.Total, s.Golden, s.percent(s.Golden), s.Replies, s.percent(s.Replies), s.Retweets, s.percent(s.Retweets), s.TotalUrls, s.MostUrls)
}

func (s *Stats) percent(i int) float64 {
    return float64(i) / float64(s.Total) * 100
}

type Tweets []*Tweet

func (t Tweets) Len() int {
    return len(t)
}

func (t Tweets) Stats() *Stats {
    s := &Stats{Total: t.Len()}
    for _, tweet := range t {
        if tweet.IsGolden() {
            s.Golden++
        }

        if tweet.IsReply() {
            s.Replies++
        }

        if tweet.IsRetweet() {
            s.Retweets++
        }

        nurls := tweet.NumberOfUrls()
        s.TotalUrls += nurls
        if nurls > s.MostUrls {
            s.MostUrls = nurls
        }
    }
    return s
}

func (t Tweets) CountGolden() int {
    c := 0
    for _, tweet := range t {
        if tweet.IsGolden() {
            c++
        }
    }
    return c
}

func ParseTweets(path string) (Tweets, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    r := csv.NewReader(file)
    r.FieldsPerRecord = -1
    // The header
    row, err := r.Read()
    if err != nil {
        return nil, err
    }

    var t Tweets
    for {
        row, err = r.Read()

        if err == io.EOF {
            break
        }

        if err != nil {
            return nil, err
        }

        if len(row) > 0 {
            t = append(t, fromRow(row))
        }
    }

    return t, nil
}
