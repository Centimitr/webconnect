package xstatistics

import (
	msg "github.com/Centimitr/xmessage"

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
			RequestTimes:  0,
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
		s.methodMap[method] = &StatisticsItem{
			RequestTimes:  0,
			ResponseTimes: 0,
		}
	}
}

func (s *StatisticsMap) get() {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for k, item := range s.methodMap {
		fmt.Println(k, item.RequestTimes, item.ResponseTimes)
	}
}

/*
	public methods used in processors
*/
func (s *StatisticsMap) Get() {
	s.get()
}

/*
	middleware hooked methods
*/

func (s StatisticsMap) AfterReceive() {
	s.addRequest("req")
}

func (s StatisticsMap) BeforeProcess() {

}

func (s StatisticsMap) AfterProcess() {

}

func (s StatisticsMap) BeforeSend() {

}

func (s StatisticsMap) AfterSend() {
	s.addResponse("res")
}

/*
	init
*/

type Statistcs struct {
	StatisticsMap
}

func init() {
	msg.LoadMiddleware(Statistcs{
		StatisticsMap{
			methodMap: make(map[string]*StatisticsItem),
		},
	})
}
