# 概述
模板标签主要作用是调取网站系统的数据，其语法格式和HTML标记一样，具有非常简单、易懂，让不懂程序设计的设计师、网页美工等人员也可制作动态网页。
PineCMS模板标签分为以下几大类：

- 全局标签/变量
- 栏目页标签
- 详情页标签

PineCMS 选用 [CloudyKit/jet](https://github.com/CloudyKit/jet) 作为模板引擎. 也是从`pine`继承过来. 

::: tip 为什么选择jet呢?
中间尝试过很多引擎, 发现只有这标签可以做到模拟标签调用, 并且可以设置伪具名参数, 如 : `{{yield ad(pos="参数内容") content} }` 这里`pos` 就可以直接作为参数逐层传递下去. 
:::

> 后期会定制化PineCMS专属解析标签 (Fork jet)
 
# 标签格式
```jettemplatelanguage
{{ yield tag(params1="", params2=1, params3=true) content }}
标签内部循环内容
内部使用可用变量: field 和 autoindex
如:
<img src="{{field.Thumb}}" alt="img{{autoindex}}"/>
{{end}}

或者

{{yield tag(param=1)}}
```

# 全局标签
全局标签: 顾名思义, 可以全局使用的标签或变量. 如: 系统配置, 广告列表, 友情链接系统. 在任何页面均可调用.
 
- adlist | 广告列表标签 
- myad | 广告标签
- artlist | 文档列表标签
- flink | 友链标签
- channel | 分类标签
- channelartlist | 分类文档标签
- type | 指定栏目标签
- query | SQL查询标签
- sonchannel | 子栏目标签

## 栏目页标签
栏目页特有的标签. 

- list | 列表标签
- pagelist | 列表分页

## 详情页标签
详情页特有的标签. 

- prenext | 上下篇标签
- likearticle | 相似文章标签

## 全局变量
系统对模板暴露了系统配置信息表`pinecms_setting`的所有属性, 对模板页为小写字母属性. 
调用方式为:
```
{{global["配置项"]}}
```

- 调用网站信息
```
{{global["site_url"]}}  // 获取网站网址(当前域名,不包括协议)
{{global["site_title"]}}  // 获取网站名称
{{global["site_keywords"]}}  // 获取网站关键字
{{global["site_description"]}}  // 获取网站描述
{{global["site_copyright"]}} // 获取网站copyright
{{global["site_icp"]}}  // 获取网站icp备案
```

> 开发期间可以在后台`系统配置`里开启动态渲染选项. 


