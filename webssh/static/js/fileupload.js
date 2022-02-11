let dropbox = document.getElementById("terms");       //要监听拖动上传的节点

let fileDrop = {
    startTime: 0,
    endTime: 0,
    uploadLength: 0, //上传数量
    //splitSize: 1024 * 1024 * 2, //文件上传分片大小
    filesList: [], // 文件列表数组
    errorLength: 0, //上传失败文件数量
    isUpload: true, //上传状态，是否可以上传
    uploadSuspend: false,  //上传暂停参数
    isUploadNumber: 150,//限制单次上传数量
    uploadAllSize: 0, // 上传文件总大小
    uploadedSize: 0, // 已上传文件大小
    topUploadedSize: 0, // 上一次文件上传大小
    uploadExpectTime: 0, // 预计上传时间
    //initTimer:0, // 初始化计时
    speedInterval: null, //平局速度定时器
    timerSpeed: 0, //速度
    uploading: false,
    cancel: false,
    done: false,
}

function file_init() {
    fileDrop = {
        startTime: 0,
        endTime: 0,
        uploadLength: 0, //上传数量
        //splitSize: 1024 * 1024 * 2, //文件上传分片大小
        filesList: [], // 文件列表数组
        errorLength: 0, //上传失败文件数量
        isUpload: true, //上传状态，是否可以上传
        //uploadSuspend:[],  //上传暂停参数
        isUploadNumber: 800,//限制单次上传数量
        uploadAllSize: 0, // 上传文件总大小
        uploadedSize: 0, // 已上传文件大小
        topUploadedSize: 0, // 上一次文件上传大小
        uploadExpectTime: 0, // 预计上传时间
        //initTimer:0, // 初始化计时
        speedInterval: null, //平局速度定时器
        timerSpeed: 0, //速度
        uploading: false,
        cancel: false,
    }
}

dropbox.addEventListener("dragleave", function (e) {
    //e.stopPropagation();
    e.preventDefault();
}, false);

dropbox.addEventListener("dragenter", function (e) {
    //e.stopPropagation();
    e.preventDefault();
}, false);

dropbox.addEventListener("dragover", function (e) {
    //e.stopPropagation();
    e.preventDefault();
}, false);

dropbox.addEventListener("drop", changes, false);

function changes(e) {
    if (!sftp_ready) {
        layer.msg("请等待服务器连接！")
        return false
    }
    e.preventDefault();
    let items = e.dataTransfer.items, time, num = 0
    if (fileDrop.uploading) {
        layer.msg("已有文件队列上传中")
        return false
    }
    if (items && items.length && items[0].webkitGetAsEntry != null) {
        if (items[0].kind != 'file') return false;
    }
    if (fileDrop.filesList == null) fileDrop.filesList = []
    for (let i = fileDrop.filesList.length - 1; i >= 0; i--) {
        if (fileDrop.filesList[i].is_upload) fileDrop.filesList.splice(-i, 1)
    }

    function update_sync(s) {
        s.getFilesAndDirectories().then(function (subFilesAndDirs) {
            return iterateFilesAndDirs(subFilesAndDirs, s.path);
        });
    }

    let iterateFilesAndDirs = function (filesAndDirs, path) {
        for (let i = 0; i < filesAndDirs.length; i++) {
            if (typeof (filesAndDirs[i].getFilesAndDirectories) == 'function') {
                update_sync(filesAndDirs[i])
            } else {
                if (num > fileDrop.isUploadNumber) {
                    //fileDrop.isUpload = false;
                    layer.msg(' ' + fileDrop.isUploadNumber + '份，无法上传,请压缩后上传!。', {icon: 2, area: '405px'});
                    //clearTimeout(time);
                    file_init()
                    return false;
                }
                fileDrop.filesList.push({
                    file: filesAndDirs[i],
                    path: path,
                    name: filesAndDirs[i].name.replace('//', '/'),
                    local: (path == "/" ? "" : path) + "/" + filesAndDirs[i].name.replace('//', '/'),
                    size: to_size(filesAndDirs[i].size),
                    upload: 0, //上传状态,未上传：0、上传中：1，已上传：2，上传失败：-1
                    is_upload: false
                });
                fileDrop.uploadAllSize += filesAndDirs[i].size
                fileDrop.uploadLength++;
            }
        }
    }
    if ('getFilesAndDirectories' in e.dataTransfer) {
        e.dataTransfer.getFilesAndDirectories().then(function (filesAndDirs) {
            return iterateFilesAndDirs(filesAndDirs, '/');
        });
    }
    //console.log(fileDrop.filesList)
    layer.load(1, {
        shade: [0.1, '#fff'] //0.1透明度的白色背景
    });
    getpwd(getpwd_callback) //采用回调函数的方式检查SFTP服务是否可用
}

function getpwd_callback(msg) {
    server_pwd = msg.data
    setTimeout(function () {
        layer.closeAll('loading')
        open_upload_window()
    }, 1500)
}

function open_upload_window() {
    //console.log(fileDrop.filesList[0])
    let template = `
    <table class="layui-table" lay-even="" lay-skin="row" id="file_upload" style="table-layout: fixed;padding-top: 0">
      <colgroup>
        <col width="35%">
        <col width="35%">
        <col width="30%">
        <col>
      </colgroup>
      <thead>
         <tr>
          <th>文件路径（共` + fileDrop.uploadLength + `个文件）</th>
          <th>文件大小（共` + to_size(fileDrop.uploadAllSize) + `）</th>
          <th>文件状态</th>
        </tr>
      </thead>
      <tbody align="center">
        
      </tbody>
    </table>
    `
    layer.open({
        type: 1,
        closeBtn: 1,
        maxmin: true,
        area: ['550px', '455px'],
        btn: [fileDrop.uploadSuspend ? "继续上传" : "开始上传", '取消上传'],
        title: '上传文件到：' + session_path,
        skin: 'file_dir_uploads',
        shade: 0.4,
        shadeClose: false,
        content: template,
        success: function () {
            for (let i = 0; i < fileDrop.filesList.length; i++) {
                $("#file_upload tbody").append(create_row(i, fileDrop.filesList[i]));
            }
        },
        yes: function (index) {
            if (fileDrop.filesList.length <= 0) {
                layer.msg("请选择要上传的文件！", {icon: 0})
                file_init()
                layer.close(index)
                return false
            }
            if (fileDrop.uploading && fileDrop.uploadSuspend == false) {
                return false
            }
            //console.log("开始上传")
            $('.layui-layer-btn0').css({
                'cursor': 'no-drop',
                'background': '#4e8ccc'
            }).attr('data-upload', 'true').text('上传中');
            fileDrop.uploadSuspend = false
            fileDrop.uploading = true
            start_upload()
        },
        btn2: function (index) {
            if(fileDrop.done){
                layer.close(index);
                file_init()
                return false
            }
            if (fileDrop.uploading) {
                layer.confirm('是否取消上传当前列表的文件，若取消上传，已上传的文件，需用户手动删除，是否取消上传？', {
                    title: '取消上传文件',
                    icon: 0,
                    btn: ['取消上传', '继续上传']
                }, function (indexs) {
                    layer.close(index);
                    layer.close(indexs);
                    file_init()
                }, function () {
                    fileDrop.uploadSuspend = true   //开启暂停标识，用于继续上传功能
                    open_upload_window()
                });
                //return false;
            } else {
                layer.close(index);
                file_init()
            }
        },
        cancel: function (index) {
            if(fileDrop.done){
                layer.close(index);
                file_init()
                return false
            }
            if (fileDrop.uploading) {
                layer.confirm('是否取消上传当前列表的文件，若取消上传，已上传的文件，需用户手动删除，是否取消上传？', {
                    title: '取消上传文件',
                    icon: 0,
                    btn: ['取消上传', '继续上传']
                }, function (indexs) {
                    layer.close(index);
                    layer.close(indexs);
                    file_init()
                }, function () {
                    fileDrop.uploadSuspend = true   //开启暂停标识，用于继续上传功能
                    open_upload_window()
                });
                //return false;
            } else {
                layer.close(index);
                file_init()
            }
        },
    });
}

function start_upload(file_index, status, msg) {
    file_index = file_index == null ? 0 : file_index
    if (status) { //这个参数只在上传回调中传递，用于更新文件上传结果，如果这个参数存在，表示file_index已经走完上传流程，故file_index需自增1，代表开始下一个文件的上传
        // if(msg){
        //     upload_view(file_index,msg)
        // }else{
        //     upload_view(file_index,status)
        // }
        fileDrop.filesList[file_index].upload = status  //更新文件列表中的状态
        upload_view(file_index, msg == undefined ? status : msg)
        file_index++
    }
    let reader = new FileReader();  //这个可以选择函数外声明，可避免重复创建
    if (reader == undefined) {
        file_init()
        layer.msg("您的浏览器暂不支持在线上传文件！")
        return false
    }
    reader.onload = function () {
        let fileData = _arrayBufferToBase64(this.result);
        fileDrop.filesList[file_index].upload = 1    //修改状态为上传中
        upload_view(file_index, 1)   //修改视图内容
        sendFile(fileDrop.filesList[file_index], fileData, file_index, start_upload)
        //console.log(fileData);
    }
    if (file_index < fileDrop.filesList.length) {     //采用回调方式上传，进度更可控
        if (fileDrop.filesList[file_index].upload < 2) {  //可继续上传未上传的文件
            let f_id = "#file_" + file_index + " td";
            $(f_id)[2].innerHTML = "<font color='#808080'>解析中</font>";
            reader.readAsArrayBuffer(fileDrop.filesList[file_index].file)
        }
    } else {
        layer.msg("上传完成")
        $('.layui-layer-btn0').text('上传完成');
        $('.layui-layer-btn1').text('关闭窗口');
        fileDrop.done = true
    }
    // for (let i = 0; i < fileDrop.filesList.length; i++) {
    //     reader.readAsArrayBuffer(fileDrop.filesList[i].file);
    // }
}

function upload_view(file_index, status) {
    let f_id = "#file_" + file_index + " td";
    $(f_id)[2].innerHTML = getstatu(status);
}

function _arrayBufferToBase64(buffer) {
    let binary = '';
    let bytes = new Uint8Array(buffer);
    let len = bytes.byteLength;
    for (var i = 0; i < len; i++) {
        binary += String.fromCharCode(bytes[i]);
    }
    return window.btoa(binary);
}

function create_row(index, file) {
    //console.log(file)
    return "<tr id='file_" + index + "'><td title='" + file.local + "' style=\"white-space:nowrap;overflow:hidden;text-overflow: ellipsis;\">" + file.local + "</td> <td>" + file.size + "</td> <td>" + getstatu(file.upload) + "</td></tr>"
}

function getstatu(statu) {
    //上传状态,未上传：0、上传中：1，已上传：2，上传失败：-1
    if (statu == -1) {
        return "<font color='red'>上传失败</font>"
    } else {
        if (statu == 0) {
            return "<font color='black'>未上传</font>"
        } else if (statu == 1) {
            return "<font color='#808080'>上传中</font>"
        } else if (statu == 2) {
            return "<font color='green'>已上传</font>"
        } else {
            return "<font color='red'>" + statu + "</font>"
        }
    }
}

function to_size(a) {
    var d = [" B", " KB", " MB", " GB", " TB", " PB"];
    var e = 1024;
    for (var b = 0; b < d.length; b += 1) {
        if (a < e) {
            var num = (b === 0 ? a : a.toFixed(2)) + d[b];
            return (!isNaN((b === 0 ? a : a.toFixed(2))) && typeof num != 'undefined') ? num : '0B';
        }
        a /= e
    }
}