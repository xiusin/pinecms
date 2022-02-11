$("#key").keydown(function(e) {
    if (e.keyCode == 13) {
        openterm();
        return false
    }
});
openterm = function () {
    var pass_info = $('form').serializeArray();
    var obj = {}; //声明一个对象
    $.each(pass_info, function (index, field) {
        obj[field.name] = field.value; //通过变量，将属性值，属性一起放到对象中
    })
    obj.id = GetQueryString("id")
    http_send("/getterm", obj, openterm_callback)
}

openterm_callback = function (result) {
    parent.layer.closeAll('iframe');
    parent.open_terminal_index = 0;
    parent.getinfo();
    var getData = GetRequest();
    if (top) {
        top.layui.home.TabController.add(getData.sername, "/term?sid=" + result.data);
    } else {
        layer.msg("请不要单独打开本页面")
    }
}

GetRequest = function () {
    var url = decodeURI(decodeURI(location.search)); //获取url中"?"符后的字串，使用了两次decodeRUI解码
    var theRequest = new Object();
    if (url.indexOf("?") != -1) {
        var str = url.substr(1);
        strs = str.split("&");
        for (var i = 0; i < strs.length; i++) {
            theRequest[strs[i].split("=")[0]] = unescape(strs[i].split("=")[1]);
        }
        return theRequest;
    }
}