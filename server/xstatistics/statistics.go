package xstatistics

import (
	"fmt"
	"sync"
	"time"
)

type StatisticsItem struct {
	// Method        string
	RequestTimes  int
	ResponseTimes int
	AvgDuration   float64
	MinDuration   float64
	MaxDuration   float64
}

type StatisticsMap struct {
	lock      sync.RWMutex
	methodMap map[string]*StatisticsItem
	timeMap   map[string]time.Time
}

/*
	private methods
*/

func (s *StatisticsMap) addRequest(method string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.methodMap[method]; ok {
		s.methodMap[method].RequestTimes++
	} else {
		s.methodMap[method] = &StatisticsItem{
			RequestTimes:  1,
			ResponseTimes: 0,
		}
	}
}

func (s *StatisticsMap) addResponse(method string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.methodMap[method]; ok {
		s.methodMap[method].ResponseTimes++
	} else {
		// not to report error is for realtime usage situation
		s.methodMap[method] = &StatisticsItem{
			RequestTimes:  0,
			ResponseTimes: 1,
		}
	}
}

func (s *StatisticsMap) get() {
	s.lock.RLock()
	defer s.lock.RUnlock()
	fmt.Printf("\n %-35s %-10s %-10s %-10s %-10s %-10s\n", "API", "ReqTimes", "ResTimes", "Avg", "Min", "Max")
	for k, item := range s.methodMap {
		fmt.Printf(" %-35s %-10v %-10v %-10f %-10f %-10f\n", k, item.RequestTimes, item.ResponseTimes, item.AvgDuration*1000, item.MinDuration*1000, item.MaxDuration*1000)
	}
}
