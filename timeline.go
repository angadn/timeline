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
	isDone   bool
}

type Callback func()

func (tl *Timeline) Append(callback Callback, triggerAt time.Time) (chain *Timeline) {
	tl.lock.Lock()
	defer tl.lock.Unlock()

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
					if !e.isDone && e.duration <= time.Now().Sub(tl.epoch) {
						go e.callback()
						tl.events[i].isDone = true
					}
				}(i, e)
			}

			tl.lock.Lock()
			defer tl.lock.Unlock()

			events := []Event{}
			for _, e := range tl.events {
				if !e.isDone {
					events = append(events, e)
				}
			}
			tl.events = events
		}()
	}
}
