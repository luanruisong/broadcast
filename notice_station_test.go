package signal_station

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNoticeStation(t *testing.T) {
	ns := NewNoticeStation()

	for i := 0; i < 10; i++ {
		go func(idx int) {
			for {
				fmt.Println(idx, ns.WaitForValue())
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		ns.Notice(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}

func TestOnNotice(t *testing.T) {
	ns := NewNoticeStation()

	for i := 0; i < 10; i++ {
		go func(idx int) {
			ns.OnNotice(context.Background(), func(value interface{}) {
				fmt.Println(idx, value)
			})
		}(i)
	}

	for i := 0; i < 10; i++ {
		ns.Notice(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}
