import { BaseService, Service } from "/@/core";

@Service("member/group")
class SysMemberGroup extends BaseService {
	select(params: any) {
		return this.request({
			url: "/select",
			method: "GET",
			data: {
				...params
			}
		});
	}
}

export default SysMemberGroup;
