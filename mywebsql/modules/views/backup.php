<link href='/mywebsql/cache?css=theme,default,alerts' rel="stylesheet" />

<style>
	div#db_objects { padding:3px;overflow:auto;height:310px;width:95%;border:3px double #efefef }
	div.objhead 	{ background-color:#ececec; padding: 5px; margin: 0 0 3px 0 }
	span.toggler 	{ display:inline-block; float:right; cursor: pointer; font-size:16px; margin: -5px 0 0 0 }
	div.obj 			{ padding:5px; margin:0 0 0 20px }
</style>

<div id="popup_wrapper">
	<div id="popup_contents">
		<%= MESSAGE %>
		<table border="0" cellpadding="5" cellspacing="8" style="width: 100%;height:90%">
		<tr>
		<td align="left" valign="top" width="45%">
			<div id="db_objects">
				<%= T("Either the database is empty, or there was an error retrieving list of database objects") %>.<br/>
				<%= T("Please try closing and re-opening this dialog again") %>.
			<div>
		</td>

		<td align="left" valign="top" width="55%">
		<fieldset>
			<legend><%= T("Backup type") %></legend>
			<table border="0" cellspacing="10" cellpadding="5" width="100%">
				<tr><td valign="top">
				<input type='radio' name='exptype' id='exptype1' value="struct" /><label class="right" for='exptype1'><%= T("Structure") %></label>
				</td></tr>
				<tr><td valign="top">
				<input type='radio' name='exptype' id='exptype2' value="data" /><label class="right" for='exptype2'><%= T("Table Data") %></label>
				</td></tr>
				<tr><td valign="top">
				<input type='radio' name='exptype' checked="1" id='exptype3' value="all" /><label class="right" for='exptype3'><%= T("Structure and Table Data") %></label>
				</td></tr>
			</table>
		</fieldset>

		<fieldset>
			<legend><%= T("Options") %></legend>
			<table border="0" cellspacing="10" cellpadding="5" width="100%">
				<tr><td valign="top">
				<input type='checkbox' name='auto_null' id='auto_null' /><label class="right" for='auto_null'><%= T("Set Auto increment field values to NULL") %></label>
				</td></tr>

				<tr><td valign="top">
				<input type='checkbox' name='dropcmd' id='dropcmd' /><label class="right" for='dropcmd'><%= T("Add DROP command before create statements") %></label>
				</td></tr>

				<tr><td valign="top">
				<input type='checkbox' name='emptycmd' id='emptycmd' /><label class="right" for='emptycmd'><%= T("Add TRUNCATE command before insert statements") %></label>
				</td></tr>

				<tr><td valign="top">
				<label><%= T("Backup filename") %>:</label><input type='text' name='filename' id='filename' value="<%= FILENAME %>" style="width:120px" />
				</td></tr>

			</table>
		</fieldset>

		<fieldset>
			<legend><%= T("Compression") %></legend>
			<table border="0" cellspacing="10" cellpadding="5" width="100%">
				<tr><td valign="top">
				<input type='radio' value="" name='compression' id='compress_none' checked="checked" /><label class="right" for='compress_none'><%= T("No Compression") %></label>
				</td>
<?php if (function_exists('bzopen')) { ?>
				<td valign="top">
				<input type='radio' value="bz" name='compression' id='compress_bzip' /><label class="right" for='compress_bzip'><%= T("BZip") %></label>
				</td>
<?php }
if (function_exists('gzopen')) { ?>
				<td valign="top">
				<input type='radio' value="gz" name='compression' id='compress_gzip' /><label class="right" for='compress_gzip'><%= T("GZip") %></label>
				</td>
<?php } ?>
				</tr>
			</table>
		</fieldset>

		</td>
		</tr>
		</table>
	</div>
	<div id="popup_footer">
		<div id="popup_buttons">
			<input type='button' id="btn_export" value='<%= T("Export") %>' />
		</div>
	</div>
</div>

<script type="text/javascript" language='javascript' src="/mywebsql/cache?script=common,jquery,ui,query,options,alerts"></script>
<script type="text/javascript" language="javascript">
window.title = "<%= T("Backup Database") %>";
var exportType = 'backup';
<?php
	foreach( $data as $name => $list ) {
		echo "var {$name} = " . json_encode( $list ) .";\n";
	}
?>

$(function() {
	$('#popup_overlay').remove();  // we do not want to show the popup overlay when form is submitted
	$('#btn_export').button().click(function() { exportBackup() });

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
