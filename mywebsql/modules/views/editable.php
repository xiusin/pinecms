<link href='/mywebsql/cache?css=theme,default,grid,alerts,editor' rel="stylesheet" />

<div id="popup_wrapper">
	<div id="popup_contents">
		<div id="grid-messages"><%= MESSAGE %></div>

		<div id="grid-tabs">
			<ul>
				<li><a href="#tab-fields"><%= T("Basic Information") %></a></li>
				<li><a href="#tab-props"><%= T("Table Properties") %></a></li>
				<li><a href="#tab-messages"><%= T("Messages") %></a></li>
			</ul>
			<div class="ui-corner-bottom">
				<div id="tab-fields">
					<div class="input">
						<span><%= T("Table Name") %>:</span><span><input type="text" size="20" name="table-name" id="table-name" value="<%= TABLE_NAME %>" /><span>
					</div>
					<table border="0" cellspacing="1" cellpadding="2" id="table_grid"><tbody>
						<tr id='fhead'>
							<th style="width:120px"><%= T("Field Name") %></th>
							<th style="width:75px"><%= T("Data Type") %></th>
							<th style="width:65px"><%= T("Length") %></th>
							<th style="width:90px"><%= T("Default value") %></th>
							<th style="width:65px"><%= T("Unsigned") %></th>
							<th style="width:65px"><%= T("Zero Fill") %></th>
							<th style="width:85px"><%= T("Primary Key") %></th>
							<th style="width:100px"><%= T("Auto Increment") %></th>
							<th style="width:70px"><%= T("Not NULL") %></th>
						</tr>
					</tbody></table>
				</div>
				<div id="tab-props">
					<div class="input"><span><%= T("Table Engine (type)") %>:</span><span><select name="enginetype" id="enginetype"><%= ENGINE %></select><span></div>
					<div class="input float"><span><%= T("Character Set") %>:</span><span><select name="charset" id="charset"><%= CHARSET %></select><span></div>
					<div class="input"><span><%= T("Collation") %>:</span><span><select name="collation" id="collation"><%= COLLATION %></select><span></div>
					<div class="input"><span><%= T("Comment") %>:</span><span><input type="text" size="40" name="comment" id="comment" value="<%= COMMENT %>" style="width:488px" /><span></div>
				</div>
				<div id="tab-messages">
					<%= T("Waiting for table information to be submitted") %>
				</div>
			</div>
		</div>
	</div>

	<div id="popup_footer">
		<div id="popup_buttons">
			<input type='button' id='btn_add' value='<%= T("Add field") %>' />
			<input type='button' id='btn_del' value='<%= T("Delete selected field") %>' />
			<input type='button' id='btn_clear' value='<%= T("Clear Table Information") %>' />
			<input type='button' id='btn_submit' value='<%= T("Submit") %>' tabindex="1" />
		</div>
	</div>

</div>

<div id="dialog-list" title="<%= T("List of values") %>">
	<div class="padded">
		<div>
			<select size="8" name="list-items" id="list-items"></select>
		</div>
		<div>
			<input type="text" name="item" id="item" class="text ui-widget-content" />
		</div>
	</div>
</div>

<script type="text/javascript" language='javascript' src="/mywebsql/cache?script=common,jquery,ui,editable,position,query,cookies,settings,alerts,hotkeys"></script>
<script type="text/javascript" language="javascript">
window.title = (<%= ALTER_TABLE %> ? "<%= T("Edit table - ") %><%= TABLE_NAME %>" : "<%= T("Create Table") %>");
var rowInfo = <%= ROWINFO %>;

$(function() {
	setupEditable(<%= ALTER_TABLE %>);
});
</script>
