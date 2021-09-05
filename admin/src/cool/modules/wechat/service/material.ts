import { BaseService, Service } from "/@/core";
import { getUrl } from "../../../../core/utils";
import store from "/@/store";

@Service("wechat/material")
class WechatMaterial extends BaseService {
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
	preview(url: string) {
		const token = store.getters.token || "";
		return getUrl(
			"/wechat/material/preview?url=" + encodeURIComponent(url) + "&token=" + token
		);
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
