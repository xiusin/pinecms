import {BaseService, Service} from "/@/core";

@Service("assets")
class SysAssets extends BaseService {
	themes(params: any) {
		return this.request({
			url: "/themes",
			method: "GET",
			data: {
				...params
			}
		});
	}

	theme(params: any) {
		return this.request({
			url: "/theme",
			method: "POST",
			data: {
				...params
			}
		});
	}

	select(params: any) {
		return this.request({
			url: "/select",
			method: "GET",
			data: {
				...params
			}
		});
	}
}

export default SysAssets;
