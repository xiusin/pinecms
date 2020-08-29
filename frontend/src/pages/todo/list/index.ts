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
    columns:  [
      
	{
		"label": "序列ID",
		"name": "id",
		"type": "text"
	},
	{
		"label": "用户ID",
		"name": "userid",
		"type": "text"
	},
	{
		"label": "信息内容",
		"name": "message",
		"type": "text"
	},
	{
		"label": "状态",
		"name": "status",
		"type": "text"
	},
	{
		"label": "多选",
		"name": "set_status",
		"type": "text"
	},
	{
		"label": "单选",
		"name": "enum_status",
		"type": "text"
	},
	{
		"label": "图片",
		"name": "image",
		"type": "text"
	},
	{
		"label": "图片集合",
		"name": "images",
		"type": "text"
	},
	{
		"label": "文件",
		"name": "file",
		"type": "text"
	},
	{
		"label": "文件集合",
		"name": "files",
		"type": "text"
	},
	{
		"label": "城市ID",
		"name": "city_id",
		"type": "text"
	},
	{
		"label": "文档内容",
		"name": "content",
		"type": "text"
	}
,
      {
        type: 'operation',
        label: '操作',
        width: 60,
        limits: ['edit', 'del'],
        limitsLogic: 'or',
        buttons: ['$preset.actions.edit', '$preset.actions.del'],
      },
    ],
  },
  definitions: {
    updateControls: {
      controls: [
	{
		"label": "序列ID",
		"name": "id",
		"type": "number"
	},
	{
		"label": "用户ID",
		"name": "userid",
		"type": "number"
	},
	{
		"label": "信息内容",
		"name": "message",
		"type": "textarea"
	},
	{
		"label": "状态",
		"name": "status",
		"type": "number"
	},
	{
		"label": "多选",
		"name": "set_status",
		"options": [
			{
				"label": 1,
				"value": 1
			},
			{
				"label": 2,
				"value": 2
			},
			{
				"label": 0,
				"value": 0
			}
		],
		"type": "checkboxes",
		"value": "'1'"
	},
	{
		"label": "单选",
		"name": "enum_status",
		"options": [
			{
				"label": 2,
				"value": 2
			},
			{
				"label": 0,
				"value": 0
			},
			{
				"label": 1,
				"value": 1
			}
		],
		"type": "radios",
		"value": "'0'"
	},
	{
		"label": "图片",
		"name": "image",
		"type": "image"
	},
	{
		"label": "图片集合",
		"multiple": true,
		"name": "images",
		"type": "image"
	},
	{
		"label": "文件",
		"name": "file",
		"type": "file"
	},
	{
		"label": "文件集合",
		"multiple": true,
		"name": "files",
		"type": "file"
	},
	{
		"label": "城市ID",
		"name": "city_id",
		"type": "number"
	},
	{
		"label": "文档内容",
		"name": "content",
		"type": "rich-text"
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
        actionType: 'drawer',
        label: '添加',
        icon: 'fa fa-plus pull-left',
        size: 'sm',
        primary: true,
        drawer: {
          title: '新增',
          size: 'xl',
          body: {
            type: 'form',
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
          title: '编辑',
          size: 'xl',
          body: {
            type: 'form',
            mode: 'normal',
            api: '$preset.apis.edit',
            $ref: 'updateControls',
          },
        },
      },
      del: {
        limits: 'del',
        type: 'action',
        icon: 'fa fa-times text-danger',
        actionType: 'ajax',
        tooltip: '删除',
        confirmText: '您确认要删除?',
        api: {
          $preset: 'apis.del',
        },
        messages: {
          success: '删除成功',
          failed: '删除失败',
        },
      },
    },
    forms: {
      filter: {
        controls: [
          
	{
		"label": "　　序列ID:",
		"name": "id",
		"options": null,
		"type": "number"
	},
	{
		"label": "　　用户ID:",
		"name": "userid",
		"options": null,
		"type": "number"
	},
	{
		"label": "　　状态:",
		"name": "status",
		"options": null,
		"type": "number"
	},
	{
		"label": "　　多选:",
		"name": "set_status",
		"options": [
			{
				"label": 1,
				"value": 1
			},
			{
				"label": 2,
				"value": 2
			},
			{
				"label": 0,
				"value": 0
			}
		],
		"type": "select"
	},
	{
		"label": "　　单选:",
		"name": "enum_status",
		"options": [
			{
				"label": 2,
				"value": 2
			},
			{
				"label": 0,
				"value": 0
			},
			{
				"label": 1,
				"value": 1
			}
		],
		"type": "select"
	},
	{
		"label": "　　城市ID:",
		"name": "city_id",
		"options": null,
		"type": "number"
	}
,
          {
            type: 'submit',
            className: 'm-l',
            label: '搜索',
          },
        ],
      },  // 搜索
    },
  },
}
