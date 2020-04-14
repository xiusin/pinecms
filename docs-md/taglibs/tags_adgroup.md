---
title: 广告分组标签:adgroup
---

# 广告分组列表标签
|功能| 描述|
| :------------- |:-------------|
| 作用      | 获取一组广告广告位列表循环输出, 返回的为切片格式 |
| 属性      | id: 读取广告的ID参数 <br/> pos: 广告位名称或ID (目前不支持多个调用) <br/> <br/>  **id和pos多个则用,分割**   |  
| 可嵌套 | 可与ad 嵌套, 自动向ad传递`id`/`pos`数据.  |
| 备注 | 数据来源: 广告管理 > 广告位列表<br/> <br/> 数据表: `pinecms_advert_space`    |   


# 标签暴露变量

|变量| 描述|
| :------------- |:-------------|
| autoindex      | 迭代时索引值 |
| field | 一条广告位的数据结构 |

# 广告属性

|属性| 描述|
| :------------- |:-------------|
| Id      | 广告位ID |
| Name      | 广告位名称   |  

# 实例说明 

### 获取所有广告位
```jettemplatelanguage
{{ yield adgroup() content}}
    <li>{{autoindex}}<img src="{{field.Image}}"> 
{{end}}
```

### 数据指定位置的广告位
```jettemplatelanguage
{{ yield adgroup(pos="首页banner,内页banner", id="1,2,3,4") content}}
    <li>{{autoindex}} {{field.Name}}
{{end}}
```

### 与ad标签嵌套获取所有广告
```jettemplatelanguage
{*嵌套调用*}
{{ yield adgroup() content}}
    {{ yield ad() content}}
        <a href="{{field.LinkUrl}}">{{autoindex}}<img src="{{field.Image}}"></a>
    {{end}}
{{end}}
```