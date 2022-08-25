package tests

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Unix(time.Now().Unix(), 0)
	t.Log(now)
}
