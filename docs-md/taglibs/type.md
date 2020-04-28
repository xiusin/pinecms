# 单分类标签
|功能| 描述|
| :------------- |:-------------|
| 作用      | 获取指定分类的信息 |
| 属性      | `typeid`: 要查询的分类ID   |  
| 可嵌套 | 不可嵌套 |
| 备注 | 数据表: `pinecms_category` |   


# 标签暴露变量

|变量| 描述|
| :------------- |:-------------|
| field.Catname | 分类名称 |
| field.Url | 分类连接 |

# 实例说明 

### 查询指定分类
```jettemplatelanguage
{{ yield type(typeid=3) content}}
   {{field.Catname}} -- {{field.Url}}
{{end}}
```
