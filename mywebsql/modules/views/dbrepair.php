<link href='/mywebsql/cache?css=theme,default,alerts,results' rel="stylesheet" />

<style>
	div#db_objects { margin-top:5px;padding:3px;overflow:auto;height:300px;width:220px;border:3px double #efefef }
	div.objhead 	{ background-color:#ececec; padding: 5px; margin: 0 0 3px 0 }
	span.toggler 	{ display:inline-block; float:right; cursor: pointer; font-size:16px; margin: -5px 0 0 0 }
	div.obj 			{ padding:5px; margin:0 0 0 20px }
</style>

<div id="popup_wrapper">
	<div id="popup_contents">
		<table border="0" cellpadding="5" cellspacing="4" style="width: 100%;height:100%">
		<tr>
		<td align="left" valign="top" width="45%">
			<%= T("Select tables to be analyzed/repaired") %><br />
			<div id="db_objects">
				<%= T("Either the database is empty, or there was an error retrieving list of database objects") %>.<br/>
				<%= T("Please try closing and re-opening this dialog again") %>.
			<div>
		</td>

		<td align="left" valign="top" width="55%">
		<fieldset>
			<legend><%= T("Operation to perform") %></legend>
			<table border="0" cellspacing="10" cellpadding="5" width="100%">
				<tr><td valign="top">
				<input type='radio' name='optype' id='optype1' value="analyze" checked="1" /><label class="right" for='optype1'><%= T("Analyze") %></label>
				</td></tr>
				<tr><td valign="top">
				<input type='radio' name='optype' id='optype2' value="check" /><label class="right" for='optype2'><%= T("Check") %></label>
				</td></tr>
				<tr><td valign="top">
				<input type='radio' name='optype' id='optype3' value="optimize" /><label class="right" for='optype3'><%= T("Optimize") %></label>
				</td></tr>
				<tr><td valign="top">
				<input type='radio' name='optype' id='optype4' value="repair" /><label class="right" for='optype4'><%= T("Repair") %></label>
				</td></tr>
			</table>
		</fieldset>

		<fieldset>
			<legend><%= T("Options") %></legend>
			<table border="0" cellspacing="10" cellpadding="5" width="100%">
				<tr><td valign="top">
				<input type='checkbox' name='skiplog' id='skiplog' /><label class="right" for='skiplog'><%= T("Skip Binary logging") %></label>
				</td></tr>

				<tr id="check_options" style="display:none"><td valign="top">
					<table border="0" cellspacing="5" cellpadding="3" width="100%">
						<tr><td><input type='radio' name='checktype' id='checktype1' value="default" checked="1" /><label class="right" for='checktype1'><%= T("Default") %></label></td></tr>
						<tr><td><input type='radio' name='checktype' id='checktype2' value="quick" /><label class="right" for='checktype2'><%= T("Quick") %></label></td></tr>
						<tr><td><input type='radio' name='checktype' id='checktype3' value="fast" /><label class="right" for='checktype3'><%= T("Fast") %></label></td></tr>
						<tr><td><input type='radio' name='checktype' id='checktype4' value="medium" /><label class="right" for='checktype4'><%= T("Medium") %></label></td></tr>
						<tr><td><input type='radio' name='checktype' id='checktype5' value="extended" /><label class="right" for='checktype5'><%= T("Extended") %></label></td></tr>
						<tr><td><input type='radio' name='checktype' id='checktype6' value="changed" /><label class="right" for='checktype6'><%= T("Changed") %></label></td></tr>
					</table>
				</td></tr>

				<tr id="repair_options" style="display:none"><td valign="top">
					<table border="0" cellspacing="5" cellpadding="3" width="100%">
						<tr><td><input type='checkbox' name='repairtype' id='repairtype1' value="quick" /><label class="right" for='repairtype1'><%= T("Quick") %></label></td></tr>
						<tr><td><input type='checkbox' name='repairtype' id='repairtype2' value="extended" /><label class="right" for='repairtype2'><%= T("Extended") %></label></td></tr>
						<tr><td><input type='checkbox' name='repairtype' id='repairtype3' value="usefrm" /><label class="right" for='repairtype3'><%= T("Use Frm files (MyISAM tables)") %></label></td></tr>
					</table>
				</td></tr>

			</table>
		</fieldset>
		</td>
		</tr>
		</table>
	</div>
	<div id="popup_footer">
		<div id="popup_buttons">
			<input type='button' id="btn_repair" value='<%= T("Submit") %>' />
		</div>
	</div>
</div>

<script type="text/javascript" language='javascript' src="/mywebsql/cache?script=common,jquery,ui,query,options,alerts"></script>
<script type="text/javascript" language="javascript">
window.title = "<%= T("Repair Tables") %>";
var repairType = 'analyze';
<?php
	echo "var tables = " . json_encode( $data ) .";\n";
?>

$(function() {
	$('#btn_repair').button().click(function() { repairTables() });

<?php
	if ( count($data) > 0 ) {
?>
		$('#db_objects').html('');
<?php
		echo "uiShowObjectList(tables, 'tables', '" . __('Tables') . "');\n";
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

	$('#optype1').click(function() { $('#check_options').hide(); $('#repair_options').hide(); });
	$('#optype2').click(function() { $('#check_options').show(); $('#repair_options').hide(); });
	$('#optype3').click(function() { $('#check_options').hide(); $('#repair_options').hide(); });
	$('#optype4').click(function() { $('#check_options').hide(); $('#repair_options').show(); });

});
</script>
<?php
	echo getGeneratedJS();
?>
