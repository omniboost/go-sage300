package trivec_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	trivec "github.com/omniboost/go-trivec"
)

func TestTickets(t *testing.T) {
	req := client.NewTicketsRequest()
	req.PathParams().Date = trivec.Date{time.Date(2021, 4, 12, 0, 0, 0, 0, time.UTC)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
