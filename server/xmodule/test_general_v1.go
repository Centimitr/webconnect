package xmodule

import (
	msg "github.com/Centimitr/xmessage"
	"github.com/Centimitr/xstatistics"
)

type Test struct {
}

func (t Test) PrintStat(c *msg.Ctx) {
	stat := c.Middleware["Statistics"].(xstatistics.Statistics)
	if c.Error.Fatal == nil {
		stat.Get()
	}
}

func init() {
	var m Test
	msg.LoadModule(m)
}
