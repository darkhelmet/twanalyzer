package tweets

import (
    "strconv"
    "time"
)

func parseUint(s string) ID {
    if s == "" {
        return ID(0)
    }

    u, err := strconv.ParseUint(s, 10, 64)
    if err != nil {
        panic(err) // Should never happen
    }
    return ID(u)
}

func parseTimestamp(s string) Timestamp {
    t, err := time.Parse(TimestampFormat, s)
    if err != nil {
        panic(err) // Should never happen
    }
    return Timestamp(t)
}

func fromRow(row []string) *Tweet {
    return &Tweet{
        Id:                  parseUint(row[0]),
        InReplyToStatusId:   parseUint(row[1]),
        InReplyToUserId:     parseUint(row[2]),
        RetweetStatusId:     parseUint(row[3]),
        RetweetStatusUserId: parseUint(row[4]),
        Timestamp:           parseTimestamp(row[5]),
        Source:              row[6],
        Text:                row[7],
        URLs:                row[8:],
    }
}
