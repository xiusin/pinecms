# 上下篇标签

|功能| 描述|
| :------------- |:-------------|
| 作用      | 表示获取文档“上一篇／下一篇”的链接列表 |
| 属性      | `get`: 要读取的连接, get="pre"获取上一篇, get="pre,next": 获取上下篇(默认) <br/>tpl: 生成模板,支持`~arturl~` 和 `~title~` |  
| 可嵌套 | 不可嵌套 |
| 备注 | 标签属于`文档页面`, 其他页面调用不生效 |   

### 实例说明 

### 获取上下篇链接
```jettemplatelanguage
{{ yield prenext() }}
```



### 获取上下篇链接
```jettemplatelanguage
{{ yield prenext(tpl="<a href="~arturl~">~title~</a>") }}
```