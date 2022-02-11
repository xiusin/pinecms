var sftp = null;
var sftp_ready = false;
var server_pwd = "";
var session_path = "";
var callback_success = undefined;
//var callback_error = undefined;
var upload_file_index = 0;

let sendmsg = {
    type: "",
    filepath: "",
    filename: "",
    filedata: "",
}

let msg_type = {
    connect: "connect",
    getpwd: "getpwd",
    upload: "upload",
}

let file_status = {
    success: 2, //上传成功
    error: -1,  //上传失败
}

create_sftp = function () {
    sftp = new WebSocket(ws_p + '://' + window.location.host + '/v1/sftp/' + GetQueryString("sid"));
    sftp.onopen = function () {
        sftp.send(JSON.stringify(auth));  //验证权限

        sftp.onmessage = default_msg;
        sftp.onerror = function (e) {
            sftp_ready = false
            console.log(e);
            layer.msg("SFTP连接出错")
        };

        sftp.onclose = function (e) {
            sftp_ready = false
            console.log(e);
            layer.msg("SFTP连接断开")
        };
    };
}

default_msg = function (msg) {
    if (isJSON(msg.data)) {
        msg = JSON.parse(msg.data)
        if (msg.type == msg_type.connect && msg.code == 200) {
            sftp_ready = true
            server_pwd = msg.data   //初始化工作路径
        } else if (msg.type == msg_type.upload) {
            if (msg.code == 200) {
                callback_success(upload_file_index, file_status.success)
            } else {
                //layer.msg("文件上传失败：" + msg.msg)      //可以回调修改状态文字内容，这里为了图方便，暂时不写了；（觉得不妥，又写上了）
                callback_success(upload_file_index, file_status.error,msg.msg)  //上传失败正常回调，开始处理下一个
            }
        }
    }
}

getpwd = function (callback) {
    sendmsg.type = msg_type.getpwd
    sftp.onmessage = function (msg) {
        if (isJSON(msg.data)) {
            msg = JSON.parse(msg.data)
            if (msg.code == 200) {
                if (session_path == "") {
                    session_path = server_pwd
                }
                callback(msg)
            } else {
                layer.msg(msg.msg)
            }
            sftp.onmessage = default_msg
        }
    }
    sftp.send(JSON.stringify(sendmsg))
}

sendFile = function (fileinfo, fileData, file_index, callback) {
    sendmsg.type = msg_type.upload
    sendmsg.filepath = session_path + fileinfo.path
    sendmsg.filename = fileinfo.name
    sendmsg.filedata = fileData
    callback_success = callback
    upload_file_index = file_index
    sftp.send(JSON.stringify(sendmsg))
}

isJSON = function (str) {
    if (typeof str == 'string') {
        try {
            var obj = JSON.parse(str);
            if (typeof obj == 'object' && obj) {
                return true;
            } else {
                return false;
            }

        } catch (e) {
            //console.log('error：' + str + '!!!' + e);
            return false;
        }
    }
    //console.log('It is not a string!')
}