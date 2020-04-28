# 文档列表标签

|功能| 描述|
| :------------- |:-------------|
| 作用      | 表示列表模板里的分页内容列表(包括下级所有同模型) |
| 属性      | pagesize: 要读取的记录数<br/>orderby: 要是调用排序字段(默认为: `id`)<br/>orderway: 排序方式(默认为:`desc`) |   
| 可嵌套 | 可与`artlist`, `channel`, `type` 嵌套  |

# 标签暴露变量

|变量| 描述|
| :------------- |:-------------|
| autoindex | 迭代数据索引值 |
| field | 一条分类信息 |

# 实例说明 

### 每页按两条记录分页
```jettemplatelanguage
{{ yield list(pagesize=2, titlelen=25) content}}
   <li>
        <a href="{{field["arcurl"]}}">
            <img src="{{field["thumb"]}}"/> 
            <span>{{field["title"]}}</span> 
        </a>
    </li>
{{end}}
{{yield pagelist()}}
```
