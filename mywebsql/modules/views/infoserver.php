<div id='results'>
	<div class='heading1'><%= T("Server information") %></div>
	<div class="hinfo">
		<div class="label"><%= T("Server") %></div><div class="info"><%= SERVER_NAME %></div>
		<div class="label"><%= T("Version") %></div><div class="info"><%= SERVER_VERSION %></div>
		<div class="label"><%= T("Version comment") %></div><div class="info"><%= SERVER_COMMENT %></div>
	</div>
	<div class='heading1'><%= T("Character sets") %></div>
	<div class="hinfo">
		<div class="label"><%= T("Server character set") %></div><div class="info"><%= SERVER_CHARSET %></div>
		<div class="label"><%= T("Client character set") %></div><div class="info"><%= CLIENT_CHARSET %></div>
		<div class="label"><%= T("Database character set") %></div><div class="info"><%= DATABASE_CHARSET %></div>
		<div class="label"><%= T("Results character set") %></div><div class="info"><%= RESULT_CHARSET %></div>
	</div>
</div>

<script type="text/javascript" language="javascript">
parent.transferInfoMessage();
<%= JS %>
</script>
