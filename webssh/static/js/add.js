show = function () {
    layer.tips('本程序承诺服务器绝不会保存您的密码信息，所以添加服务器需要输入你密码的加密密码，此密码用于将您的SSH密码加密后存入数据库，加密密码仅存放于您的浏览器缓存，用于解密密码建立SSH连接。', '#show', {
        tips: [1, '#3595CC'],
        time: 1000 * 10
    });
}

$("#passwd").bind("input propertychange", function (event) {
    //console.log($("#key").val())
    var key = $('#key').val();
    if (key == "") {
        return
    }
    var pass = $('#passwd').val();
    $('#aes').val(aes_en(pass, key));
});

$("#key").bind("input propertychange", function (event) {
    //console.log($("#key").val())
    var key = $('#key').val();
    var pass = $('#passwd').val();
    if (pass == "") {
        return
    }
    //console.log(aes_en(pass,key))
    $('#aes').val(aes_en(pass, key));
});

addser = function () {
    var addinfo = $('form').serializeArray();
    var obj = {}; //声明一个对象
    $.each(addinfo, function (index, field) {
        obj[field.name] = field.value; //通过变量，将属性值，属性一起放到对象中
    })
    if (!isValidIP(obj.ip)) {
        layer.msg("IP格式错误", function () {

        })
        return
    }
    obj.password = obj.desc
    delete obj.setpass; //不提交用户的加密密码
    console.log(obj)
    http_send("/addser", obj, add_callback)
}

editpass = function () {
    var addinfo = $('form').serializeArray();
    var obj = {}; //声明一个对象
    $.each(addinfo, function (index, field) {
        obj[field.name] = field.value; //通过变量，将属性值，属性一起放到对象中
    })
    obj.id = GetQueryString("id")
    obj.password = obj.desc
    delete obj.setpass; //不提交用户的加密密码
    console.log(obj)
    parent.open_repass_index = 0;
    http_send("/repass", obj)
}

add_callback = function (result) {
    parent.open_add_index = 0;
    parent.getinfo();
    parent.layer.closeAll('iframe');
}

isValidIP = function (ip) {
    var reg = /^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/
    return reg.test(ip);
}