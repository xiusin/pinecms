import {BaseService, Service} from "/@/core";

@Service("backup")
class SysDatabaseBackupList extends BaseService {
	download(params: any) {
		return this.request({
			url: "/download",
			method: "POST",
			data: {
				...params
			}
		});
	}
}
export default SysDatabaseBackupList;
