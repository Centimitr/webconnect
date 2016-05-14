class WebMessage {
    constructor(url) {
        this.url = url;
        this.websocket;
        this.queue = [];
        this.queue.queuedNum = function(){
        	let num = 0;
        	for (let item of this){
        		if (item.isQueued()) {
        			num++;
        		}
        	}
        	return num;
        };
        this.queue.sentNum = function(){
        	let num = 0;
        	for (let item of this){
        		if (item.isSent()) {
        			num++;
        		}
        	}
        	return num;
        };
        this.queue.getFirstQueued = function () {
            for (let item of this) {
                if (item.isQueued()) {
                    return item;
                }
            }
        };
        this._requestOrder = 0;
        this._isConnecting = false;
        this.MAX_WAITING_NUM = 64;
        this.CONTINUE_SEND_REST_DELAY = 10;
        this.CONTINUE_SEND_REST_RATIO = 0.5;
        this.REQUEST_TIMEOUT = 250;
        this.PROMISE_TIMEOUT = 6400;
        // this.CONNECTION_TIMEOUT = 4000;
        this._connect();
    };

    _isConnected() {
        return this.websocket && this.websocket.readyState === 1;
    };

    _sendQueueRequests() {
        if (this._isConnected()) {
        	if (this.queue.sentNum() <= this.MAX_WAITING_NUM*this.CONTINUE_SEND_REST_RATIO) {
        		let sendNum = Math.min(this.queue.queuedNum(),this.MAX_WAITING_NUM-this.queue.sentNum());
            	while (sendNum--) {
            	    let w = this.queue.getFirstQueued();
            	    this.websocket.send(JSON.stringify(w.request));
            	    w.startTiming(this.REQUEST_TIMEOUT);
            	    w.setSent();
            	}
        	}
        	if (this.queue.queuedNum()>0) {
        		this._continueSendRest();
        	}
        }else{
        	this._connect();
        }
    }
    _continueSendRest(){
    	//console.log('RETRY');
    	setTimeout(()=>{
    		this._sendQueueRequests();
    	}, this.CONTINUE_SEND_REST_DELAY);
    }

    call(method, params, data) {
        return new Promise((resolve, reject) => {
            // add new task.
            let w = new Task((this._requestOrder++) + '.' + method, method, params, data);
        	w.onreceive = (data) =>{
        		w.setResolve();
            	resolve(data);
            	return;
        	};
        	//console.log('METHOD: Send.');
        	this.queue.push(w);
        	// do task in queue
        	this._sendQueueRequests();
        	// timeout
        	w.ontimeout = ()=>{
        		w.setReject();
        		reject('Request Timeout.');
        	};
        	setTimeout(() => {
        		w.setReject();
        		reject('Promise Timeout.');
        	},this.PROMISE_TIMEOUT);
    	});
    }

    // connection
    _connect() {
    	if (!this.websocket || this.websocket.readyState===3) {
    		this._isConnecting = true;
        	//console.log('SOCKET: Try to connect.');
        	this.websocket = new WebSocket(this.url);
        	this.websocket.onopen = (event) =>{
        	    //console.log('SOCKET: Open.');
        	    // send all requests in queue
        	    //console.log('SOCKET: Start Clearing the Queue.');
        	    this._sendQueueRequests();
        	};
        	this.websocket.onmessage = (event) =>{
        	    let data = JSON.parse(event.data);
        	    //console.log('SOCKET: Message.', data.id);
        	    // match and clear 1 request.
        	    let havntMatch = true;
        	    this.queue.map((w, i) => {
        	        if(havntMatch && w.isSent() &&w.id === data.id){
        	        	//console.log("METHOD: Catch matched response: " + w.id);
        	        	w.onreceive(data);
        	        	havntMatch = false;
        	        }
        	    })
        	};
        	this.websocket.onclose = () =>{};
        	this.websocket.onerror = () =>{};
        	this._isConnecting = false;
    	}    
    };

    disconnect() {
        if (this._isConnected()) {
            this.websocket.close();
            //console.log('SOCKET: Close.');
        }
    };
}