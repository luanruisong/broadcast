package signal_station

import (
	"context"
	"sync"
)

type (
	NoticeStation struct {
		lock   sync.RWMutex
		signal *SignalStation
		curr   interface{}
	}
)

func (ss *NoticeStation) WaitForValue() interface{} {
	<-ss.CurrSignal()
	return ss.CurrValue()
}

func (ss *NoticeStation) CurrSignal() <-chan struct{} {
	return ss.signal.CurrSignal()
}

func (ss *NoticeStation) CurrValue() interface{} {
	ss.lock.RLock()
	defer ss.lock.RUnlock()
	return ss.curr
}

func (ss *NoticeStation) OnNotice(ctx context.Context, f func(value interface{})) {
	ss.signal.OnSignal(ctx, func() {
		f(ss.CurrValue())
	})
}

func (ss *NoticeStation) OnNoticeAsync(ctx context.Context, f func(value interface{})) {
	go ss.OnNoticeAsync(ctx, f)
}

func (ss *NoticeStation) Notice(value interface{}) {
	ss.lock.Lock()
	defer ss.lock.Unlock()
	ss.curr = value
	ss.signal.Send()
}

func NewNoticeStation() *NoticeStation {
	return &NoticeStation{
		signal: NewSignalStation(),
		lock:   sync.RWMutex{},
	}
}
