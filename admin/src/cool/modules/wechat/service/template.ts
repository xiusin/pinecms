import { BaseService, Service } from "/@/core";

@Service("wechat/template")
class WechatTemplate extends BaseService {
	total() {
		return this.request({
			url: "/total",
			method: "POST"
		});
	}
	sync() {
		return this.request({
			url: "/sync",
			method: "POST"
		});
	}
	clear() {
		return this.request({
			url: "/clear",
			method: "POST"
		});
	}
	send(data: any) {
		return this.request({
			url: "/send",
			method: "POST",
			data
		});
	}
}

export default WechatTemplate;
