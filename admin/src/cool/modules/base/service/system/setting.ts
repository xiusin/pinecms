import { BaseService, Service } from "/@/core";

@Service("setting")
class SysSetting extends BaseService {
	groupList() {
		return this.request({
			url: "/groups",
			method: "POST"
		});
	}
}

export default SysSetting;
