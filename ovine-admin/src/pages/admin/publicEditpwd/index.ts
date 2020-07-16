export const schema = {
  type: 'page',
  body: {
    "type": "form",
    "title": "修改密码",
    "initApi": '$preset.apis.getAdminInfo', // 初始化表单时调用数据
    "api": '$preset.apis.editPwd', // 保存表单数据调用
    "mode": "horizontal",
    "autoFocus": false,
    "horizontal": {
      "leftFixed": "xs"
    },
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
        "name": "realname",
        "label": "真实姓名",
        "required": true,
        "size": "md",
        "disabled": true,
      },
      {
        "type": "text",
        "name": "realname",
        "label": "真实姓名",
        "required": true,
        "size": "md",
        "disabled": true,
      },
      {
        "type": "text",
        "name": "email",
        "label": "邮箱",
        "required": true,
        "size": "md",
        "disabled": true,
      },
      {
        "type": "password",
        "name": "old_password",
        "label": "旧密码",
        "required": true,
        "size": "md"
      },
      {
        "type": "password",
        "name": "new_password",
        "label": "新密码",
        "required": true,
        "size": "md"
      },
      {
        "type": "password",
        "name": "new_pwdconfirm",
        "label": "确认新密码",
        "required": true,
        "validations":"equalsField:new_password",
        "size": "md",
      },
    ]
  }
}
