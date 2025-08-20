package pomodoro

import (
	"sync"
	"time"
)

type Pomodoro struct {
	Duration  time.Duration
	StartTime time.Time
	IsRunning bool
	clock     sync.Mutex
}

func New() *Pomodoro {
	return &Pomodoro{}
}

func (p *Pomodoro) Start(duration time.Duration) {
	p.clock.Lock()
	defer p.clock.Unlock()

	p.Duration = duration
	p.StartTime = time.Now()
	p.IsRunning = true
}

func (p *Pomodoro) Status() (time.Duration, bool) {
	p.clock.Lock()
	defer p.clock.Unlock()

	if !p.IsRunning {
		return 0, false
	}

	space := time.Since(p.StartTime)
	remaining := p.Duration - space
	if remaining <= 0 {
		remaining = 0
		p.IsRunning = false
	}

	return remaining, p.IsRunning
}
