import { BaseService, Service } from "/@/cool";

@Service("stat")
class SysStat extends BaseService {
	data() {
		return this.request({
			url: "/data",
			method: "GET"
		});
	}
}

export default SysStat;
