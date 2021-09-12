import { BaseService, Service } from "/@/core";

@Service("public")
class Common extends BaseService {
	/**
	 * 文件上传模式
	 */
	uploadMode() {
		return this.request({
			url: "/uploadMode"
		});
	}

	/**
	 * 图片验证码 svg
	 *
	 * @param {*} { height, width }
	 * @returns
	 * @memberof CommonService
	 */
	captcha({ height, width }: any) {
		return this.request({
			url: "/captcha",
			params: {
				height,
				width
			}
		});
	}

	/**
	 * 文件上传，如果模式是 cloud，返回对应参数
	 *
	 * @returns
	 * @memberof CommonService
	 */
	upload(params: any) {
		return this.request({
			url: "/upload",
			method: "POST",
			params
		});
	}

	/**
	 * 权限信息
	 *
	 * @returns
	 * @memberof CommonService
	 */
	permMenu() {
		return this.request({
			url: "/menu"
		});
	}
}

export default Common;
