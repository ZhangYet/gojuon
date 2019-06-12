package common

import (
	"fmt"
	"testing"
)

func TestQueryKanaRows(t *testing.T) {
	testKey := []string{"k", "ka", "か", "カ"}
	for _, key := range testKey {
		testName := fmt.Sprintf("testing key %s", key)
		t.Run(testName, func(t *testing.T) {
			ret := QueryKanaRows("hira", key)
			if len(ret) <= 0 {
				t.Fatal("cannot get gana from data")
			}
			t.Log(ret)
		})
	}
}
