import { BaseService, Service } from "/@/core";

@Service("stat")
class SysStat extends BaseService{
	data() {
		return this.request({
			url: "/data",
			method: "GET"
		});
	}
}

export default SysStat;
