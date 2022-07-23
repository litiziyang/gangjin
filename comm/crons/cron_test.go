package crons

import (
	"testing"
	"time"
)

func Test_cron(t *testing.T) {
	ts := 3
	InserTimerTask("ss", 2*time.Minute, func(id string, err error) error {
		td := ts + 1
		t.Log(td)
		return nil
	})
}

func Test_te(t *testing.T) {
	err := InitProcessTimer()
	if err != nil {
		t.Error(err)
	}
}
