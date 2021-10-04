<link href="/mywebsql/cache?css=theme,default,help" rel="stylesheet" />

<div id="popup_wrapper">

	<div class="docinfo">
	    <%= MSG %>
		<a target="_blank" href="//gowebdb.xiusin.cn/docs/">MyWebSQL Online Documentation</a>
	</div>

	<ul class="links">
	<%= for (x, y) in indexs { %>
	    <%= if (y == page) { %>
	         <li class="current"><img border="0" align="absmiddle" src='/mywebsql/img/help/t_<%= y %>.gif' alt="" /><%= pages[y] %></li>
	    <% } else { %>
	    	 <li><a href="#<%= y %>"><img border="0" align="absmiddle" src='/mywebsql/img/help/t_<%= y %>.gif' alt="" /><%= pages[y] %></a></li>
	    <% } %>
	<% } %>
	</ul>
	<div class="content">
	<%= contents %>
	</div>

</div>

<script language="javascript" src="/mywebsql/cache?script=jquery,options" type="text/javascript"></script>
<script type="text/javascript" language="javascript">
	window.title = "<%= T("Help") %>";
	$(function() {
		$('ul.links a').click(function() {
			page = $(this).attr('href').replace('#', '');
			navigatePage('help', page);
		});
	});
</script>
