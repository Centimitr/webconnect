class WebMessage {
	constructor(url){
		// websocket
		this.url = url;
		this.websocket;
		// webconnct
		this.queue = [];
		this.queue.queuedNum = 0;
		this.queue.getFirstQueued = function(){
			for (let item of this){
				if (item.isQueued) {
					return item;
				}
			}
		};
		// this.queue.add = function(item){
		// 	this.push(item);
		// };
		// this.queue.remove = ()=>{};
		this._order = 0;
		this._currentWaitingRequestNum = 0;
		this.MAX_WAITING_REQUEST_NUM = 100;
		this.REQUEST_TIMEOUT = 3000;
		// this.CONNECTION_TIMEOUT = 4000;
		this._connect();
	};
	_isConnected(){
		return this.websocket && this.websocket.readyState === 1;
	};
	_sendQueueRequests(){
		if (this._isConnected()) {
			let sendNum = this.MAX_WAITING_REQUEST_NUM - this._currentWaitingRequestNum;
			while (sendNum--) {
				let w = this.queue.getFirstQueued();
				this.websocket.send(JSON.stringify(w.request));
				// this.queue.queuedNum++;
				this._currentWaitingRequestNum++;
				w.setSent();
			}
		}
	}
	call(method, params, data){
		return new Promise((resolve,reject)=>{
			// add new task.
			let w = new Task((this._order++)+'.'+method,method,params,data);
			w.onreceive = (data)=>{
				this._currentWaitingRequestNum--;
				resolve(data);
				return;
			};
			console.log('METHOD: Send.');
			this.queue.push(w);
			// do task in queue
			this._sendQueueRequests();
			// timeout
			setTimeout(()=> {
				this._currentWaitingRequestNum--;
				reject('Timeout.');
			}, this.REQUEST_TIMEOUT);
		})
	}
	// connection
	_connect(){
		console.log('SOCKET: New.');
		this.websocket = new WebSocket(this.url)
		this.websocket.onopen = (event)=>{
			console.log('SOCKET: Open.');
			// send all requests in queue
			console.log('SOCKET: Start Clearing the Queue.');
			this._sendQueueRequests();
		};
		this.websocket.onmessage = (event)=>{
			let data = JSON.parse(event.data);
			console.log('SOCKET: Message.',data.id);
			// match and clear 1 request
			.
			let havntMatch = true;
			this.queue.map((w,i)=>{	
				if(havntMatch && w.id===data.id){
					console.log("METHOD: Catch matched response: "+w.id);
					w.onreceive(data);
					this.queue.splice(i,1);
					havntMatch = false;
				}
			})
		};
		this.websocket.onclose = ()=>{};
	};
	disconnect(){
		if (this._isConnected()) {
			this.websocket.close();
			console.log('SOCKET: Close.');
		}
	};
}