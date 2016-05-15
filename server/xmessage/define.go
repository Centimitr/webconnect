package xmessage

type Req struct {
	Id     string `json:'id'`
	Method string `json:'method'`
	Params string `json:'params'`
	Data   string `json:'data'`
}

type Res struct {
	Id     string `json:'id'`
	Method string `json:'method'`
	Params string `json:'params'`
	Data   string `json:'data'`
	Error  string `json:'error'`
}

// type Controller func(Req, *Res)

// type Router map[string]Controller
