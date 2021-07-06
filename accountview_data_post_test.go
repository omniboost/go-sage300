package accountviewnet_test

import (
	"encoding/json"
	"log"
	"testing"

	accountviewnet "github.com/omniboost/go-accountview.net"
)

func TestAccountviewDataPost(t *testing.T) {
	req := client.NewAccountviewDataPostRequest()
	req.SetRequestBody(accountviewnet.AccountviewDataPostRequestBody{
		BookDate:       "2021-7-2T10:39:05.276Z",
		BusinessObject: "DJ2",
		Table: accountviewnet.Table{
			Definition: accountviewnet.TableDefinition{
				Name: "DJ_PAGE",
				Fields: accountviewnet.TableDefinitionFields{
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
			DetailDefinitions: accountviewnet.TableDetailDefinitions{
				{
					Name: "DJ_LINE",
					Fields: accountviewnet.TableDefinitionFields{
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
							Name:      "SRC_TYPE",
							FieldType: "N",
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
		TableData: accountviewnet.TableData{
			Data: accountviewnet.TableDataData{
				Rows: accountviewnet.Rows{
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
			DetailData: accountviewnet.DetailData{
				{
					Rows: accountviewnet.Rows{
						{
							Values: []interface{}{
								"1",
								"1",
								"TEST",
								"1337",
								"TEST",
								"TEST",
								"2016-11-30T13:37:37.450Z",
								"TEST",
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
