class WebMessageTask {
	constructor(id,method, params, data){
		this.id = id|| Date.now();
		this.status = 'queued';
		this.request = {
			id: this.id,
			method: method || "",
			params: params || {},
			data: data || {}
		}
	}
	isQueued(){
		return this.status === 'queued';
	}
	isSent(){
		return this.status === 'sent';
	}
	setSent(){
		this.status='sent';
	}
	setReject(){
		this.status='reject';
	}
	setResolve(){
		this.status='resolve';
	}
	startTiming(delay){
		setTimeout(()=>{
			if (this.isSent()) {
				this.ontimeout()
			}
		}, delay)
	}
	onreceive(){}
	ontimeout(){}
}