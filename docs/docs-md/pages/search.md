# 搜索页面
搜索页面一般用用来针对用户即时搜索使用的页面. 开发者需要在各个主题下创建模板 `search.jet` 

::: warning 告知
不支持多模型内容同时查询, 请悉知.
 
如需改造搜索内容可以修改`src/controllers/frontend/search.go`控制器.
:::


下面列举可用字段:

|字段| 描述|
| :------------- |:-------------|
| keywords      | 搜索关键字, 用于提取内容, 不得小于两个字节, 多个用空格分离 |
| page      | 页码   |  
| pagesize | 分页记录条数, 默认15  |
|orderby| 排序条件字段: 想要查询的表字段, 排序规则为逆序|
| starttime | 起始时间: <br/>如果为纯数字就视为天数. <br/>正确格式为: 10 或 2020-03-12 或 2020-03-12 11:11:11  |
| typeid | 分类ID (会同时查询该分类下的所有分类的内容) |
| kwtype | 多关键字查询的条件: <br/>   1: 只要满足一个即可<br/>   0: 必须都满足 |
| channeltype | 模型ID 与 typeid相互影响<br/> typeid=0,channeltype!=0时,根据 channeltype查询需要关联的表名<br/>typeid!=0,channeltype=0则根据typeid决定channeltype和表名<br/>两者都为0时就默认为文章模型|

# 控制器导出的模板变量如下:

|变量| 描述|
| :------------- |:-------------|
| list      | 符合条件的记录数组. 格式为: `[]map[string]string` |
| pagelist | 页码组件生成HTML的函数 |

# 使用案例

```jettemplatelanguage

{{range field := list }}
    <a href="{{field["arcurl"]}}" target="_blank">{{field["title"]}}</a><span class="state other"> {{format_time(field["pubtime"], "01月02日")}}</span>
{{end}}

{{pagelist() | unsafe}}
```