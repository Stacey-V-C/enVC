export namespace types {
	
	export class NVC_Event {
	    action: string;
	    id: string;
	    payload: any;
	
	    static createFrom(source: any = {}) {
	        return new NVC_Event(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.action = source["action"];
	        this.id = source["id"];
	        this.payload = source["payload"];
	    }
	}

}

