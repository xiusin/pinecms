<link href="/mywebsql/cache?css=theme,default,help" rel="stylesheet" />

<div id="popup_wrapper">

	<div class="docinfo">
		<?php echo str_replace('<%= LINK %>', '', __('To see most up-to-date help contents, please visit <%= LINK %>')); ?>
		<a target="_blank" href="<%= PROJECT_SITEURL %>/docs">MyWebSQL Online Documentation</a>
	</div>

	<ul class="links">
	<?php
	foreach($data['pages'] as $x=>$y) {
			if ($data['page'] == $x)
				echo "<li class=\"current\"><img border=\"0\" align=\"absmiddle\" src='img/help/t_$x".".gif' alt=\"\" />$y</li>";
			else
				echo "<li><a href=\"#$x\"><img border=\"0\" align=\"absmiddle\" src='img/help/t_$x".".gif' alt=\"\" />$y</a></li>";
		}
	?>
	</ul>
	<div class="content">
	<%= CONTENT %>
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
