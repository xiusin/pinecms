# Query标签
|功能| 描述|
| :------------- |:-------------|
| 作用      | 用于自由查询SQL, 最大化扩展灵活性 |
| 属性      | `sql`: 要执行的查询语句, 返回结构为: `[]map[string]string`    |  
| 可嵌套 | 可与任意标签嵌套 |
| 备注 | 只支持查询语句, 数据表前缀用`#@_`代替 |   


# 标签暴露变量

|变量| 描述|
| :------------- |:-------------|
| autoindex      | 迭代时索引值 |
| field | 一条查询数据记录内容: `map[string]string` |

# 实例说明 

### 查询文章总数
```jettemplatelanguage
{{ yield query(sql="select count(1) as total from #@_articles") content}}
    模型下共有内容{{field["total"]}}
{{end}}
```
