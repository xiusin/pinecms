<fieldset style="margin: 10px">
	<legend><%= T("Record Editing") %></legend>
	<table border="0" cellspacing="10" cellpadding="5" width="100%">
		<tr><td valign="top">
		<input type='checkbox' name='ui-edit-popup' id='ui-edit-popup' <?php if(Options::get('ui-edit-popup', false)) echo 'checked="checked" '; ?>/><label class="right" for='ui-edit-popup'><%= T("Show popup dialog for editing large text data") %></label>
		</td></tr>
		<tr><td valign="top">
		<input type='submit' id="save" value="<%= T("Save") %>" />
		</td></tr>
	</table>
</fieldset>
<script type="text/javascript" language="javascript">
   $(function () {
   	$("#save").click(function () {
   		optionsSet('ui-edit-popup', $('#ui-edit-popup').prop("checked") ? 1 : 0);
   	});
   });
</script>
