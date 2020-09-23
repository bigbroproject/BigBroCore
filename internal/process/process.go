package process

import (
	"time"
)

type Process struct {
	Function func()
	Channel  chan string
}

func NewProcess(function func(), channel chan string) *Process {
	return &Process{Function: function, Channel: channel}
}

func ScheduleProcess(process *Process, ms int64) {
	ticker := time.NewTicker(time.Duration(ms) * time.Millisecond)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				process.Function()
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()
}
