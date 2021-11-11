package sage300_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestPostedTransactionsGet(t *testing.T) {
	req := client.NewPostedTransactionsGetRequest()
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
