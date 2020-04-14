$(document).ready(function () {
    $("#btnComment").click(function () {
        postdata();
    });
    $("#btnReplyComment").click(function () {
        self.parent.tb_remove();
        postdata();
    });
});


 function postdata(){
     $.ajax({
         type: "POST",
         url: "/ithome/postComment.aspx",
         data: "newsid=" + $("#newsid").val() + "&commentNick=" + escape($("#commentNick").val()) + "&commentContent=" + escape($("#commentContent").val()) + "&parentCommentID=" + $("#parentCommentID").val() + "&txtCode=" + $("#txtCode").val() + "&type=comment",
         success: function (msg) {

             var messageobj = GetObj('commentMessage');

             messageobj.innerHTML = '<span style="color:red">' + msg + '</span>';
             if (msg.indexOf("评论成功") >= 0) {
                 self.parent.LoadData();
                 var comment = GetObj("commentContent");
                 comment.value = '';

                 setCookie('username', $("#commentNick").val(), 1);

             }
             // randomNoImg.src = 'http://www.ithome.com/validate.aspx?' + Math.random();
         }
     });
 }

 function GetObj(objName) {
     if (document.getElementById) {
         return eval('document.getElementById("' + objName + '")');
     } else if (document.layers) {
         return eval("document.layers['" + objName + "']");
     } else {
         return eval('document.all.' + objName);
     }
 }

 function ShowReplay(commentid,newsid) {

     var ReplyDiv = GetObj('Reply' + commentid);
     ReplyDiv.style.display = 'block';


     if (ReplyDiv.innerHTML.length == 0) {

            replydivHtml = '  <div class="reply_post_comment">';
            replydivHtml += '	<a href="javascript:void(0)" class="close_comm" onclick="CloseReplay(' + commentid + ')"></a>';
            replydivHtml += '  <div class="add_comm">';
           //replydivHtml += '  <p class="c_alt">请自觉遵守互联网相关的政策法规，严禁发布色情、暴力、反动的言论。</p>';
			replydivHtml += '  <textarea name="txtBody" id="commentContent' + commentid + '" cols="60" rows="5" class="ipt-txt" onfocus="if(this.value==\'IT之家有您参与更精彩！\')this.value=\'\';" onblur="if(this.value==\'\')this.value=\'IT之家有您参与更精彩！\';">IT之家有您参与更精彩！</textarea>';
			replydivHtml += '  <div class="comm-con"><span>签名：</span>  ';
			 if(getCookie("username")=='')
			replydivHtml += '  <input name="txtNickname" type="text" id="commentNick' + commentid + '" size="36" class="ipt-txt" value="匿名" onfocus="if(this.value==\'匿名\')this.value=\'\';" onblur="if(this.value==\'\')this.value=\'匿名\';" />';
			else
			replydivHtml += '  <input name="txtNickname" type="text" id="commentNick' + commentid + '" size="36" class="ipt-txt" value="' + getCookie("username") + '" />';
			replydivHtml += '  </div>';
			//replydivHtml += '  <span id="divValidate"><input name="txtCode" type="text" id="txtCode' + commentid + '" size="18" class="ipt-txt fl" value="请输入验证码" onfocus="if(this.value==\'请输入验证码\')this.value=\'\';" onblur="if(this.value==\'\')this.value=\'请输入验证码\';" /><img  onclick="this.src=\'/validate.aspx?\'+Math.random()"  src="/validate.aspx" alt="将图中的文字填到左边输入框中" id="randomNoImg" class="fl"  /></span>';
			replydivHtml += '  <input type="submit" name="btnComment" value="发表评论"  onclick="PostQuickComment(' + commentid + ',' + newsid + ')" id="btnComment" class="button" />';
			replydivHtml += '  <span id="commentMessage' + commentid + '" style="color:red;"></span>';
            replydivHtml += '  </div>';

            ReplyDiv.innerHTML = replydivHtml;
     }
    }
    function CloseReplay(commentid) {
        var ReplyDiv = GetObj('Reply' + commentid);
        ReplyDiv.style.display = 'none';
}




 function PostQuickComment(commentid, newsid) {
     $.ajax({
         type: "POST",
         url: "/ithome/postComment.aspx",
         data: "newsid=" + newsid + "&commentNick=" + escape($("#commentNick" + commentid).val()) + "&commentContent=" + escape($("#commentContent" + commentid).val()) + "&parentCommentID=" + commentid + "&txtCode=" + $("#txtCode" + commentid).val() + "&type=comment",
         success: function (msg) {

             var messageObj = GetObj('commentMessage' + commentid);
             messageObj.innerHTML = '<span id="commentMessage" style="color:red">' + msg + '</span>';
             if (msg.indexOf("评论成功") >= 0) {
                 setCookie('username', $("#commentNick" + commentid).val(), 1);
                 var ReplyDiv = GetObj('Reply' + commentid);
                 ReplyDiv.style.display = 'none';
                 ReplyDiv.innerHTML = "";
                 self.parent.LoadData();

                

             }

         }
     });
 }

 function commentVote(commentid, typeid) {


     $.ajax({
         type: "POST",
         url: "/ithome/postComment.aspx",
         data: "commentid=" + commentid + "&type=replyVote&typeid=" + typeid,
         success: function (msg) {

             if (msg.indexOf("您") >= 0) {
                 alert(msg);
             }
             else {
                 if (typeid == 1) {
                     $("#agree" + commentid).text('支持(' + msg + ')');
                     $("#agree" + commentid).removeAttr("href");

                     $("#agree" + commentid).css({ "position": "relative" });
                     $("#agree" + commentid).append("<span class='flower'></span>");
                     $("#agree" + commentid).find(".flower").css({ "position": "absolute", "text-align": "center", "left": "6px", "top": "-10px", "display": "block", "width": "30px", "height": "30px", "background": "url(http://file.ithome.com/images/agree.gif) left center no-repeat", "opacity": "0" }).animate({ top: '-30px', opacity: '1' }, 300, function () { $(this).delay(300).animate({ top: '-35px', opacity: '0' }, 300) });
                     $("#agree" + commentid).find(".flower").removeClass();

                 }
                 else {
                     $("#against" + commentid).text('反对(' + msg + ')');
                     $("#against" + commentid).removeAttr("href");

                     $("#against" + commentid).css({ "position": "relative" });
                     $("#against" + commentid).append("<span class='shit'></span>");
                     $("#against" + commentid).find(".shit").css({ "position": "absolute", "text-align": "center", "left": "6px", "top": "-60px", "display": "block", "width": "30px", "height": "30px", "background": "url(http://file.ithome.com/images/against.gif) left center no-repeat", "opacity": "0" }).animate({ top: '-30px', opacity: '1' }, 300, function () { $(this).delay(300).animate({ top: '-5px', opacity: '0' }, 300) });
                     $("#against" + commentid).find(".shit").removeClass();

                 }
             }
         }
     });
 }

 function hotCommentVote(commentid, typeid) {
     $.ajax({
         type: "POST",
         url: "/ithome/postComment.aspx",
         data: "commentid=" + commentid + "&type=replyVote&typeid=" + typeid,
         success: function (msg) {
             if (msg.indexOf("您") >= 0) {
                 alert(msg);
             }
             else {
                 if (typeid == 1) {
                     $("#hotagree" + commentid).text('支持(' + msg + ')');
                     $("#hotagree" + commentid).removeAttr("href");
                     $("#hotagree" + commentid).css({ "position": "relative" });
                     $("#hotagree" + commentid).append("<span class='flower'></span>");
                     $("#hotagree" + commentid).find(".flower").css({ "position": "absolute", "text-align": "center", "left": "6px", "top": "-10px", "display": "block", "width": "30px", "height": "30px", "background": "url(http://file.ithome.com/images/agree.gif) left center no-repeat", "opacity": "0" }).animate({ top: '-30px', opacity: '1' }, 300, function() { $(this).delay(300).animate({ top: '-35px', opacity: '0' }, 300) });
                     $("#hotagree" + commentid).find(".flower").removeClass();

                 }
                 else {
                     $("#hotagainst" + commentid).text('反对(' + msg + ')');
                     $("#hotagainst" + commentid).removeAttr("href");

                     $("#hotagainst" + commentid).css({ "position": "relative" });
                     $("#hotagainst" + commentid).append("<span class='shit'></span>");
                     $("#hotagainst" + commentid).find(".shit").css({ "position": "absolute", "text-align": "center", "left": "6px", "top": "-60px", "display": "block", "width": "30px", "height": "30px", "background": "url(http://file.ithome.com/images/against.gif) left center no-repeat", "opacity": "0" }).animate({ top: '-30px', opacity: '1' }, 300, function() { $(this).delay(300).animate({ top: '-5px', opacity: '0' }, 300) });
                     $("#hotagainst" + commentid).find(".shit").removeClass();

                 }
             }
         }
     });
 }


function clearComment() {
    var comment = document.getElementById("commentContent");

    if (comment.value == "IT之家有您参与更精彩！") {
        comment.value = '';
        comment.onclick = function ()
        { }
    }
}

function showValidate() {
            var comment = GetObj('commentContent');
            if (comment.value == "IT之家有您参与更精彩！")
                comment.value = '';


  //  if ($("#divValidate").is(":hidden")) {
    //    randomNoImg.src = 'http://www.ithome.com/validate.aspx?' + Math.random();
  //      $("#divValidate").show();
  //  }

//    var Valdiv = GetObj('divValidate');
//    if (Valdiv.style.display == 'none') {

//        Valdiv.style.display = 'block';
//        randomNoImg.src = 'http://www.ithome.com/validate.aspx?' + Math.random();
//    }

}

function LoadData() {
    $("#LoadArticleReply").load("/ithome/GetAjaxData.aspx", { "newsID": $("#newsid").val(), "type": "comment" }, function () {

        $(document).ready(function () {
            location.replace("#LoadArticleReply");
        });
    }
     );
}

function pagecomment(page) {
    $("<div/>").load("/ithome/GetAjaxData.aspx", { "newsID": $("#newsid").val(), "type": "commentpage", "page": page }, function () {
    }).appendTo($("#ulcommentlist")).fadeIn('slow');
    var commentcount = parseInt(document.getElementById("commentcount").innerHTML)
    if (page * 50 > commentcount)
        $(".more_comm").hide();

}

$('.mobile a').attr('title','下载IT之家客户端，炫耀我的尾巴！');

/* Ctrl+Enter回复 */
$('textarea').attr('onkeydown','if(event.ctrlKey&&event.keyCode==13){document.getElementById("btnComment").click();return false}');