package request

import (
	"encoding/json"
	"log"
	"time"
)

type CaughtRequest struct {
	Id            int64
	Time          time.Time `db:"created_at"`
	Method        string    `db:"method"`
	ContentLength int64     `db:"content_length"`
	RemoteAddr    string    `db:"remote_addr"`
	Url           string    `db:"url"`
	Headers       string    `db:"headers"`
	Body          string    `db:"body"`
}

func (v *CaughtRequest) ParsedHeaders() map[string][]string {
	var headers map[string][]string
	err := json.Unmarshal([]byte(v.Headers), &headers)
	if err != nil {
		log.Print("Failed to unmarshal headers json", err)
	}
	return headers
}
