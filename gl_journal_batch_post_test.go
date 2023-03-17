package sage300_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGLJournalBatchPost(t *testing.T) {
	req := client.NewGLJournalBatchPostRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
