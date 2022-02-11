/**
 *  入口文件索引
 *  使用说明：将此文件引入到页面中，可在script标签上定义一个data-main=""属性，
 *  此属性指定页面入口文件。
 *
 **/
(function () {

    var entry,
        // 配置所有应用的入口文件，程序将会按照data-main属性中设置的值进行索引查找
        // 如果你在引入此脚本的script标签上没有设置data-main属性，程序将会默认访问home.js文件
        app = {
            home: '{/}static/js/home'
        };

    (function () {

        var dataMain, scripts = document.getElementsByTagName('script'),
            eachScripts = function (el) {
                dataMain = el.getAttribute('data-main');
                if (dataMain) {
                    entry = dataMain;
                }
            };

        [].slice.call(scripts).forEach(eachScripts);

    })();

    layui.config({
        base: '/static/assets/lay/modules/'
    }).extend(app).use(entry || 'home');

})();