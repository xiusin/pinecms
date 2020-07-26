export const schema = {
  type: 'page',
  body: {
    "type": "form",
    "title": "修改信息",
    "initApi": '$preset.apis.getAdminInfo', // 初始化表单时调用数据
    "api": '$preset.apis.editAdminInfo', // 保存表单数据调用
    "mode": "horizontal",
    "autoFocus": false,
    // "horizontal": {
    //   "leftFixed": "xs"
    // },
    "controls": [
      {
        "type": "text",
        "name": "username",
        "label": "用户名",
        "required": true,
        "size": "md",
        "disabled": true,
      },
      {
        "type": "text",
        "name": "lastloginip",
        "label": "登录IP",
        "required": true,
        "size": "md",
        "disabled": true,
      },
      {
        "type": "email",
        "name": "email",
        "label": "邮箱",
        "required": true,
        "size": "md",
      },
      {
        "type": "text",
        "name": "realname",
        "label": "真实姓名",
        "required": true,
        "size": "md",
      }
    ]
  }
}
