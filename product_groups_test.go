package accountviewnet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestProductGroups(t *testing.T) {
	req := client.NewProductGroupsRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
