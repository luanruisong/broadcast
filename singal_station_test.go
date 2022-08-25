package signal_station

import (
	"context"
	"testing"
	"time"
)

func TestSignalStation(t *testing.T) {

	ss := NewSignalStation()
	for i := 0; i < 10; i++ {
		go func(idx int) {
			for {
				select {
				case <-ss.CurrSignal():
					t.Log(idx, "receive signal")
				}
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		ss.Send()
		time.Sleep(time.Second)
	}
}

func TestOnSignal(t *testing.T) {

	ss := NewSignalStation()
	for i := 0; i < 10; i++ {
		idx := i
		go ss.OnSignal(context.TODO(), func() {
			t.Log(idx, "receive signal")
		})
	}
	for i := 0; i < 10; i++ {
		ss.Send()
		time.Sleep(time.Second)
	}
}
