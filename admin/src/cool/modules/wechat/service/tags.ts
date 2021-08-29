import { BaseService, Service } from "/@/core";

@Service("wechat/user/tags")
class WechatTags extends BaseService {
	sync(data: any) {
		return this.request({
			url: "/sync",
			method: "POST",
			data
		});
	}
}

export default WechatTags;
