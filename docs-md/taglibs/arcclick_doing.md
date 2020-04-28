# 内容列表标签
|功能| 描述|
| :------------- |:-------------|
| 作用      | 获取指定文档列表 |
| 属性      | `typeid`: 要读取的栏目的内容<br/>`offset`: 数据读取的Offset位置<br/>`row`:要读取的条数<br/>`orderby`: 排序字段 <br/>`orderway`: 排序方式<br/> `keywords`: 要筛选的关键字<br/>`modelid`: 模型ID 默认为`1(文章模型)`<br/>`flag`: 要读取的属性, 多个用`,`分割 <br/> `noflag`: 不包含的属性: 多个用`,`分割<br/>`titlelen`: 标题长度<br/>`getall`:是否读取所有下级栏目的内容(非列表、详情页)为`false`, 其他页面为`true` <br/>`subday`: 几天内发布的数据 |  
| 可嵌套 | 可与channelartlist 嵌套  |
| 备注 | 数据表: 各个模型的内容, 不可与`pagelist`达到分页效果   |   

# 标签暴露变量

|变量| 描述|
| :------------- |:-------------|
| autoindex      | 迭代时索引值 |
| field | 一条模型记录 `map[string]string` |

# 模型文档属性

可查看数据库各个文档对应数据表的字段定义. 

# 实例说明
### 获取指定内容 
```html
{{yield artlist(row=10, offset=10, keyword="CMS", orderby="id", orderway="asc", modelid=2, getall=true, subday=3, flag="c") content}}
    - [{{field["catname"]}}]<a href="{{field["arcurl"]}}">{{field["title"]}}</a>
{{end}}
```