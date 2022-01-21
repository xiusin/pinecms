import { BaseService, Service } from "/@/cool";

@Service("district")
class SysDistrict extends BaseService {
	select(params: any) {
		return this.request({
			url: "/select",
			method: "GET",
			params
		});
	}

	import() {
		return this.request({
			url: "/import",
			method: "POST"
		});
	}
}

export default SysDistrict;
