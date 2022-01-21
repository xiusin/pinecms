import { BaseService, Service } from "/@/cool";

@Service("wechat/account")
class WechatAccount extends BaseService {
	clearQuota(data: any) {
		return this.request({
			url: "/clear",
			method: "POST",
			data
		});
	}
	distribution(data: any) {
		return this.request({
			url: "/distribution",
			method: "POST",
			data
		});
	}
	select() {
		return this.request({
			url: "/select",
			method: "POST"
		});
	}
}

export default WechatAccount;
