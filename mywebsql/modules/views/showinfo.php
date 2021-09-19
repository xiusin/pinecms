<div id='results'>
	<div class="message ui-state-default"><%= T("Create command for <%= TYPE %> <%= NAME %>") %></div>
	<div class="sql-text ui-state-default">
		<%= COMMAND %>
	</div>
</div>

<script type="text/javascript" language="javascript">
parent.transferInfoMessage();
//parent.addCmdHistory("<%= SQL %>");
</script>
