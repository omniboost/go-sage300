package sage300_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGLAccountsGet(t *testing.T) {
	req := client.NewGLAccountsGetRequest()
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
