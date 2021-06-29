package accountviewnet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestPaymodeGroups(t *testing.T) {
	req := client.NewPaymodeGroupsRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
