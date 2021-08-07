import { BaseService, Service } from "/@/core";

@Service("dict/category")
class SysDictCategory extends BaseService {
	select(params: any) {
		return this.request({
			url: "/select",
			method: "GET",
			params
		});
	}
}

export default SysDictCategory;
