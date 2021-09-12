import { BaseService, Service } from "/@/core";

@Service("setting")
class SysSetting extends BaseService {
	groupList() {
		return this.request({
			url: "/groups",
			method: "POST"
		});
	}
	sendTestEmail(data: any) {
		return this.request({
			url: "/test",
			method: "POST",
			data
		});
	}
}

export default SysSetting;
