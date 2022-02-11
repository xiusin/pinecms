// var is_login = { watchValue:false };
// var lastTimeValue=is_login.watchValue;
// Object.defineProperty(is_login, 'watchValue', {
//     get: function() {
//         console.log('get：' + watchValue);
//         return watchValue;
//     },
//     set: function(value) {
//         watchValue = value;
//         if(lastTimeValue!=watchValue){
//             lastTimeValue=watchValue;
//             console.log('value changed!! set: ' + watchValue);
//             if(watchValue == true){
//                 create_sftp()
//             }
//         }
//     }
// });

var is_login = false
const protocol = document.location.protocol.split(':')[0];
var ws_p = "ws";
if (protocol == "https") {
    ws_p = "wss";
}
const token = window.localStorage.getItem("token")
if (token == "") {
    if (window != top) {
        top.location.href = "/login";
    }
    window.location.href = "/login";
}
const auth = {
    type: "auth",
    token: token,
}

function completeLoading() {
    console.log("Loading Success")
    if (document.readyState == "complete") {
        layer.close(index);
    }
}

const socket = new WebSocket(ws_p + '://' + window.location.host + '/v1/term/' + GetQueryString("sid"));

const term = new Terminal({cols: 180, rows: 50, screenKeys: true, cursorBlink: true, cursorStyle: "block"});
term.open(document.getElementById('terms'));
window.onresize = function () {
    fit.fit(term);
};
socket.onopen = function () {
    socket.send(JSON.stringify(auth));  //验证权限
    term.write("正在验证\r\n");
    term.toggleFullscreen(true);
    fit.fit(term);
    term.on('data', function (data) {
        let sdata = {
            type: "cmd",
            cmd: data,
        }
        socket.send(str2utf8(JSON.stringify(sdata)));
    });

    term.on('resize', size => {
        //console.log('resize', [size.cols, size.rows]);
        let sdata = {
            type: "resize",
            cols: size.cols,
            rows: size.rows,
        }
        socket.send(JSON.stringify(sdata));
    });

    socket.onmessage = function (msg) {
        if (!is_login) {
            term.clear()
            create_sftp()
            is_login = true
        }
        let reader = new FileReader();
        reader.onload = function (event) {
            let content = reader.result;//内容就在这里
            //delete reader
            //console.log(content)
            term.write(content);
            update_path(content);
        };
        reader.readAsText(msg.data);
        //term.write(msg.data);

    };
    socket.onerror = function (e) {
        is_login = false
        layer.msg("链接出错：" + JSON.stringify(e))
        console.log(e);
    };

    socket.onclose = function (e) {
        is_login = false
        layer.msg("链接断开：" + JSON.stringify(e))
        term.write("连接已断开" + "\r\n");
        //term.destroy();
    };
};

function str2utf8(str) {
    let encoder = new TextEncoder('utf8');
    return encoder.encode(str);
}

function update_path(str) {     //判断是否有效路径后更新当前所在路径，用于SFTP功能
    //console.log(str)
    let splice = String.fromCharCode(7)
    let start = str.indexOf(":")
    let end = str.indexOf(splice)
    if (start >= 0 && end >= 0) {
        let path = trimStr(str.substring(start + 1, end))
        let verify = path.indexOf(" ")
        if (verify == -1) {
            //console.log(path)
            if (path.substr(0, 1) == "~" || path.substr(0, 1) == "/") {
                if (path.substr(0, 1) == "~") {    //替换~为用户目录
                    path = server_pwd + path.substr(1)
                }
                //console.log(path)
                if (session_path != path) {
                    session_path = path
                }
            }
        }
    }
}

function trimStr(str) {
    return str.replace(/(^\s*)|(\s*$)/g, "");
}