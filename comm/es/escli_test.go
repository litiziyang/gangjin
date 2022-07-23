package es

import "testing"

func Test_escli(t *testing.T) {
	escli, err := GetESClient()
	if err != nil {
		t.Error(err)
	}

}
