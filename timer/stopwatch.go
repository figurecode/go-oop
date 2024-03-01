package timer

import "time"

type Stopwatch struct {
	start  time.Time
	splits []time.Time
}

func (s *Stopwatch) Start() {
	s.start = time.Now()
	s.splits = nil
}

func (s *Stopwatch) SaveSplit() {
	s.splits = append(s.splits, time.Now())
}

func (s Stopwatch) GetResults() (retResult []time.Duration) {
	for _, splitTime := range s.splits {
		retResult = append(retResult, splitTime.Sub(s.start))
	}

	return
}
