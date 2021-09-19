<link href='/mywebsql/cache?css=theme,default,alerts,grid' rel="stylesheet" />
<div id="popup_wrapper">
	<div id="popup_contents">
		<%= MESSAGE %>
		<div id="grid-tabs" class="padded">
			<div class="input"><span><%= T("Table Engine (type)") %>:</span><span><select name="enginetype" id="enginetype"><%= ENGINE %></select><span></div>
		</div>
	</div>
	<div id="popup_footer">
		<div id="popup_buttons">
			<input type='button' id="btn_alter" value='<%= T("Submit") %>' />
		</div>
	</div>
</div>

<script type="text/javascript" language='javascript' src="/mywebsql/cache?script=common,jquery,ui,query,options"></script>
<script type="text/javascript" language='javascript'>
window.title = "<%= T("Change Table Type") %>";
$('#btn_alter').button().click(function() {
	wrkfrmSubmit('enginetype', 'alter', '<%= TABLE_NAME %>', '')
});
</script>
