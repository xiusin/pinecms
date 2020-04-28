var fontSize = $.cookie('fontSize'),
	textIndent = $.cookie('textIndent');

/* 字号设置 */

if(fontSize !== null){
	$('.post_content').addClass(fontSize);
}

$('#fs1').click(function(){
	$('.post_content').removeClass('fs1 fs2 fs3');
	$('.post_content').addClass('fs1');
	$.cookie('fontSize', null,{expires: 9999,path:'/'});
});
$('#fs2').click(function(){
	$('.post_content').removeClass('fs1 fs2 fs3');
	$.cookie('fontSize', null);
	$('.post_content').addClass('fs2');
	if( fontSize !== 'fs2' ){
		$.cookie('fontSize', 'fs2',{expires: 9999,path:'/'});
	}
});
$('#fs3').click(function(){
	$('.post_content').removeClass('fs1 fs2 fs3');
	$.cookie('fontSize', null);
	$('.post_content').addClass('fs3');
	if( fontSize !== 'fs3' ){
		$.cookie('fontSize', 'fs3',{expires: 9999,path:'/'});
	}
});

/* 段落缩进设置 */
if(textIndent !== null){
	$('.post_content').addClass('indt');
	$('#indt').addClass('hide');
	$('#noindt').removeClass('hide');
}else{
	$('.post_content').removeClass('indt');
	$('#indt').removeClass('hide');
	$('#noindt').addClass('hide');
}

$('#indt').click(function(){
	$('#noindt').removeClass('hide');
	$('#indt').addClass('hide');
	$('.post_content').addClass('indt');
	$.cookie('textIndent', 'indt',{expires: 9999,path:'/'});
});
$('#noindt').click(function(){
	$('#noindt').addClass('hide');
	$('#indt').removeClass('hide');
	$('.post_content').removeClass('indt');
	$.cookie('textIndent', null,{expires: 9999,path:'/'});
});