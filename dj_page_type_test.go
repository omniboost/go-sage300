package accountviewnet_test

import (
	"encoding/json"
	"log"
	"testing"

	accountviewnet "github.com/omniboost/go-accountview.net"
)

func TestDjPageTypeTest(t *testing.T) {
	o := accountviewnet.DjPage{
		DjCode:  "860",
		HdrDesc: "TEST",
	}
	req, err := o.ToAccountviewDataPostRequest(client)
	if err != nil {
		t.Error(err)
	}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
