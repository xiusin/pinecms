<fieldset style="margin: 10px">
	<legend><%= T("Database Objects") %></legend>
	<table border="0" cellspacing="10" cellpadding="5" width="100%">
		<tr><td valign="top">
		<input type='button' id="reset" value="<%= T("Reset all confirmation dialogs") %>" />
		</td></tr>
	</table>
</fieldset>
<script type="text/javascript" language="javascript">
   $(function () {
   	$("#reset").click(function () {
   		$(this).button("disable");
   	});
   });
</script>
