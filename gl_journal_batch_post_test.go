package sage300_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGLJournalBatchPost(t *testing.T) {
	req := client.NewGLJournalBatchPostRequest()
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
