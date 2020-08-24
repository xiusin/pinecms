export const schema = {
  type: 'page',
  body: {
    type: 'lib-crud',
    api: '$preset.apis.list',
    filter: '$preset.forms.filter',
    filterTogglable: true,
    perPageAvailable: [50, 100, 200],
    defaultParams: {
      size: 50,
    },
    perPageField: 'size',
    pageField: 'page',
    headerToolbar: [
      'filter-toggler',
      {
        type: 'columns-toggler',
        align: 'left',
      },
      {
        type: 'pagination',
        align: 'left',
      },
      '$preset.actions.add',
    ],
    footerToolbar: ['statistics', 'switch-per-page', 'pagination'],
    columns: [
      {
        name: 'name',
        label: '模板名称',
        type: 'text',
      },
      {
        name: 'size',
        label: '文件大小',
        type: 'text',
      },
      {
        name: 'updated',
        label: '修改时间',
        type: 'text',
      },
      {
        type: 'operation',
        label: '操作',
        width: 60,
        limits: ['edit'],
        limitsLogic: 'or',
        buttons: ['$preset.actions.edit'],
      },
    ],
  },
  definitions: {
    updateControls: {
      controls: [
        {
          name: 'name',
          label: '名称',
          type: 'text',
          required: true,
          disabledOn: "data.name"
        },
        {
          name: 'content',
          label: '内容',
          type: 'html-editor',
        }
      ],
    },
  },
  preset: {
    actions: {
      add: {
        limits: 'add',
        type: 'button',
        align: 'right',
        actionType: 'dialog',
        label: '添加',
        icon: 'fa fa-plus pull-left',
        size: 'sm',
        primary: true,
        dialog: {
          title: '新增文档',
          size: 'lg',
          body: {
            type: 'form',
            // debug: true,  // 调试期可以用来打印表单变量的值
            api: '$preset.apis.add',
            mode: 'normal',
            $ref: 'updateControls',
          },
        },
      },
      edit: {
        limits: 'edit',
        type: 'button',
        icon: 'fa fa-pencil',
        tooltip: '编辑',
        actionType: 'dialog',
        dialog: {
          title: '编辑文档',
          size: 'lg',
          body: {
            type: 'form',
            mode: 'normal',
            api: '$preset.apis.edit',
            $ref: 'updateControls',
          },
        },
      },
    }
  },
}
