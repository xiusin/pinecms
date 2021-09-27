<div id='results'>
	<div class="message ui-state-default"><%= MESSAGE %></div>
	<div class="sql-text ui-state-default">
		<%= COMMAND %>
	</div>
</div>

<script type="text/javascript" language="javascript">
parent.transferInfoMessage();
parent.addCmdHistory("<%= SQL %>");
</script>
