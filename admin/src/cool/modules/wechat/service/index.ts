import Account from "./account";
import User from "./user";
import Msg from "./msg";
import Qrcode from "./qrcode";
import Material from "./material";
import Rule from "./rule";
import Template from "./template";
import Tags from "./tags";
import Menu from "./menu";

export default {
	wechat: {
		account: new Account(),
		user: new User(),
		msg: new Msg(),
		qrcode: new Qrcode(),
		rule: new Rule(),
		material: new Material(),
		template: new Template(),
		tags: new Tags(),
		menu: new Menu(),
	}
};
