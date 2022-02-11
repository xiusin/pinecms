window.onload = function () {
    var token = window.localStorage.getItem("token")
    if (token) {
        window.location.href = "/console"
    }
}

showproto = function () {
    var info = `
    <p style="text-indent:2em; ">本网站所提供的服务，只供参考学习之用。</p>
    
    <p style="text-indent:2em; ">本网站及其雇员一概毋须以任何方式就任何信息传递或传送的失误、不准确或错误对用户或任何其他人士负任何直接或间接的责任。</p>
    
    <p style="text-indent:2em; ">在法律允许的范围内，本网站在此声明,不承担用户或任何人士就使用或未能使用本网站所提供的信息或任何链接或项目所引致的任何直接、间接、附带、从属、特殊、惩罚性或惩戒性的损害赔偿（包括但不限于收益、预期利润的损失或失去的业务、未实现预期的节省）。</p>
    
    <p style="text-indent:2em; ">本网站所提供的信息，若在任何司法管辖地区供任何人士使用或分发给任何人士时会违反该司法管辖地区的法律或条例的规定或会导致本网站或其第三方代理人受限于该司法管辖地区内的任何监管规定时，则该等信息不宜在该司法管辖地区供该等任何人士使用或分发给该等任何人士。用户须自行保证不会受限于任何限制或禁止用户使用或分发本网站所提供信息的当地的规定。</p>
    
    <p style="text-indent:2em; ">本网站图片，文字之类版权申明，因为网站可以由注册用户自行上传图片或文字，本网站无法鉴别所上传图片或文字的知识版权，如果侵犯，请及时通知我们，本网站将在第一时间及时删除。</p>
    
    <p style="text-indent:2em; ">凡以任何方式登陆本网站或直接、间接使用本网站资料者，视为自愿接受本网站声明的约束。</p>`;
    layer.alert(info, {
        title: '服务协议',
        area: ['50vw'],
        skin: 'layui-layer-lan' //样式类名
        , closeBtn: 0
    })
}

login = function () {
    var phone = $('#phone').val();
    var code = $('#code').val();
    if (!(/^1(3|4|5|7|8)\d{9}$/.test(phone))) {
        errshow("手机号码有误，请重填！");
        return false;
    }
    if (isNaN(code) || code.length < 4) {
        errshow("验证码错误，请重试！");
        return false;
    }
    http_send("/login", {phone: phone, code, code}, login_callback)
    //加载层-风格2
    //layer.load(1);
}

login_callback = function (result) {
    layer.msg(result.msg);
    window.location.href = "/console"
}

send = function () {
    var phone = $('#phone').val();
    if (!(/^1(3|4|5|7|8)\d{9}$/.test(phone))) {
        errshow("手机号码有误，请重填！");
        return false;
    }
    http_send("/send", {phone: phone}, sendmsg_callback)
}

sendmsg_callback = function (result) {
    layer.msg(result.msg);
    countdown() //按钮禁用并显示倒计时
}


let countdowns = 120;
let button = $('#send');

countdown = function () {
    setTimeout(function () {
        if (countdowns <= 0) {
            button.prop('disabled', false);
            button.text("发送");
            countdowns = 120;//60秒过后button上的文字初始化,计时器初始化;
            return;
        } else {
            button.prop('disabled', true);
            button.text("(" + countdowns + "s)");
            countdowns--;
        }
        countdown()
    }, 1000) //每1000毫秒执行一次
}

errshow = function (msg) {
    layer.msg(msg, function () {
//关闭后的操作
    });
}