package sage300_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestPostingJournalDetailsGet(t *testing.T) {
	req := client.NewPostingJournalDetailsGetRequest()
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
