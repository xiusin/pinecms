export const schema = {
  type: 'page',
  body: {
    type: 'lib-crud',
    api: '$preset.apis.list',
    perPageField: 'size',
    pageField: 'page',
    perPageAvailable: [50, 100, 200],
    defaultParams: {
      size: 50,
    },
    headerToolbar: [
      {
        type: 'columns-toggler',
        align: 'left',
      },
      {
        type: 'pagination',
        align: 'left',
      },
      {
        $preset: 'actions.add',
        align: 'right',
      },
    ],
    footerToolbar: ['statistics', 'switch-per-page', 'pagination'],
    columns: [
      {
        name: 'userid',
        label: 'ID',
        type: 'text',
      },
      {
        name: 'username',
        label: '登录账号',
        type: 'text',
      },
      {
        name: 'realname',
        label: '名称',
        type: 'text',
      },
      {
        name: 'avatar',
        label: '头像',
        type: 'tpl',
        tpl: '<img style="width:30px;" src="${avatar}" />',
        popOver: {
          body: '<div class="w-xxl"><img class="w-full" src="${avatar}"/></div>',
        },
      },
      {
        name: 'rolename',
        label: '角色名',
        type: 'tpl',
        tpl: '<%= data.rolename + " (" + data.roleid +")" %>',
      },
      {
        name: 'lastlogintime',
        label: '登录时间',
        type: 'datetime',
      },
      {
        name: 'lastloginip',
        label: '登录IP',
        type: 'text',
      },
      {
        type: 'operation',
        label: '操作',
        limits: ['editItem', 'delItem'],
        limitsLogic: 'or',
        buttons: ['$preset.actions.edit', '$preset.actions.del'],
      },
    ],
  },
  definitions: {
    updateControls: {
      controls: [
        {
          type: 'text',
          name: 'username',
          label: '账号',
          required: true,
        },
        {
          type: 'password',
          name: 'password',
          label: '密码',
          requiredOn: 'typeof data.userid === "undefined"',
        },
        {
          type: 'password',
          name: 'pwdconfirm',
          label: '确认密码',
          validations:"equalsField:password",
          requiredOn: 'typeof data.userid === "undefined"',
        },
        {
          type: 'email',
          name: 'email',
          label: '邮箱',
          required: true,
        },
        {
          type: 'text',
          name: 'realname',
          label: '名称',
        },
        {
          type: 'select',
          name: 'roleid',
          label: '所属角色',
          required: true,
          source: "GET public/select?type=role"
        },
      ],
    },
  },
  preset: {
    actions: {
      add: {
        limits: 'addItem',
        type: 'button',
        align: 'right',
        actionType: 'dialog',
        label: '添加',
        icon: 'fa fa-plus pull-left',
        size: 'sm',
        primary: true,
        dialog: {
          title: '新增',
          body: {
            type: 'form',
            name: 'sample-edit-form',
            api: '$preset.apis.add',
            $ref: 'updateControls',
          },
        },
      },
      edit: {
        limits: 'editItem',
        type: 'button',
        icon: 'fa fa-pencil',
        tooltip: '编辑',
        actionType: 'dialog',
        dialog: {
          title: '编辑',
          body: {
            type: 'form',
            name: 'sample-edit-form',
            api: '$preset.apis.edit',
            $ref: 'updateControls',
          },
        },
      },
      del: {
        limits: 'delItem',
        type: 'button',
        icon: 'fa fa-times text-danger',
        actionType: 'ajax',
        tooltip: '删除',
        confirmText: '您确认要删除?',
        api: '$preset.apis.del',
        messages: {
          success: '删除成功',
          failed: '删除失败',
        },
      },
    },
  },
}
