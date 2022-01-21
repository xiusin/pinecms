import { BaseService, Service } from "/@/cool";

@Service("wechat/user")
class WechatUser extends BaseService {
	sync(data: any) {
		return this.request({
			url: "/sync",
			method: "POST",
			data
		});
	}
}

export default WechatUser;
