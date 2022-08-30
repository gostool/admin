package tests

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Unix(time.Now().Unix(), 0)
	t.Log(now)
	const shortForm = "2006-Jan-02"
	ti, _ := time.Parse(shortForm, "2012-Feb-03")
	t.Log("time.Parse:", ti)

	timer, err := time.Parse("2006-01-02 15:04:05", "2018-08-08 08:08:08")
	if err != nil {
		fmt.Printf("parse time failed, err:%v\n", err)
		return
	}
	t.Logf("timeobj:%v\n", timer)
}
