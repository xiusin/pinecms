import Common from "./common";
import Open from "./open";
import SysUser from "./system/user";
import SysMenu from "./system/menu";
import SysRole from "./system/role";
import SysDept from "./system/dept";
import SysTask from "./system/task";
import SysSetting from "./system/setting";
import SysLog from "./system/log";
import PluginInfo from "./plugin/info";
import SysAd from "./system/ad";
import SysLink from "./system/link";
import SysAssets from "./system/assets";
import SysAttachment from "./system/attachment";
import SysDatabaseList from "./system/database_list";
import SysDict from "./system/dict";
import SysDictCategory from "./system/dict_category";
import SysDatabaseBackupList from "./system/database_backup_list";
import SysModel from "./system/dict";
import SysCategory from "./system/category";
import SysDocument from "./system/document";

export default {
	common: new Common(),
	open: new Open(),
	system: {
		user: new SysUser(),
		menu: new SysMenu(),
		role: new SysRole(),
		dept: new SysDept(),
		task: new SysTask(),
		setting: new SysSetting(),
		log: new SysLog(),
		ad: new SysAd(),
		link: new SysLink(),
		assets: new SysAssets(),
		attachment: new SysAttachment(),
		databaseList: new SysDatabaseList(),
		databaseBackupList: new SysDatabaseBackupList(),
		dict: new SysDict(),
		model: new SysModel(),
		dictCategory: new SysDictCategory(),
		category: new SysCategory(),
		document: new SysDocument(),
	},
	plugin: {
		info: new PluginInfo()
	}
};
