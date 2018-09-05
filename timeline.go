package timeline

import (
	"sync"
	"time"
)

type Timeline struct {
	epoch  time.Time
	events []Event
	lock   sync.Mutex
}

func New() (tl *Timeline) {
	tl = new(Timeline)
	tl.epoch = time.Now()
	return
}

type Event struct {
	duration time.Duration
	callback Callback
}

type Callback func()

func (tl *Timeline) Append(callback Callback, triggerAt time.Time) (chain *Timeline) {
	var event Event
	event.callback = callback
	event.duration = triggerAt.Sub(tl.epoch)
	tl.events = append(tl.events, event)

	chain = tl
	return
}

func (tl *Timeline) AppendInterspersed(
	startAt time.Time, interval time.Duration, callbacks ...Callback,
) (chain *Timeline) {
	if len(callbacks) == 0 {
		return
	}

	tl.Append(func() {
		go callbacks[0]()
		tl.AppendInterspersed(
			startAt.Add(interval),
			interval,
			callbacks[1:]...,
		)
	}, startAt)

	chain = tl
	return
}

func (tl *Timeline) Start() {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		go func() {
			for i, e := range tl.events {
				func(i int, e Event) {
					if e.duration <= time.Now().Sub(tl.epoch) {
						go e.callback()
					}

					tl.lock.Lock()
					tl.events = append(tl.events[0:i], tl.events[i+1:]...)
					tl.lock.Unlock()
				}(i, e)
			}
		}()
	}
}
