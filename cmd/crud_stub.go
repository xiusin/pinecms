package cmd

const (
	controllerDir = "src/application/controllers/backend/"
	modelDir      = "src/application/models/"
	tableDir      = modelDir + "tables/"
	feDir         = "frontend/src/pages/"
	theme         = "vim"
	controllerTpl = `package backend
import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type [ctrl] struct {
	BaseController
}

func (c *[ctrl]) Construct() {
	c.BindType = "form"
	c.SearchFields = map[string]searchFieldDsl{}
	c.Orm = pine.Make(controllers.ServiceXorm).(*xorm.Engine)
	c.Table = &tables.[table]{}
	c.Entries = &[]tables.[table]{}
}`
	modelTpl = `package models

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine/di"
)

type [model] struct {
	orm *xorm.Engine
}

func New[model]() *[model] {
	return &[model]{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}`

	tableTpl = `package tables

[struct]
`

	indexTsTpl = `export const schema = {
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
      [tableDSL]
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
      controls: [formDSL],
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
          [filterDSL]
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
`

	presetTsTpl = `export default {
  limits: {
    $page: {
      label: '查看列表',
    },
    add: {
      label: '添加',
    },
    edit: {
      label: '编辑',
    },
    del: {
      label: '删除',
    },
  },
  apis: {
    list: {
      url: 'GET [table]/list',
      limits: '$page',
      onPreRequest: (source) => {
        const { dateRange } = source.data
        if (dateRange) {
          const arr = dateRange.split('%2C')
          source.data = {
            ...source.data,
            startDate: arr[0],
            endDate: arr[1],
          }
        }
        return source
      },
    },
    add: {
      url: 'POST [table]/add',
      limits: 'add',
    },
    edit: {
      url: 'POST [table]/edit?linkid=$id',
      limits: 'edit',
    },
    del: {
      url: 'POST [table]/delete?id=$id',
      limits: 'del',
    },
  },
}`
)
