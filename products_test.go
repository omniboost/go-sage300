package trivec_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestProducts(t *testing.T) {
	req := client.NewProductsRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
