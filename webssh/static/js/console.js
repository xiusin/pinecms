settmpdata = function (res) {
    var tmp = JSON.stringify(res);
    window.localStorage.setItem("server", tmp);
}

setnickname = function (data) {
    http_send("/nickname", {
        id: data.ID,
        nickname: data.Nickname,
        ip: data.Ip,
        port: data.Port,
        username: data.Username
    })
}
var open_repass_index = 0;
openpass = function (ID) {
    //console.log(ID)
    if(open_repass_index > 0){
        return false
    }
    open_repass_index = layer.open({
        type: 2,
        maxmin : true,
        title: '重置加密密码',
        shade: 0.4,
        mask: true,
        //maxmin: true, //开启最大化最小化按钮
        area: ['30vw', '50vh'],
        content: '/setpass?id=' + ID,
        shadeClose: false,
        cancel: function(){
            open_repass_index = 0
        }
    });
}

getinfo = function () {
    var token = window.localStorage.getItem("token")
    var table = layui.table;
    //展示已知数据
    table.render({
        elem: '#server_list'
        , id: 'server_list'
        , url: '/v1/userinfo'
        , headers: {
            Authorization: "Bearer " + token,
        }
        , title: '服务器列表'
        , cols: [[
            {field: 'ID', title: 'ID', sort: true, totalRow: true, hide: true}
            , {field: 'Ip', title: 'IP地址', sort: true, totalRow: true, edit: 'text'}
            , {field: 'Port', title: '端口', sort: true, totalRow: true, edit: 'text'}
            , {field: 'Username', title: '用户', sort: true, edit: 'text'}
            , {field: 'Nickname', title: '备注', edit: 'text'}
            , {field: 'BeforeTime', title: '上次登录', sort: true, totalRow: true}
            , {field: 'Controller', title: '操作'}
        ]]
        , skin: "row"
        , even: 'true'
        , page: true //是否显示分页
        , limits: [10, 15, 20]
        , limit: 10 //每页默认显示的数量
        , response: {
            statusCode: 200 //重新规定成功的状态码为 200，table 组件默认为 0
        }
        , parseData: function (res) { //将原始数据解析成 table 组件所规定的数据
            if (res.token) {
                window.localStorage.setItem("token", res.token)   //更新token
            }
            if (res.code == 301 || res.code == 302) {   //Token校验失败或过期
                layer.msg(res.msg, function () {
                    window.localStorage.clear();        //清空浏览器缓存
                    if (window != top) {
                        top.location.href = "/login";
                    }
                    window.location.href = "/login";
                })
                return
            }
            //console.log(res.data.List)
            for (key in res.data.List) {
                res.data.List[key].Controller = "<button onclick=\"open_terminal(" + res.data.List[key].ID + ",'" + res.data.List[key].Nickname + " - " + res.data.List[key].Ip + "'" + ")\" class=\"layui-btn-xs layui-btn-blue\">打开SSH终端</button>&nbsp;<button onclick=\"openpass(" + res.data.List[key].ID + ")\" class=\"layui-btn-xs layui-btn-warm\">重置密码</button>&nbsp;<button onclick=\"del_info(" + res.data.List[key].ID + ")\" class=\"layui-btn-xs layui-btn-danger\">删除</button>"
            }
            settmpdata(res);
            return {
                "code": res.code, //解析接口状态
                "msg": res.msg, //解析提示文本
                "count": res.data.Count, //解析数据长度
                "data": res.data.List //解析数据列表
            };
        },
    });
    table.on('edit(server_list)', function (obj) { //注：edit是固定事件名，test是table原始容器的属性 lay-filter="对应的值"
        //console.log(obj.value); //得到修改后的值
        //console.log(obj.field); //当前编辑的字段名
        //console.log(obj.data); //所在行的所有相关数据
        setnickname(obj.data)
    });
}

getinfo()

search = function () {
    var keyword = $('#keyword').val()
    if (keyword == "") {
        getinfo()
        return
    }
    var data = JSON.parse(window.localStorage.getItem("server"))
    var tmp_data = [];
    //data.data.Count = 1
    //data.data.List.splice(0, 1);
    data.data.List.forEach(function (e) {
        if (JSON.stringify(e).indexOf(keyword) != -1) {
            tmp_data.push(e);
        }
    });
    //console.log(tmp_data)
    if (tmp_data.length == 0) {
        layer.msg("搜索内容不存在")
        return
    }
    data.data.Count = tmp_data.length;
    data.data.List = tmp_data;
    var table = layui.table;
    table.render({
        elem: '#server_list'
        , id: 'server_list'
        , title: '服务器列表'
        , cols: [[
            {field: 'ID', title: 'ID', sort: true, totalRow: true, hide: true}
            , {field: 'Ip', title: 'IP地址', sort: true, totalRow: true, edit: 'text'}
            , {field: 'Port', title: '端口', sort: true, totalRow: true, edit: 'text'}
            , {field: 'Username', title: '用户', sort: true, edit: 'text'}
            , {field: 'Nickname', title: '备注', edit: 'text'}
            , {field: 'BeforeTime', title: '上次登录', sort: true, totalRow: true}
            , {field: 'Controller', title: '操作', width: 100}
        ]]
        , skin: "row"
        , even: 'true'
        , page: true //是否显示分页
        , limits: [10, 15, 20]
        , limit: 10 //每页默认显示的数量
        , response: {
            statusCode: 200 //重新规定成功的状态码为 200，table 组件默认为 0
        }
        , data: data.data.List
    });
}
var open_add_index = 0;
add = function () {
    if(open_add_index > 0){
        return false
    }
    open_add_index = layer.open({
        maxmin : true,
        type: 2,
        title: '添加SSH服务器',
        shade: 0.4,
        mask: true,
        //maxmin: true, //开启最大化最小化按钮
        area: ['30vw', '50vh'],
        content: '/add',
        shadeClose: false,
        cancel: function(){
            open_add_index = 0
        }
    });
}

del_info = function (ID) {
    http_send("/delete", {id: ID}, del_callbacl)
}

del_callbacl = function (result) {
    layer.msg(result.msg)
    getinfo()
}

open_terminal_index = 0;
open_terminal = function (ID, sername) {
    if(open_terminal_index > 0){
        return false
    }
    open_terminal_index = layer.open({
        maxmin : true,
        type: 2,
        title: '打开SSH终端：' + sername,
        shade: 0.4,
        mask: true,
        //maxmin: true, //开启最大化最小化按钮
        area: ['30vw', '30vh'],
        content: '/openterm?id=' + ID + "&sername=" + encodeURI(sername),
        shadeClose: false,
        cancel: function(){
            open_terminal_index = 0
        }
    });
}
