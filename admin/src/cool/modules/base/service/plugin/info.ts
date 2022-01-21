import { BaseService, Service, Permission } from "/@/cool";

@Service("plugin")
class PluginInfo extends BaseService {
	config(data: any) {
		return this.request({
			url: "/config",
			method: "POST",
			data
		});
	}

	install(data: any) {
		return this.request({
			url: "/install",
			method: "POST",
			data
		});
	}

	getConfig(params: any) {
		return this.request({
			url: "/config",
			params
		});
	}

	enable(data: any) {
		return this.request({
			url: "/enable",
			method: "POST",
			data
		});
	}
}

export default PluginInfo;
