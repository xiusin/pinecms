# 广告列表标签
|功能| 描述|
| :------------- |:-------------|
| 作用      | 获取一组广告列表循环输出, 返回的为切片格式 |
| 属性      | `id`: 读取广告的ID参数 <br/> `pos`: 广告位名称或ID (目前不支持多个调用) <br/>  `orderby`: 排序字段 <br/><br/>  **id和pos最少有一个有值, 如果两个都传递了,如果ID不pos内, 则返回空**   |  
| 备注 | 数据来源: 广告管理 > 广告列表<br/> <br/> 数据表: `pinecms_advert`   |   


# 标签暴露变量

|变量| 描述|
| :------------- |:-------------|
| autoindex      | 迭代时索引值 |
| field | 一条广告的数据结构 |

# 广告属性

|属性| 描述|
| :------------- |:-------------|
| Id      | 广告ID |
| Name      | 广告名称   |  
| Image | 图片地址  |
| LinkUrl | 点击广告需要跳转的地址    |  

# 实例说明 

### 获取指定的广告
```jettemplatelanguage
{{ yield adlist(id="1,2,3") content}}
    <li>{{autoindex}}<img src="{{field.Image}}"> 
{{end}}
```

### 数据指定位置的广告位
```jettemplatelanguage
{{ yield adlist(pos="首页banner") content}}
    <li>{{autoindex}}<img src="{{field.Image}}"> 
{{end}}
```

### 获取所有广告位下的所有广告
```jettemplatelanguage
{*嵌套调用*}
{{ yield adgroup() content}}
    {{ yield adlist() content}}
        <a href="{{field.LinkUrl}}">{{autoindex}}<img src="{{field.Image}}"></a>
    {{end}}
{{end}}
```