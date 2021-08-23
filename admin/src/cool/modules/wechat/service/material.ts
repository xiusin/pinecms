import { BaseService, Service } from "/@/core";

@Service("wechat/material")
class WechatMaterial extends BaseService {
	total() {
		return this.request({
			url: "/total",
			method: "POST"
		});
	}
	upload(data: any) {
		return this.request({
			url: "/upload",
			method: "POST",
			data,
			headers: { "Content-Type": "multipart/form-data" }
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
