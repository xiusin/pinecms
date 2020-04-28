# 模板
模板分为后端模板和前端模板:

# 后端模板
后端模板位于目录`resources/views/`下, 这里使用的模板渲染引擎是`html/template`, 所以您直接在这里使用原生的Go模板方法即可.

:::warning 提醒
模板后缀为:`.html`, 因为`Pine`可以同时注册多个模板引擎, 以后缀查找模板引擎. 
::: 

## 模板创建
后端模板是基于`easyui`作为主要页面渲染, 一些`Datagrid` 和 `Treegrid` 已经被封装为go函数调用:
```go
import "github.com/xiusin/pinecms/src/common/helper"

helper.Treegrid("组件名称", "组件需要异步请求的地址", helper.EasyuiOptions{    // easyui的属性配置项
		"title":     models.NewMenuModel().CurrentPos(menuid),
		"toolbar":   "category_categorylist_treegrid_toolbar",// 需要关联的tool组件ID
		"idField":   "catid",   // id字段
		"treeField": "catname", // 名称字段
	}, helper.EasyuiGridfields{
		"字段name":   {"field": "数据字段", "width": "表格宽度", "align": "对齐", "formatter": "字段格式化函数", "index": "索引"},
	})
```
> 需要注意的是, `index`必须按照索引顺序, 否则会渲染失败. 虽然您可以分别写到不同页面, 但是您的组件名称和formatter必须是全局唯一的, 否则可能会引发一些不必要的问题


# 前端页面
由于前端支持多主题, 所以以主题目录(`resources/themes`)为根目录, 可以定义多个主题, 每个主题在一个单独的目录下, 如:
```
.
├── default
├── cc
├── my
└── blog
```
每个目录就代表一个主题, 您可以在后端`资源管理 > 主题列表`选择:
![](https://raw.githubusercontent.com/xiusin/assets/master/20200427161345.png)

:::tip
前端模板引擎使用的是`jet`, 您可以先熟悉一下模板引擎的使用方式, `pinecms`内置了一些常用的标签, 您可以在`taglibs/`中查看代码: 
:::

内置标签的具体使用您可以查看: [标签系统](/taglibs/) 


或者导入织梦模板: [导入模板](/models/import_dede_tpl) 

