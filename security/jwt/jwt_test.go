package jwt

import (
	"fmt"
	"testing"
)

func TestBuildToken(t *testing.T) {
	tokenString, err := ParseUser("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjozfSwiZXhwIjoxNjU1NjM2NDI2LCJpc3MiOiJpdC1pcy1nYW5namluIn0.4nCjyrwvS4tOYSc5niCrfTzEEJB6Opz_cPscFNO-3o0")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(tokenString)
}
