GetQueryString = function (name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]);
    return null;
}

default_callback = function (result) {
    layer.msg(result.msg)
}

http_send = function (url, data, callback = default_callback, method = "POST") {
    layer.load(1);
    url = "/v1" + url;
    $.ajax({
        url: url,
        type: method,
        data: data,
        beforeSend: function (xhr) {
            var token = window.localStorage.getItem("token")
            if (token) {
                xhr.setRequestHeader("Authorization", "Bearer " + token);
            }
        },
        success: function (result) {
            layer.closeAll('loading');
            if (result.token) {
                window.localStorage.setItem("token", result.token)   //更新token
            }
            if (result.code == 301 || result.code == 302) {   //Token校验失败或过期
                layer.msg(result.msg, function () {
                    window.localStorage.clear();        //清空浏览器缓存
                    if (window != top) {
                        top.location.href = "/login";
                    }
                    window.location.href = "/login";
                })
            }
            if (result.code != 200) {
                layer.msg(result.msg, function () {

                })
            } else {
                callback(result);
            }
        },
        error: function (e) {
            console.log("err")
            layer.closeAll('loading');
            console.log(e.status);
            console.log(e.responseText)
            layer.msg(e.responseText);
        },
        complete: function (e) {
            layer.closeAll('loading');
            console.log("请求完成")
        }
    });
}