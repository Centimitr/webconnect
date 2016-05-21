package xstatistics

import (
	msg "github.com/Centimitr/xmessage"
)

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
