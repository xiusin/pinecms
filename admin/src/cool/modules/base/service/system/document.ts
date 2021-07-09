import { BaseService, Service } from "/@/core";

@Service("document")
class SysDocument extends BaseService {
	select(params: any) {
		return this.request({
			url: "/select",
			method: "POST",
			data: {
				...params
			}
		});
	}
}

export default SysDocument;
