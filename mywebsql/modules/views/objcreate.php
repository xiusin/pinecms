<link href='/mywebsql/cache?css=theme,default' rel="stylesheet" />
	<input type='hidden' name='q' value="wrkfrm" />
	<input type='hidden' name='type' value="objcreate" />
	<input type='hidden' name='id' value="<%= ID %>" />

<div id="popup_wrapper">
	<div id="popup_contents">
		<div class="<%= MESSAGE_TYPE %>"><%= MESSAGE %></div>

		<div class="code-editor">
			<textarea cols="86" rows="16" name="objinfo" id="objinfo" class="text-editor"><%= OBJINFO %></textarea>
		</div>
	</div>

	<div id="popup_footer">
		<div id="popup_buttons">
			<input type="button" id="btn-submit" value="<%= T("Submit") %>" />
		</div>
	</div>
</div>

<script type="text/javascript" language='javascript' src="/mywebsql/cache?script=common,jquery,ui"></script>
<%= EDITOR_LINK %>
<script type="text/javascript" language="javascript">
window.title = "<%= T("Create new database object") %>";
var code_editor = null;
$(function() {
	document.frmquery.objinfo.focus();
	code_editor = CodeMirror.fromTextArea('objinfo', { <%= EDITOR_OPTIONS %>,
		width: '100%', height: '320px', tabMode : 'default',
		stylesheet: '/mywebsql/cache?css=editor', onLoad : function() { }
	});
	$('#btn-submit').button().click(submit_form);
	if ("<%= MESSAGE_TYPE %>" == "success")
		parent.objectsRefresh();
});

function submit_form() {
	if(code_editor)
		document.frmquery.objinfo.value = code_editor.getCode();
	document.frmquery.submit();
}

if (<%= REFRESH %>)
	window.parent.objectsRefresh();
</script>
</script>


