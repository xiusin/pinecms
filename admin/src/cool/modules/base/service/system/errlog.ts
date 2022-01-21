import { BaseService, Permission, Service } from "/@/cool";

@Service("errlog")
class SysErrLog extends BaseService {
	@Permission("clear")
	clear() {
		return this.request({
			url: "/clear",
			method: "POST"
		});
	}
}

export default SysErrLog;
