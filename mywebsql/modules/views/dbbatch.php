<link href='/mywebsql/cache?css=theme,default,alerts' rel="stylesheet" />

<style>
	div#db_objects { margin-top:5px;padding:3px;overflow:auto;height:300px;width:95%;border:3px double #efefef }
	div.objhead 	{ background-color:#ececec; padding: 5px; margin: 0 0 3px 0 }
	span.toggler 	{ display:inline-block; float:right; cursor: pointer; font-size:16px; margin: -5px 0 0 0 }
	div.obj 			{ padding:5px; margin:0 0 0 20px }
</style>

<div id="popup_wrapper">
	<div id="popup_contents">
		<table border="0" cellpadding="5" cellspacing="8" style="width: 100%;height:100%">
		<tr>
		<td align="left" valign="top" width="45%">
			<%= T("Select objects to operate upon") %><br />
			<div id="db_objects">
				<%= T("Either the database is empty, or there was an error retrieving list of database objects") %>.<br/>
				<%= T("Please try closing and re-opening this dialog again") %>.
			<div>
		</td>

		<td align="left" valign="top" width="55%">
		<fieldset>
			<legend><%= T("Operations to perform") %></legend>
			<table border="0" cellspacing="10" cellpadding="5" width="100%">
				<tr><td valign="top">
					<label for='old_prefix'><%= T("Delete prefix string from name") %></label><input type='text' name='old_prefix' id='old_prefix' maxlength="10" style="width:70px" />
				</td></tr>

				<tr><td valign="top">
					<label for='new_prefix'><%= T("Add prefix string to name") %></label><input type='text' name='new_prefix' id='new_prefix' maxlength="10" style="width:70px" />
				</td></tr>

				<tr><td valign="top">
				<input type='checkbox' name='dropcmd' id='dropcmd' /><label class="right" for='dropcmd'><%= T("DROP selected database objects") %></label>
				</td></tr>
			</table>
		</fieldset>

		<fieldset>
			<legend><%= T("Generate SQL") %></legend>
			<table border="0" cellspacing="10" cellpadding="5" width="100%">
				<tr><td valign="top">
					<label for='command'><%= T("Command text") %></label><input type='text' name='command' id='command' maxlength="100" style="width:180px" />
				</td></tr>
			</table>
		</fieldset>

		<fieldset>
			<legend><%= T("Options") %></legend>
			<table border="0" cellspacing="10" cellpadding="5" width="100%">
				<tr><td valign="top">
					<input type='checkbox' name='skip_fkey' id='skip_fkey' /><label class="right" for='skip_fkey'><%= T("Skip Foreign Key checks") %></label>
				</td></tr>
			</table>
		</fieldset>

		</td>
		</tr>
		</table>
	</div>
	<div id="popup_footer">
		<div id="popup_buttons">
			<input type='button' id="btn_submit" value='<%= T("Submit") %>' />
		</div>
	</div>
</div>

<script type="text/javascript" language='javascript' src="/mywebsql/cache?script=common,jquery,ui,query,options,alerts"></script>
<script type="text/javascript" language="javascript">
window.title = "<%= T("Batch operations") %>";
<?php
	foreach( $data as $name => $list ) {
		echo "var {$name} = " . json_encode( $list ) .";\n";
	}
?>

$(function() {
	$('#btn_submit').button().click(function() {
		if ( $("#db_objects").find("input[type=checkbox]").filter(":checked").length == 0 ) {
		 	jAlert(__('Select objects to operate upon'));
		} else if ($("#dropcmd").prop("checked")) {
			// if drop command is selected, confirm user for the operation
			jConfirm(__('Are you sure you want to DROP selected objects?'), __('Confirm Action'), function(result) {
				if (result)
					wrkfrmSubmit('dbbatch', 'batch', '', '');
			}, '');
		 } else if ( $("#old_prefix").val() == '' && $("#new_prefix").val() == '' && $("#command").val() == '' ) {
			jAlert(__('Please select one or more operations to perform'));
		 } else {
			wrkfrmSubmit('dbbatch', 'batch', '', '');
		 }
	});

	$("#dropcmd").click(function() {
		var on = $(this).prop("checked");
		$("#new_prefix").add("#old_prefix").attr("disabled", on);
	});

	if (tables.length == 0 && views.length == 0 && procs.length == 0 && funcs.length == 0 && triggers.length == 0)
		return;

<?php
	if ( count($data) > 0 ) {
?>
		$('#db_objects').html('');
<?php
		foreach( $data as $name => $list ) {
			echo "uiShowObjectList({$name}, '{$name}', '" . __( ucfirst($name) ) . "');\n";
		}
	}
?>
	$('.selectall').click(function(e) {
		chk = $(this).attr('checked');
		chk ? $(this).parent().next().find('input').attr('checked', "checked") : $(this).parent().next().find('input').removeAttr('checked');
	});

	$('#db_objects .toggler').click(function() {
		$(this).parent().next().toggle();
		if ($(this).hasClass('c')) {
			$(this).removeClass('c').html('&#x25B4;');
		} else {
			$(this).addClass('c').html('&#x25BE;');
		}
		return false;
	});
});
</script>
<?php
	echo getGeneratedJS();
?>
