package xstatistics

import (
	"fmt"
	"sync"
	"time"
)

type StatisticsItem struct {
	RequestTimes  int
	ResponseTimes int
	TotalDuration time.Duration
	MinDuration   time.Duration
	MaxDuration   time.Duration
	LastDuration  time.Duration
}

type StatisticsMap struct {
	mutex     sync.RWMutex
	methodMap map[string]*StatisticsItem
}

/*
	private methods
*/

func (s *StatisticsMap) get() {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	fmt.Printf("\n %-35s %-10s %-10s %-10s %-10s %-10s\n", "API", "Res/Req", "Avg", "Min", "Max", "Last")
	for k, item := range s.methodMap {
		fmt.Printf(" %-35s %-10v %-10.4f %-10.4f %-10.4f %-10.4f\n", k, fmt.Sprintf("%v/%v", item.ResponseTimes, item.RequestTimes),
			item.TotalDuration.Seconds()/float64(item.ResponseTimes)*1000,
			item.MinDuration.Seconds()*1000,
			item.MaxDuration.Seconds()*1000,
			item.LastDuration.Seconds()*1000,
		)
	}
}

func (s *StatisticsMap) recordReq(method string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.methodMap[method]; ok {
		s.methodMap[method].RequestTimes++
	} else {
		s.methodMap[method] = &StatisticsItem{
			RequestTimes:  1,
			ResponseTimes: 0,
			TotalDuration: 0,
			MaxDuration:   0,
			MinDuration:   time.Second,
			LastDuration:  0,
		}
	}
}

func (s *StatisticsMap) recordResAndStat(method string, duration time.Duration) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// record response
	if m, ok := s.methodMap[method]; ok {
		m.ResponseTimes++
		switch {
		case true:
			m.TotalDuration += duration
			m.LastDuration = duration
			fallthrough
		case m.MaxDuration < duration:
			m.MaxDuration = duration
			fallthrough
		case m.MinDuration > duration:
			m.MinDuration = duration
			// fallthrough
		}
	} else {
		// may not have a statistics item when in realtime send situation
		s.methodMap[method] = &StatisticsItem{
			RequestTimes:  0,
			ResponseTimes: 1,
			TotalDuration: duration,
			MinDuration:   duration,
			MaxDuration:   duration,
			LastDuration:  duration,
		}
	}
}
