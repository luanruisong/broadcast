package signal_station

import (
	"context"
	"sync"
)

type (
	SignalStation struct {
		lock sync.RWMutex
		cp   chan struct{}
	}
)

func (ss *SignalStation) CurrSignal() <-chan struct{} {
	ss.lock.RLock()
	defer ss.lock.RUnlock()
	return ss.cp
}

func (ss *SignalStation) OnSignal(ctx context.Context, f func()) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-ss.CurrSignal():
			f()
		}
	}
}

func (ss *SignalStation) Send() {
	ss.lock.Lock()
	defer ss.lock.Unlock()
	cc := ss.cp
	ss.cp = make(chan struct{})
	close(cc)
}

func NewSignalStation() *SignalStation {
	return &SignalStation{
		lock: sync.RWMutex{},
		cp:   make(chan struct{}),
	}
}
