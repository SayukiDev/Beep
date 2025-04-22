package Beep

import (
	"testing"
)

func TestPlay(t *testing.T) {
	done := make(chan bool)
	err := PlayFromPath("./test_data/1.mp3", func() {
		done <- true
	})
	if err != nil {
		t.Error(err)
		return
	}
	<-done
}
