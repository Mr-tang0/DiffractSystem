export namespace components {
	
	export class HVPSSetpointInfo {
	    HV: number;
	    HI: number;
	    FV: number;
	    FI: number;
	
	    static createFrom(source: any = {}) {
	        return new HVPSSetpointInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.HV = source["HV"];
	        this.HI = source["HI"];
	        this.FV = source["FV"];
	        this.FI = source["FI"];
	    }
	}

}

export namespace services {
	
	export class GitHubRelease {
	    tag_name: string;
	    html_url: string;
	    browser_download_url: string;
	    body: string;
	    current_version: string;
	
	    static createFrom(source: any = {}) {
	        return new GitHubRelease(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tag_name = source["tag_name"];
	        this.html_url = source["html_url"];
	        this.browser_download_url = source["browser_download_url"];
	        this.body = source["body"];
	        this.current_version = source["current_version"];
	    }
	}

}

