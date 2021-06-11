import { BaseService, Service } from "/@/core";

@Service("setting")
class SysSetting extends BaseService {
	groupList(data: any) {
		return this.request({
			url: "/groupList",
			method: "POST",
			data
		});
	}
}

export default SysSetting;
