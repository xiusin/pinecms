import { BaseService, Service } from "/@/cool";

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
