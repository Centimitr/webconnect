class Task {
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
	onreceive(){}
}