package accountviewnet_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	accountviewnet "github.com/omniboost/go-accountview.new"
)

func TestTickets(t *testing.T) {
	req := client.NewTicketsRequest()
	req.PathParams().Date = accountviewnet.Date{time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
