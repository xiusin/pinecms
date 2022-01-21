import { BaseService, Service } from "/@/cool";

@Service("department")
class SysDepartment extends BaseService {
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

export default SysDepartment;
