package statistics

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

func (s *StatisticsMap) AddRequest(method string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.methodMap[method]; ok {
		s.methodMap[method].RequestTimes++
	} else {
		s.methodMap[method] = &StatisticsItem{
			RequestTimes:  0,
			ResponseTimes: 0,
		}
	}
}

func (s *StatisticsMap) AddResponse(method string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.methodMap[method]; ok {
		s.methodMap[method].ResponseTimes++
	} else {
		s.methodMap[method] = &StatisticsItem{
			RequestTimes:  0,
			ResponseTimes: 0,
		}
	}
}

func (s *StatisticsMap) Get() {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for k, item := range s.methodMap {
		fmt.Println(k, item.RequestTimes, item.ResponseTimes)
	}
}

var Stat *StatisticsMap

func init() {
	Stat = &StatisticsMap{
		methodMap: make(map[string]*StatisticsItem),
	}
}
