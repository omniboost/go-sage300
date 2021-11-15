package sage300_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestGLJournalBatchGet(t *testing.T) {
	req := client.NewGLJournalBatchGetRequest()
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	os.Exit(12)

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
