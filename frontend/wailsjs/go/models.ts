export namespace main {
	
	export class DirList {
	    name: string;
	    isDir: boolean;
	    path: string;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new DirList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.isDir = source["isDir"];
	        this.path = source["path"];
	        this.type = source["type"];
	    }
	}

}

