import { BaseService, Service } from "/@/cool";

@Service("table")
class SysTable extends BaseService {
	fields() {
		return this.request({
			url: "/fields",
			method: "GET"
		});
	}
}

export default SysTable;
