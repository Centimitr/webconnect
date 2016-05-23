package xjsonbase

import (
	"fmt"
	msg "github.com/Centimitr/xmessage"
	"io/ioutil"
)

func (j *JSONBase) Load(c *msg.Ctx) {
	filename := c.Method
	data, err := ioutil.ReadFile(filename + ".json")
	if err != nil {
		fmt.Println(err)
	}
	c.Data = string(data)
}

func (j *JSONBase) Save(c *msg.Ctx) {
	filename := c.Method
	err := ioutil.WriteFile(filename+".json", []byte(c.Data), 0777)
	if err != nil {
		fmt.Println(err)
	}
}
