import { BaseService, Service } from "/@/core";

@Service("wechat/material")
class WechatMaterial extends BaseService {
	total() {
		return this.request({
			url: "/total",
			method: "POST"
		});
	}
	clear() {
		return this.request({
			url: "/clear",
			method: "POST"
		});
	}
}

export default WechatMaterial;
