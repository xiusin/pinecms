<fieldset style="margin: 10px">
	<legend><%= T("Results") %></legend>
	<table border="0" cellspacing="10" cellpadding="5" width="100%">
		<tr><td valign="top">
		<label for='res-max-count'><%= T("Maximum records to display in result set") %></label>
			<input type='text' name='res-max-count' id='res-max-count' style="width:35px" value="<?php echo Options::get('res-max-count', MAX_RECORD_TO_DISPLAY); ?>" />
		</td></tr>
		<tr><td valign="top">
		<input type='submit' id="save" value="<%= T("Save") %>" />
		</td></tr>
	</table>
</fieldset>
<script type="text/javascript" language="javascript">
   $(function () {
   	$("#save").click(function () {
   		optionsSet( 'res-max-count', $('#res-max-count').val() );
   	});
   });
</script>
