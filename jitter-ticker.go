package time

import (
	"errors"
	"math/rand"
	"time"
)

// JitterTicker has some probalistic component to it
type JitterTicker struct {
	C  chan time.Time
	R  *rand.Rand
	on bool
}

// Stop stops the ticker
func (self *JitterTicker) Stop() {
	self.on = false
}

// NewJitterTicker returns
func NewJitterTicker(low, high time.Duration) *JitterTicker {

	if low > high {
		panic(errors.New("low must be a shorter duration than high"))
	}

	s := rand.NewSource(time.Now().Unix())

	ticker := JitterTicker{
		C:  make(chan time.Time, 1),
		R:  rand.New(s),
		on: true,
	}

	go func() {
		for ticker.on {
			jitter := ticker.R.Uint64() % uint64(high-low)
			t := time.NewTimer(low + time.Duration(jitter))
			ticker.C <- <-t.C
			if t.Stop() {
				<-t.C
			}
		}
	}()

	return &ticker
}
