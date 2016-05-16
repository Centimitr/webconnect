package xmessage

import (
	"errors"
	"reflect"
	"strings"
)

type Processor struct {
	Module  string
	PkgPath string
	Name    string
	Func    func(*Ctx) []reflect.Value
}

func registerProcessor(p *Processor) {
	Table[p.Module+"."+p.Name] = p
}

/*
	Client use a part of the table's key (usually is ProcessorName) to match a processor.
	Currently, the key is "PkgName.ProcessorName" so team members should use different package names to do a replacement though their package path is not the same.
*/
func matchProcessor(suffix string) (func(*Ctx) []reflect.Value, error) {
	matchedList := []string{}
	// key is "PkgName.ProcessorName"
	for key, _ := range Table {
		if strings.HasSuffix(key, suffix) {
			matchedList = append(matchedList, key)
		}
	}
	if len(matchedList) == 1 {
		return Table[matchedList[0]].Func, nil
	}
	return nil, errors.New("Cannot select 1 processor.")
}
