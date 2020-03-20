$(function () {
    click_small_btn();
    header_search_btn();
    mouse_scroll_change();
    banner_img_url_click();
    click_show_hide_article_btn_box_item();
    mousehover_changeimgsrc();
});
//记录鼠标滑轮事件 显示 一键向上的按钮  和把个人介绍固定
function mouse_scroll_change() {
    if($(".self_info").length>0){
        $(window).load(function () {
        $(".self_info_scroll_bottom").html($(".self_info").html());
        var right_height=$(".right").height();
        var right_right=($(window).width()-$(".article_box").width())/2;
        // var right_left=$(".right").offset().left
        // console.log("right的高度现在是"+($(".right").offset().top+right_height))
        $(window).scroll(function () {
            // console.log("页面已经被吞入"+$(window).scrollTop())
            // console.log("right的高度现在是"+($(".right").offset().top+right_height))
            if($(window).scrollTop()>($(".right").offset().top+right_height)){
                $(".self_info_scroll_bottom").stop().animate({top:0,right:right_right,opacity: 1});
            }else if($(window).scrollTop()<=($(".right").offset().top+right_height)) {
                $(".self_info_scroll_bottom").stop().animate({top:0,right:0,opacity: 0});
            }
        })
    })
    }

    if($(".about_me_info").length>0){
        $(window).load(function () {
            $(".about_me_info_scroll_bottom").html($(".about_me_info").html());
            var right_height=$(".right").height();
            var right_right=($(window).width()-$(".article_box").width())/2;
            $(window).scroll(function () {
                if($(window).scrollTop()>($(".right").offset().top+right_height)){
                    $(".about_me_info_scroll_bottom").stop().animate({top:0,right:right_right,opacity: 1});
                }else if($(window).scrollTop()<=($(".right").offset().top+right_height)) {
                    $(".about_me_info_scroll_bottom").stop().animate({top:0,right:0,opacity: 0});
                }
            })
        })
    }

}
//学习笔记btn按钮显示影藏下面的文章列表
function click_show_hide_article_btn_box_item() {
    $(".article_btn ul li").click(function () {
            $(".article_btn_box>div").eq($(this).index()).show().siblings().hide();
            $(this).css({"border-bottom":"1px solid #004FCB","color":"#004FCB"}).siblings().css({"border-bottom":"none","color":"#333"})


            //当前盒子变量因为有4个盒子所有用个变量记录
            var _thisitem=$(".article_btn_box>div").eq($(this).index());

            // console.log(_thisitem.find("ul li:first").find("a").length)
            //
            // console.log(_thisitem.find("ul li:first").find("a").attr("imgsrc"))

            if(_thisitem.find("ul li:first").find("a").length>0){

            _thisitem.find(".article_btn_box_item_img").find("img").attr("src",_thisitem.find("ul li:first").find("a").attr("imgsrc"));
            _thisitem.find(".article_btn_box_item_img").attr("href",_thisitem.find(".article_btn_box_item_img").find("ul li:first").find("a").attr("href"));
      }

    });
}
//学习笔记下面的li鼠标悬停改变左侧图片的路劲
function mousehover_changeimgsrc() {
        //第一个盒子
        var _first_box=$(".article_btn_box .article_btn_box_item:first");
        // console.log($(".article_btn_box .article_btn_box_item:first").find("img").attr("src",$(".article_btn_box .article_btn_box_item:first").find("ul li:first").find("a").attr("imgsrc")));
        _first_box.find(".article_btn_box_item_img").find("img").attr("src",_first_box.find("ul li:first").find("a").attr("imgsrc"));
        _first_box.find(".article_btn_box_item_img").attr("href",_first_box.find("ul li:first").find("a").attr("href"));
        var _itemfather=$(".article_btn_box .article_btn_box_item");
        _itemfather.find("ul li").hover(function () {
                // console.log($(this).parent().parent().html())
                $(this).parent().parent().find("img").attr("src",$(this).find("a").attr("imgsrc"));
                $(this).parent().parent().find(".article_btn_box_item_img").attr("href",$(this).find("a").attr("href"));
        },function () {
                $(this).parent().parent().find("img").attr("src",$(this).parent().find("li:first").find("a").attr("imgsrc"));
                $(this).parent().parent().find(".article_btn_box_item_img").attr("href",$(this).parent().find("li:first").find("a").attr("href"));
        });
}
//手机端页面按钮点击三横变一×
function click_small_btn() {
    $(".header_nav_small_btn").click(function () {
        $(this).toggleClass("click_small_btn");
        //手机端导航点击显示隐藏
        $(".header_nav").toggle("slow",function () {
            if($(".header_nav_small_btn").css("display")=="block"){
                if($(".search_box").css("display")=="block"){
                    $(".search_box").hide(350);
                }
            }
        });
    });
}
//头部搜索按钮
function header_search_btn() {
        $(".header_search_btn").click(function () {
            $(".search_box").slideToggle(300,function () {
                if($(".header_nav_small_btn").css("display")=="block"){
                        if($(".header_nav").css("display")=="block"){
                            $(".header_nav_small_btn").removeClass("click_small_btn");
                            $(".header_nav").hide(350);
                    }
                }
            });
        });
}
//banner图下url点击
function banner_img_url_click() {
    $(".banner_box img").click(function () {
        window.location.href = $(this).attr('url');
    })
}
