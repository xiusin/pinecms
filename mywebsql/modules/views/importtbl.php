<link href='/mywebsql/cache?css=theme,default,alerts,grid' rel="stylesheet" />

<div id="popup_wrapper">
	<div id="popup_contents">
		<%= MESSAGE %>
		<div class="padded"><%= T("Select data file to import") %></div>
		<div class="padded"><input type='file' name='impfile' size="40" /></div>
		<div class="input padded"><label for="table_name"><%= T("Select table for data import") %></label><select name="table" id="table_name"><%= TABLE_LIST %></select></div>
		<div class="padded"><input type='checkbox' name='header' id="field_header" value='yes' /><label class="right" for="field_header"><%= T("First line contains column names") %></label></div>
		<div class="padded"><input type='checkbox' name='ignore_errors' id="ignore_errors" value='yes' /><label class="right" for="ignore_errors"><%= T("Continue processing even if error occurs") %></label></div>
	</div>
	<div id="popup_footer">
		<div id="popup_buttons">
			<input type='button' id="btn_import" value='<%= T("Import") %>' />
		</div>
	</div>
</div>

<script type="text/javascript" language='javascript' src="/mywebsql/cache?script=common,jquery,ui,query,options,alerts"></script>
<script type="text/javascript" language='javascript'>
window.title = "<%= T("Import") %>";
$('#btn_import').button().click(function() {
	if (document.frmquery.impfile.value == '') {
		jAlert('<%= T("Select data file to import") %>');
		return false;
	}
	wrkfrmSubmit('importtbl', '', '', '')
});
if (<%= REFRESH %>)
	window.parent.objectsRefresh();
</script>
