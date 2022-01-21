import { BaseService, Service } from "/@/cool";

@Service("wechat/user/tags")
class WechatTags extends BaseService {
	sync(data: any) {
		return this.request({
			url: "/sync",
			method: "POST",
			data
		});
	}
	tagging(data: any) {
		return this.request({
			url: "/tagging",
			method: "POST",
			data
		});
	}
}

export default WechatTags;
