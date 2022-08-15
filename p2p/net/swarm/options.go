package swarm

import "time"

type Options struct {
	dialDelay time.Duration
}

var (
	options = &Options{
		dialDelay: time.Second * 3,
	}
)

func WithDialDelay(delay time.Duration) {
	options.dialDelay = delay
}
