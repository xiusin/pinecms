import { BaseService, Service } from "/@/core";

@Service("content")
class SysContent extends BaseService {
	select(params: any) {
		return this.request({
			url: "/select",
			method: "GET",
			params
		});
	}

	getPageInfo(params: any) {
		return this.request({
			url: "/page",
			method: "GET",
			params,
		});
	}

	savePageInfo(data: any) {
		return this.request({
			url: "/page",
			method: "POST",
			data,
		});
	}
}

export default SysContent;
