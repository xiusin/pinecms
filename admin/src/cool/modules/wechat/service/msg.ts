import { BaseService, Service } from "/@/core";

@Service("wechat/msg")
class WechatMsg extends BaseService {
	reply(data: any) {
		return this.request({
			url: "/reply",
			method: "POST",
			data
		});
	}
}

export default WechatMsg;
