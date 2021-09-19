<link href='/mywebsql/cache?css=theme,default,alerts,grid' rel="stylesheet" />

<style>
	#db_names option		{ padding: 3px 6px; }
</style>

<div id="popup_wrapper" style="display:none">
	<div id="popup_contents">
		<div class="input" id="button-list">
			<span><%= T("Select a User") %>:</span><span><select name="userlist" id="userlist"><%= USERS %></select></span>
			<span><input type='button' id='btn_del' value='<%= T("Delete selected User") %>' /></span>
			<span><input type='button' id='btn_add' value='<%= T("Add User") %>' /></span>
		</div>

		<div id="grid-tabs">
			<div id="grid-messages"><%= MESSAGE %></div>
			<ul>
				<li><a href="#tab-general"><%= T("User Information") %></a></li>
				<li><a href="#tab-global"><%= T("Global Privileges") %></a></li>
				<li><a href="#tab-db"><%= T("Database Privileges") %></a></li>
			</ul>
			<div class="ui-corner-bottom">
				<div id="tab-general">
					<div class="input"><span><%= T("User Name") %>:</span><span><input type="text" size="30" name="username" id="username" /><span></div>
					<div class="input"><span><%= T("Host") %>:</span><span><input type="text" size="30" name="hostname" id="hostname" /><span></div>
					<div class="input"><span><%= T("Password") %>:</span><span><input autocomplete="off" type="password" size="30" name="userpass" id="userpass" /><span></div>
					<div class="input"><span><%= T("Confirm Password") %>:</span><span><input autocomplete="off" type="password" size="30" name="userpass2" id="userpass2" /><span></div>
					<div class="input"><span><input type="checkbox" name="nopass" id="nopass" /><label class="right" for="nopass"><%= T("Remove Password") %></label></span></div>
				</div>
				<div id="tab-global">
					&nbsp;
				</div>
				<div id="tab-db">
					&nbsp;
				</div>
			</div>
		</div>
	</div>

	<div id="popup_footer">
		<div id="popup_buttons">
			<div id="checkboxes" style="float:left"><input type="checkbox" name="selectall" id="selectall"><label class="right" for="selectall"><%= T("Select All/None") %></label></div>
			<div style="float:right"><input type='button' id='btn_cancel' value='<%= T("Cancel") %>' /></div>
			<div style="float:right"><input type='button' id='btn_add2' value='<%= T("Add User") %>' /></div>
			<div style="float:right"><input type='button' id='btn_submit' value='<%= T("Update User") %>' tabindex="1" /></div>
		</div>
	</div>

</div>

<script type="text/javascript" language='javascript' src="/mywebsql/cache?script=common,jquery,ui,editable,position,query,cookies,settings,alerts,users"></script>
<script type="text/javascript" language="javascript">
window.title = "<?php echo __('User Manager')?>";

var USER_INFO = <%= USER_INFO %>;
var DATABASES = <%= DATABASES %>;
var PRIVILEGES = <%= PRIVILEGES %>;
var DB_PRIVILEGES = <%= DB_PRIVILEGES %>;
var PRIVILEGE_NAMES = <%= PRIVILEGE_NAMES %>;
var DB_PRIVILEGE_NAMES = <%= DB_PRIVILEGE_NAMES %>;

$(function() {
	$('#grid-tabs').tabs({
		select: function(event, ui) {
			if (ui.index == 0)
				$('#checkboxes').hide();
			else if ( (ui.index == 2 && !$('#db_names').val()) )
				$('#checkboxes').hide();
			else {
				$('#checkboxes').show();
				cls = (ui.index==1) ? '#tab-global .prv' : '#tab-db .dbprv';
				$('#selectall').attr('checked', $(cls).not(':checked').length == 0);
			}
		}
	});
	$('#userlist').change(selectUser);
	$('#btn_add').button().click(addUser);
	$('#btn_del').button().click(deleteUser);
	$('#btn_cancel').button().click(cancelOperation).hide();
	$('#btn_add2').button().click(addNewUser).hide();
	$('#btn_submit').button().click(updateUser);
	$('#selectall').click(selectAll);
	$('#nopass').click(function() {
		checked = $(this).attr('checked');
		if (checked)
			$('#userpass,#userpass2').attr('disabled', 'disabled');
		else
			$('#userpass,#userpass2').removeAttr('disabled');
	});
	loadUserData();
});
</script>
