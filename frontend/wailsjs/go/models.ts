export namespace main {
	
	export class Container {
	    ID: string;
	    Names: string[];
	    State: string;
	    Image: string;
	    Command: string;
	    Created: string;
	    Status: string;
	    Ports: types.Port[];
	
	    static createFrom(source: any = {}) {
	        return new Container(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Names = source["Names"];
	        this.State = source["State"];
	        this.Image = source["Image"];
	        this.Command = source["Command"];
	        this.Created = source["Created"];
	        this.Status = source["Status"];
	        this.Ports = this.convertValues(source["Ports"], types.Port);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Image {
	    ID: string;
	    Repository: string;
	    Tag: string;
	    Created: string;
	    Size: string;
	
	    static createFrom(source: any = {}) {
	        return new Image(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Repository = source["Repository"];
	        this.Tag = source["Tag"];
	        this.Created = source["Created"];
	        this.Size = source["Size"];
	    }
	}

}

export namespace registry {
	
	export class ServiceConfig {
	    InsecureRegistryCIDRs: NetIPNet[];
	    IndexConfigs: {[key: string]: IndexInfo};
	
	    static createFrom(source: any = {}) {
	        return new ServiceConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.InsecureRegistryCIDRs = this.convertValues(source["InsecureRegistryCIDRs"], NetIPNet);
	        this.IndexConfigs = source["IndexConfigs"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace types {
	
	export class Info {
	    CpuCfsPeriod: boolean;
	    CpuCfsQuota: boolean;
	    BridgeNfIp6tables: boolean;
	    HttpProxy: string;
	    HttpsProxy: string;
	
	    static createFrom(source: any = {}) {
	        return new Info(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.CpuCfsPeriod = source["CpuCfsPeriod"];
	        this.CpuCfsQuota = source["CpuCfsQuota"];
	        this.BridgeNfIp6tables = source["BridgeNfIp6tables"];
	        this.HttpProxy = source["HttpProxy"];
	        this.HttpsProxy = source["HttpsProxy"];
	    }
	}
	export class Port {
	    IP?: string;
	    PrivatePort: number;
	    PublicPort?: number;
	    Type: string;
	
	    static createFrom(source: any = {}) {
	        return new Port(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.IP = source["IP"];
	        this.PrivatePort = source["PrivatePort"];
	        this.PublicPort = source["PublicPort"];
	        this.Type = source["Type"];
	    }
	}
	export class Runtime {
	    path: string;
	    runtimeArgs?: string[];
	
	    static createFrom(source: any = {}) {
	        return new Runtime(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.runtimeArgs = source["runtimeArgs"];
	    }
	}

}

