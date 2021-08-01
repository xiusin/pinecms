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
}

export default SysContent;
