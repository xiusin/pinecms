# 友情链接标签

|功能| 描述|
| :------------- |:-------------|
| 作用      | 获取一组友情链接 |
| 属性      | `row`: 读取条数 <br/> `ids`: 指定的ID集合 <br/> `sort`: 排序内容 |  
| 可嵌套 | 不可嵌套 |
| 备注 | 数据表: `pinecms_link` |   


# 标签暴露变量

|变量| 描述|
| :------------- |:-------------|
| field | 一条友情链接信息 |

# 广告属性

|属性| 描述|
| :------------- |:-------------|
| Linkid      | 友链ID |
| Name      | 名称   |  
| Url | 地址  |
| Logo | 图标    |  
| Introduce | 描述    |  

# 实例说明 

### 查询指定的友链
```jettemplatelanguage
{{ yield flink(ids="1,2,3,4,5") content}}
   <li>
        <a href="{{field.Url}}">
            <img src="{{field.Logo}}"/> 
            <span>{{field.Name}}</span> 
        </a>
    </li>
{{end}}
```
