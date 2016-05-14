class WebMessageQueue {
	constructor() {

	}
    queuedNum = function(){
    	let num = 0;
    	for (let item of this){
    		if (item.isQueued()) {
    			num++;
    		}
    	}
    	return num;
    };
    sentNum = function(){
    	let num = 0;
    	for (let item of this){
    		if (item.isSent()) {
    			num++;
    		}
    	}
    	return num;
    };
    getFirstQueued = function () {
        for (let item of this) {
            if (item.isQueued()) {
                return item;
            }
        }
    };
}