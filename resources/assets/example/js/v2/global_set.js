var fontFamily = $.cookie('fontFamily');

if(fontFamily !== null){
	$('body').addClass('song');
	$('#yahei').removeClass('hide');
	$('#song').addClass('hide');
}else{
	$('body').removeClass('song');
	$('#yahei').addClass('hide');
	$('#song').removeClass('hide');
}

$('#song').click(function(){
	$('#yahei').removeClass('hide');
	$('#song').addClass('hide');
	$('body').addClass('song');
	$('body').addClass('song');
	$.cookie('fontFamily', 'song',{expires: 9999,path:'/'});
});

$('#yahei').click(function(){
	$('#song').removeClass('hide');
	$('#yahei').addClass('hide');
	$('body').removeClass('song');
	$.cookie('fontFamily', null,{expires: 9999,path:'/'});
});

function modifyheight() { 
 $("#top_hl").css({"height":"60px"});
} 