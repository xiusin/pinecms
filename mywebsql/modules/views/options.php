<link href="/mywebsql/cache?css=theme,default,help" rel="stylesheet" />
<script language="javascript" src="/mywebsql/cache?script=jquery,ui,cookies,options,settings" type="text/javascript"></script>

<div id="popup_wrapper">

	<ul class="links">
	<%= lis %>
	</ul>
	<div class="content">
	<%= CONTENT %>
	</div>

</div>

<script type="text/javascript" language="javascript">
	window.title = "<%= T("Options") %>";
	var COOKIE_LIFETIME = 1440;
	$(function() {
		$('ul.links a').click(function() {
			page = $(this).attr('href').replace('#', '');
			navigatePage('options', page);
		});

		$("#save").button().click(function() { $(this).button("disable"); });
		$("#reset").button().click( optionsReset );

		// enable save button when any setting is changed
		$("input[type=radio], input[type=checkbox]").click(function() {
			$("#save").button("enable");
		});
		$("input[type=text]").change(function() {
			$("#save").button("enable");
		});
	});
</script>
