import { BaseService, Service, Permission } from "/@/cool";

@Service("task")
class SysTask extends BaseService {
	@Permission("stop")
	stop(data: any) {
		return this.request({
			url: "/stop",
			method: "POST",
			data
		});
	}

	@Permission("start")
	start(data: any) {
		return this.request({
			url: "/start",
			method: "POST",
			data
		});
	}

	@Permission("once")
	once(data: any) {
		return this.request({
			url: "/once",
			method: "POST",
			data
		});
	}

	@Permission("log")
	log(params: any) {
		return this.request({
			url: "/log",
			params
		});
	}

	scriptList() {
		return this.request({
			url: "/script_list",
			method: "POST"
		});
	}

	scriptInfo(data: any) {
		return this.request({
			url: "/script_info",
			method: "POST",
			data
		});
	}
	// 新增或保存内容
	saveInfo(data: any) {
		return this.request({
			url: "/script_save",
			method: "POST",
			data
		});
	}
}

export default SysTask;
