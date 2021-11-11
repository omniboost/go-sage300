package sage300_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestTaxGroupsGet(t *testing.T) {
	req := client.NewTaxGroupsGetRequest()
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
