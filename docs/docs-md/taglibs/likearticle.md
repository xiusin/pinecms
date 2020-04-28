# 相似文章标签

|功能| 描述|
| :------------- |:-------------|
| 作用      | 获取一组相似的文档内容 |
| 属性      | row: 读取条数 <br/> kws: 要匹配的关键字(默认为本文档的关键字) <br/> sort: titlelen: 返回标题长度 |  
| 可嵌套 | 不可嵌套 |
| 备注 | 标签属于`文档页面`, 其他页面调用不生效 |   


# 标签暴露变量

|变量| 描述|
| :------------- |:-------------|
| field | 一条文章信息 `map[string]string` |

> 具体参数可以参考各模型数据表 

# 实例说明 

### 查找两条包含`苹果` 或 `香蕉`的数据记录, 标题最大保留25个字符
```jettemplatelanguage
{{ yield likearticle(row=2, kws="苹果,香蕉", titlelen=25) content}}
   <li>
        <a href="{{field.Url}}">
            <img src="{{field.Logo}}"/> 
            <span>{{field.Name}}</span> 
        </a>
    </li>
{{end}}
```
