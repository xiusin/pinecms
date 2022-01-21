import { BaseService, Service } from "/@/cool";

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
