import { BaseService, Service, Permission } from "/@/core";

@Service("user")
class SysUser extends BaseService {
	/**
	 * 用户退出
	 */
	userLogout() {
		return this.request({
			url: "/logout",
			method: "POST"
		});
	}

	/**
	 * 用户信息
	 *
	 * @returns
	 * @memberof CommonService
	 */
	userInfo() {
		return this.request({
			url: "/admin_info"
		});
	}

	/**
	 * 用户信息
	 *
	 * @returns
	 * @memberof CommonService
	 */
	saveInfo(data: any) {
		return this.request({
			url: "/person_update",
			method: "POST",
			data
		});
	}
	//
	// /**
	//  * 用户信息修改
	//  *
	//  * @param {*} params
	//  * @returns
	//  * @memberof CommonService
	//  */
	// userUpdate(params: any) {
	// 	return this.request({
	// 		url: "/personUpdate",
	// 		method: "POST",
	// 		data: {
	// 			...params
	// 		}
	// 	});
	// }
}

export default SysUser;
