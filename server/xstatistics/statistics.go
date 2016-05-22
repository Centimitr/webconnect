package xstatistics

import (
	"fmt"
	"sync"
)

type StatisticsItem struct {
	// Method        string
	RequestTimes  int
	ResponseTimes int
}

type StatisticsMap struct {
	lock      sync.RWMutex
	methodMap map[string]*StatisticsItem
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
	for k, item := range s.methodMap {
		fmt.Println(k, "ReqTimes:", item.RequestTimes, "ResTimes:", item.ResponseTimes)
	}
}
