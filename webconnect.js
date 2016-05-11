class Webconnect {
	constructor(url){
		this.url = url;
		this.connect();
	};
	url = '';
	websocket;
	protocol;
	dataQueue = [];
	// websocket original methods
	send(data){
		this.websocket.send(data);
	};
	receive(){
		return this.dataQueue.pop();
	};
	// connection
	connenct(url){
		this.websocket = new WebSocket(url)
		this.websocket.onopen = function(event){
			console.log('Connected.');
		};
		this.websocket.onmessage = function(event){
			console.log(event.data);
			this.dataQueue.push(event.data);
		};
	};
	disconnect(){
		this.websocket.close();
	};
}
