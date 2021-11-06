function highlightBlock() {
    if('hljs' in window) {
        $('.markdown-body pre code').each(function(i, block) {
            hljs.highlightBlock(block)
        })
    }
}
$(function() {
	$(".doc-nav > ul").addClass("menus");
	$(".doc-nav > ul > li").addClass("menu");
	$(".doc-nav > ul > li > ul").addClass("pages");
	$(".doc-nav > ul > li > ul > li").addClass("page");
	const as = $(".doc-nav > ul a");
	as.each((index) => {
		if (as.eq(index).attr("href").indexOf(currentDocmd) > -1) {
			as.eq(index).parent().addClass('active');
			as.eq(index).parents(".menu").addClass("active");
		}
	})

	var onleft=1;//这个是移动端的变量
	
    // ie8 以上使用代码高亮
    highlightBlock()
    // 现代浏览器采用前端路由
    if(window.history && history.pushState) {

     /*
		window.addEventListener("popstate", function(e) {
            changePage()
        })
        $("ul.menus").on('click', 'li.page>a', function(e) {
            var href = e.target.href
            history.pushState(null, null, href)
            changePage()
            if($(window).width()<740){
    			$(".doc-content").css("left","0")
    			onleft=1;
    		}
            return false
        })
        function changePage() {
            var pathname = location.pathname
            var title = ''
            $('li.page').removeClass('active')
            var $pageItem = $('li.page > a[href="' + pathname + '"]')
            if($pageItem) {
                $pageItem.parent().addClass('active')
                title = $pageItem.text()
            } else {
                $pageItem = $('li.page').eq(0).addClass('active').children('a')
                pathname = $pageItem.attr('href')
                title = $pageItem.text()
            }
            var $title = $('title')
            $title.text($title.text().replace(/-\s.+/, '- ' + title))
            var mdPath = pathname.replace(/(\/\S+?\/)/, '$1pages/') + '.md'
            $.ajax({
                url: mdPath,
                success: function(data) {
                    $('#markdown').html(marked(data)).scrollTop(0)
                    highlightBlock()
                }
            })
        }*/
    }
    
    //移动端页面切换
    $(".nav_bar em").on("click",function(event){
    	event.stopPropagation();
 		$(".doc-header").addClass("bounceInDown");
    	if(onleft){
    		$(".doc-content").css("left",$(".doc-nav").width());
    		onleft=0;
    	}else{
    		$(".doc-content").css("left","0");
    		onleft=1;
    	}
    });
    $("body,.doc-content").click(function(event){
    	event.stopPropagation();
    	if($(window).width()<740){
    		$(".doc-content").css("left","0");
    		onleft=1;
    	}
    });
    $(".menu_top").click(function(event){
//  	event.stopPropagation();
    	$(".doc-header").stop().show();
    })
    
    
    //判断浏览器
	var userAgent = navigator.userAgent, 
	rMsie = /(msie\s|trident.*rv:)([\w.]+)/, 
	rFirefox = /(firefox)\/([\w.]+)/, 
	rOpera = /(opera).+version\/([\w.]+)/, 
	rChrome = /(chrome)\/([\w.]+)/, 
	rSafari = /version\/([\w.]+).*(safari)/;
	
	var version;
	var ua = userAgent.toLowerCase();
	function uaMatch(ua) {
		var match = rMsie.exec(ua);
		if (match != null) {
			return { browser : "IE", version : match[2] || "0" };
		}
		var match = rFirefox.exec(ua);
		if (match != null) {
			return { browser : match[1] || "", version : match[2] || "0" };
		}
		var match = rOpera.exec(ua);
		if (match != null) {
			return { browser : match[1] || "", version : match[2] || "0" };
		}
		var match = rChrome.exec(ua);
		if (match != null) {
			return { browser : match[1] || "", version : match[2] || "0" };
		}
		var match = rSafari.exec(ua);
		if (match != null) {
			return { browser : match[2] || "", version : match[1] || "0" };
		}
		if (match != null) {
			return { browser : "", version : "0" };
		}
	}
	var browserMatch = uaMatch(userAgent.toLowerCase());
	if (browserMatch.browser) {
		browser = browserMatch.browser;
		version = browserMatch.version;
	}
	if(browser=="IE"){
		if(version=="10.0"||version=="9.0"||version=="8.0"||version=="7.0"){
			$(".doc-content").height($("html").height()-$(".doc-header").height());
		}else if(version=="9.0"){
//			$(".doc-content").height($("html").height()-$(".doc-header").height());
		}
	}
	
    
})