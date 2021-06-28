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

	exec(params: any) {
		return this.request({
			url: "/exec",
			method: "POST",
			data: {
				...params
			}
		});
	}
}

export default SysDatabaseList;
