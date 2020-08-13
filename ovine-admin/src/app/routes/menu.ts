/**
 * 项目内路由，每次新增或者修改页面需要在这里添加路由信息
 * 详细文档：https://ovine.igroupes.com/org/docs/advance/configurations#%E8%B7%AF%E7%94%B1%E9%85%8D%E7%BD%AE
 *
 * 会根据 nodePath 字段，自动匹配 pages 文件夹下的 xxx/index  文件。
 * 由于使用了 文件动态引入，每次新增加页面时，dev server 可能会报 找不到文件错误，重启下 dev server 就好。
 */

import {LimitMenuItem} from '@core/routes/types'
import {getStore} from "@core/utils/store";
import {storeKeys} from "~/app/constants";
import {env} from "~/app/env";


let menus: LimitMenuItem = {
  nodePath: '/',
  limitLabel: '菜单目录',
  label: '',
  children: [
    {
      path: '/',
      label: '后台首页',
      nodePath: 'dashboard',
      exact: true,
      icon: 'fa fa-home',
      pathToComponent: 'dashboard',
      sideVisible: true, // 不会显示在侧边栏
    },
    // {
    //   label: '系统管理',
    //   icon: 'fa fa-wrench',
    //   nodePath: 'system',
    //   children: [
    //     {
    //       label: '管理员用户',
    //       nodePath: 'user_list', // 对应 src/pages/system/user_list
    //     },
    //     {
    //       label: '管理员角色',
    //       nodePath: 'user_role', // 对应 src/pages/system/user_role
    //     },
    //     {
    //       label: '系统操作日志',
    //       nodePath: 'user_log', // 对应 src/pages/system/user_log
    //     },
    //   ],
    // },
  ],
};

var xhr = new XMLHttpRequest();
xhr.open("GET", env.localhost.domains.api + "/index/menu", false);
const { key, token } = getStore(storeKeys.auth) || {}
xhr.setRequestHeader(key, token)
xhr.onload = function (e) {
  if (xhr.readyState === 4) {
    if (xhr.status === 200) {
      const data=JSON.parse(xhr.responseText);
      // 追加自己的菜单
      menus.children = menus.children.concat(data.data);
    } else {
      console.error(xhr.statusText);
    }
  }
};
xhr.onerror = function (e) {
  console.error(xhr.statusText);
};
xhr.send(null);

// todo 动态合并主菜单
export const menuRoutes = menus;
