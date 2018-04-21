package path

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	vectors := []string{
		"",
		"M 10 20",
		"m 20 30",
		// TODO: "M 10 20 L 20 30",
	}
	for _, vector := range vectors {
		fmt.Println("d:", vector)
		d, err := Parse("grammar_test.go", []byte(vector), Recover(false))
		if err != nil {
			t.Error(err)
		}
		jsn, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			t.Error(err)
		}
		fmt.Println("JSON:", string(jsn))
	}
}
