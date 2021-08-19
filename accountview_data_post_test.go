package netsuite_test

import (
	"encoding/json"
	"log"
	"testing"

	netsuite "github.com/omniboost/go-netsuite"
)

func TestNetsuiteDataPost(t *testing.T) {
	req := client.NewNetsuiteDataPostRequest()
	req.SetRequestBody(netsuite.NetsuiteDataPostRequestBody{
		BookDate:       "2021-07-02T10:39:05.276Z",
		BusinessObject: "DJ2",
		Table: netsuite.Table{
			Definition: netsuite.TableDefinition{
				Name: "DJ_PAGE",
				Fields: netsuite.TableDefinitionFields{
					{
						Name:      "RowId",
						FieldType: "C",
					},
					{
						Name:      "PAGE_NR",
						FieldType: "C",
					},
					{
						Name:      "DJ_CODE",
						FieldType: "C",
					},
					{
						Name:      "HDR_DESC",
						FieldType: "C",
					},
				},
			},
			DetailDefinitions: netsuite.TableDetailDefinitions{
				{
					Name: "DJ_LINE",
					Fields: netsuite.TableDefinitionFields{
						{
							Name:      "RowId",
							FieldType: "C",
						},
						{
							Name:      "HeaderId",
							FieldType: "C",
						},
						{
							Name:      "PAGE_NR",
							FieldType: "C",
						},
						{
							Name:      "DJ_CODE",
							FieldType: "C",
						},
						{
							Name:      "SRC_ID",
							FieldType: "C",
						},
						{
							Name:      "TRN_DATE",
							FieldType: "T",
						},
						{
							Name:      "ACCT_NR",
							FieldType: "C",
						},
						{
							Name:      "REC_ID",
							FieldType: "C",
						},
					},
				},
			},
		},
		TableData: netsuite.TableData{
			Data: netsuite.TableDataData{
				Rows: netsuite.Rows{
					{
						Values: []interface{}{
							"1",
							"TEST",
							"860",
							"Test",
						},
					},
				},
			},
			DetailData: netsuite.DetailData{
				{
					Rows: netsuite.Rows{
						{
							Values: []interface{}{
								"1",
								"1",
								"TEST",
								"TEST",
								"TEST",
								"2021-11-30T13:37:37.450Z",
								"3811",
								"TEST",
							},
						},
					},
				},
			},
		},
	})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
