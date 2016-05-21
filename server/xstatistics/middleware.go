package xstatistics

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
