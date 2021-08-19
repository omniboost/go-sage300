package netsuite_test

import (
	"encoding/json"
	"log"
	"testing"

	netsuite "github.com/omniboost/go-netsuite"
)

func TestDjPageTypeTest(t *testing.T) {
	o := netsuite.DjPage{
		DjCode:  "860",
		HdrDesc: "TEST",
	}
	lines := []netsuite.DjLine{
		{
			AcctNr: "9999",
			Amount: 12.1,
			RecID:  "REC_ID",
		},
		{
			AcctNr: "9999",
			Amount: -12.1,
			RecID:  "REC_ID",
		},
	}
	req, err := o.ToNetsuiteDataPostRequest(client, lines)
	if err != nil {
		t.Error(err)
		return
	}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
