import { BaseService, Service } from "/@/cool";

@Service("model")
class SysModel extends BaseService {
	getSQL(params: any) {
		return this.request({
			url: "/sql",
			method: "GET",
			params
		});
	}
	modelTable(params: any) {
		return this.request({
			url: "/table",
			method: "GET",
			params
		});
	}
	select(params: any) {
		return this.request({
			url: "/select",
			method: "GET",
			params
		});
	}
}

export default SysModel;
