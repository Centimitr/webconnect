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
	TotalDuration time.Duration
	MinDuration   time.Duration
	MaxDuration   time.Duration
	LastDuration  time.Duration
}

type StatisticsMap struct {
	// methodMapLock sync.RWMutex
	// timeMapLock   sync.RWMutex
	mutex     sync.RWMutex
	methodMap map[string]*StatisticsItem
	// timeMap       map[string]time.Time
}

/*
	private methods
*/

// func (s *StatisticsMap) timesStatAR(method string) {
// 	s.methodMapLock.Lock()
// 	defer s.methodMapLock.Unlock()
// 	if _, ok := s.methodMap[method]; ok {
// 		s.methodMap[method].RequestTimes++
// 	} else {
// 		s.methodMap[method] = &StatisticsItem{
// 			RequestTimes:  1,
// 			ResponseTimes: 0,
// 		}
// 	}
// }

// func (s *StatisticsMap) timesStatAS(method string) {
// 	s.methodMapLock.Lock()
// 	defer s.methodMapLock.Unlock()
// 	if _, ok := s.methodMap[method]; ok {
// 		s.methodMap[method].ResponseTimes++
// 	} else {
// 		// not to report error is for realtime usage situation
// 		s.methodMap[method] = &StatisticsItem{
// 			RequestTimes:  0,
// 			ResponseTimes: 1,
// 		}
// 	}
// }

// func (s *StatisticsMap) durationStatAR(id string) {
// 	s.timeMapLock.Lock()
// 	defer s.timeMapLock.Unlock()
// 	s.timeMap[id] = time.Now()
// }

// func (s *StatisticsMap) durationStatAS(method string, id string) {
// 	s.methodMapLock.Lock()
// 	defer s.methodMapLock.Unlock()
// 	s.timeMapLock.Lock()
// 	defer s.timeMapLock.Unlock()

// 	duration := time.Now().Sub(s.timeMap[id]).Seconds()
// 	s.methodMap[method].LastDuration = duration
// 	if s.methodMap[method].MinDuration < 1e-4 {
// 		s.methodMap[method].MinDuration = duration
// 	}
// 	switch {
// 	case s.methodMap[method].MinDuration-duration > 1e-4:
// 		s.methodMap[method].MinDuration = duration
// 	case duration-s.methodMap[method].MaxDuration > 1e-4:
// 		s.methodMap[method].MaxDuration = duration
// 	case true:
// 		times := s.methodMap[method].ResponseTimes
// 		s.methodMap[method].AvgDuration = (float64(times-1)*s.methodMap[method].AvgDuration + duration) / float64(times)
// 	}
// 	// delete(s.timeMap, id)
// 	// if len(s.timeMap) > 64000 {
// 	// 	for id, t := range s.timeMap {
// 	// 		d := time.Now().Sub(t)
// 	// 		if d > time.Second {
// 	// 			delete(s.timeMap, id)
// 	// 		}
// 	// 	}
// 	// }
// }

func (s *StatisticsMap) get() {
	// s.methodMapLock.RLock()
	// defer s.methodMapLock.RUnlock()
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	fmt.Printf("\n %-35s %-10s %-10s %-10s %-10s %-10s\n", "API", "Res/Req", "Avg", "Min", "Max", "Last")
	for k, item := range s.methodMap {
		// fmt.Printf(" %-35s %-10v %-10.6f %-10.6f %-10.6f %-10.6f\n", k, fmt.Sprintf("%v/%v", item.ResponseTimes, item.RequestTimes),
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
	// add duration stat

}
