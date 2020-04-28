# 频道标签
|功能| 描述|
| :------------- |:-------------|
| 作用      | 用于获取栏目列表|
| 属性      | typeid: 要查询的分类ID<br/>reid: 上级栏目ID<br/> row: 读取栏目数<br/> type: son表示下级栏目,self表示同级栏目,top顶级栏目 <br/> noself 是否排除当前栏目 (可选值为yes) |  
| 嵌套 | 可与 channelartlist 嵌套 |

# 注意事项
1. 如果reid和typeid均设置有值, 那么reid会被修正为typeid的父ID
2. type=self时会只关注reid参数
3. type=top会直接读取父ID为0的栏目
4. type=son会读取直属子栏目`(如果没有子栏目会返回同级分类)`
5. 在没有指定typeid的情况下，type标记与模板的环境有关，如: 模板生成到栏目一或者栏目一下的文章，那么`type='son'`就表示栏目一的所有子类.
6. typeid不支持传入多个

# 标签暴露变量
|变量| 描述|
| :------------- |:-------------|
| field | 一条分类信息 |

# 实例说明 

### 读取分类3下的所有子分类
```html
{{ yield channel(typeid=3, type="son") content}}
   {{field.Url}} -- {{field.Catname}}
{{end}}
```

### 读取顶级栏目
> 设置为top, typeid的设置将失效. reid会被设置为0
```html
{{ yield channel(type="top") content}}
   {{field.Url}} -- {{field.Catname}}
{{end}}
```

### 读取同级栏目(同父ID)
```html
{{ yield channel(type="self", typeid=3) content}}
   {{field.Url}} -- {{field.Catname}}
{{end}}
```