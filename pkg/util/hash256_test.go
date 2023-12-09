package util

import "testing"

func TestHash256(t *testing.T) {
	str := "go_reminder_bot"
	strHash := Hash256(str)
	t.Log(strHash)
}
