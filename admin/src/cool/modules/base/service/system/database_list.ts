import {BaseService, Service} from "/@/core";

@Service("database")
class SysDatabaseList extends BaseService {
	backup(params: any) {
		return this.request({
			url: "/backup",
			method: "POST",
			data: {
				...params
			}
		});
	}

	repair(params: any) {
		return this.request({
			url: "/repair",
			method: "POST",
			data: {
				...params
			}
		});
	}

	optimize(params: any) {
		return this.request({
			url: "/optimize",
			method: "POST",
			data: {
				...params
			}
		});
	}
}

export default SysDatabaseList;
