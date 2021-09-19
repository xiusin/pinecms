<link href='/mywebsql/cache?css=theme,default,alerts' rel="stylesheet" />

<style>
	div#db_objects { margin-top:5px;padding:3px;overflow:auto;height:300px;width:95%;border:3px double #efefef }
	div.objhead 	{ background-color:#ececec; padding: 5px; margin: 0 0 3px 0 }
	span.toggler 	{ display:inline-block; float:right; cursor: pointer; font-size:16px; margin: -5px 0 0 0 }
	div.obj 			{ padding:5px; margin:0 0 0 20px }
</style>

<div id="popup_wrapper">
	<div id="popup_contents">
		<?php if( isset( $data['stats'] ) ) { ?>
			<div class="message ui-state-default">
				<?php if( isset($data['stats']['drop']) ) {
					$txt = '<p><span class="ui-icon ui-icon-check"></span>' . str_replace('<%= NUM %>', $data['stats']['drop']['success'], __('<%= NUM %> queries successfully executed')) . '</p>';
					if ( $data['stats']['drop']['errors'] > 0 )
						$txt .= '<p><span class="ui-icon ui-icon-close"></span>' . str_replace('<%= NUM %>', $data['stats']['drop']['errors'], __('<%= NUM %> queries failed to execute')) . '</p>';
					echo $txt;
				?>
				<?php } ?>
			</div>
		<?php } else { ?>
			<div class="message ui-state-error"><%= T("WARNING') . ': ' . __('The following operation is irreversible') . '. ' . __('Potential data loss might occur") %></div>
		<?php }?>

		<table border="0" cellpadding="5" cellspacing="8" style="width: 100%;height:100%">
		<tr>
		<td align="left" valign="top" width="45%">
			<%= T("Select databases to operate upon") %><br />
			<div id="db_objects">
				<%= T("Either the database is empty, or there was an error retrieving list of database objects") %>.<br/>
				<%= T("Please try closing and re-opening this dialog again") %>.
			<div>
		</td>

		<td align="left" valign="top" width="55%">
		<fieldset>
			<legend><%= T("Operations to perform") %></legend>
			<table border="0" cellspacing="10" cellpadding="5" width="100%" id="dbops">
				<tr><td valign="top">
				<input type='checkbox' name='dropcmd' id='dropcmd' /><label class="right" for='dropcmd'><%= T("DROP selected databases") %></label>
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
window.title = "<%= T("Database Manager") %>";
<?php
	echo "var databases = " . json_encode( $data['objects'] ) .";\n";
?>

$(function() {
	$('#btn_submit').button().click(function() {
		if ( $("#db_objects").find("input[type=checkbox]").filter(":checked").length == 0 ) {
		 	jAlert(__('Select databases to operate upon'));
		} else if ($("#dropcmd").prop("checked")) {
			jConfirm(__('Are you sure you want to DROP selected databases?'), __('Confirm Action'), function(result) {
				if (result)
					wrkfrmSubmit('databases', 'batch', '', '');
			}, '');
		} else if ($("#dbops input[type=checkbox]").filter(":checked").length == 0) {
			jAlert(__('Please select one or more operations to perform'));
		}
	});


	$("#dbops input[type=checkbox]").click(function() {
		var id = $(this).attr("id");
		if (id == "dropcmd") {
			var disabled = $(this).prop("checked");
			if (disabled) {
				$("#dbops input[type=checkbox]:gt(0)").prop("checked", false).attr("disabled", "disabled");
			} else {
				$("#dbops input[type=checkbox]:gt(0)").removeAttr("disabled");
			}
		}
	});

	if (databases.length == 0)
		return;

<?php
	if ( count($data['objects']) > 0 ) {
?>
		$('#db_objects').html('');
<?php
		echo "uiShowObjectList(databases, 'databases', '" . __( 'Databases' ) . "', true);\n";
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
