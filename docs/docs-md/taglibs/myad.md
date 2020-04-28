# 广告标签
|功能| 描述|
| :------------- |:-------------|
| 作用      | 获取指定的一条广告,返回html内容 |
| 属性      | `id`: 读取广告的ID参数 <br/> `name`: 广告名称   |  
| 可嵌套 | 单独使用即可  |
| 备注 | 数据来源: `广告管理 > 广告列表`<br/> <br/> 数据表: `pinecms_advert`   |   

# 标签暴露变量
标签不对外部暴露任何变量

# 实例说明 

### 获取指定的广告
```jettemplatelanguage
{{ yield myad(id=1) }}
```

### 数据指定名称的广告
```jettemplatelanguage
{{ yield ad(name="首页banner") }}
```
