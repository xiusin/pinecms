aes_en = function (src, mkey) {
    var key_str = mkey;
    var plaintText = src;
    var iv = CryptoJS.enc.Utf8.parse(md5(key_str).substring(0, 16));  //24位字符串
    var key = CryptoJS.enc.Utf8.parse(md5(key_str).substring(0, 24));    // 16 24 32
    var encryptedData = CryptoJS.AES.encrypt(plaintText, key, {
        iv: iv,
        mode: CryptoJS.mode.CBC,
        padding: CryptoJS.pad.Pkcs7
    });
    return encryptedData
}

aes_de = function (enstr, mkey) {
    var key_str = mkey;
    var key = CryptoJS.enc.Utf8.parse(key_str);
    var iv = CryptoJS.enc.Utf8.parse(md5(key_str).substring(0, 16));  //16位字符串
    var dec = CryptoJS.AES.decrypt(enstr, key, {
        iv: iv,
        mode: CryptoJS.mode.CBC,
        padding: CryptoJS.pad.Pkcs7
    })
    CryptoJS.enc.Utf8.stringify(dec);
}


