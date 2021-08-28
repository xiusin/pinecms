import { BaseService, Service } from "/@/core";

@Service("wechat/account")
class WechatAccount extends BaseService {
	clearQuota() {
		return this.request({
			url: "/clear",
			method: "POST"
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
