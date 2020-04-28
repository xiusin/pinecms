# 导入织梦模板

如果您已经有喜欢的织梦主题, 可以直接通过提供的导入命令完成一键转换模板标签内容

# 命令
```shell
./pinecms import import dedeTpl \
    --dedepath xxx\ 
    --dirname xxx\
    --force xxx
```
- `dedepath`: 要导入的织梦模板地址
- `dirname`: 要生成pinecms的主题名, 不填写默认为`dedepath`的目录名
- `force`: 如果存在同名目录, 是否强制删除

# 案例
```shell
./pinecms import import dedeTpl \
    --dedepath /dede/templates/default\ 
    --dirname dede\
    --force true
```
以上命令表示为:要导入织梦项目下`default`模板并且生成为`pinecms`的主题目录下`dede`目录, 如果存在`dede`目录, 则删除后再导出. 

:::warning 注意
仅支持转换可检测的标签, 大致的字段映射, 一般翻译过来直接渲染会报错, 您可以根据报错修改标签的内容.
这样做的意义是您不用再将模板的`{dede:xxx}{/dede:xxx}` 这样的内容挨个的手动替换, 节省您无聊的工作过程.
不过目前能力有限, 无法达到一键可用~. 最理想的情况下可以达到, 不过您还是不要抱太大希望😢, 准备好手动修改部分错误吧. 在后期的版本中,尽量让体验变得更好.  
:::


一般的执行得到的的结果是这样的: 

![](https://raw.githubusercontent.com/xiusin/assets/master/20200428113856.png)

此处会打印每个标签的输出过程, 您可以核对, 也可以忽略. 