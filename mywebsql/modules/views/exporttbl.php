<link href='/mywebsql/cache?css=theme,default,alerts' rel="stylesheet" />

<div id="popup_wrapper">
	<div id="popup_contents">
		<table border="0" cellpadding="5" cellspacing="4" style="width: 100%">
		<tr>
			<td align="left" valign="top" width="45%">

			<fieldset>
				<legend><%= T("Export As") %></legend>
				<table border="0" cellspacing="10" cellpadding="5" width="100%">
					<tbody id="exp-types">
					<tr><td valign="top">
					<input type='radio' name='exptype' id='insert' value="insert" checked="checked" /><label class="right" for='insert'><%= T("Insert Statements") %></label>
					</td></tr>

					<tr><td valign="top">
					<input type='radio' name='exptype' id='xml' value="xml" /><label class="right" for='xml'><%= T("XML") %></label>
					</td></tr>

					<tr><td valign="top">
					<input type='radio' name='exptype' id='xhtml' value="xhtml" /><label class="right" for='xhtml'><%= T("XHTML") %></label>
					</td></tr>

					<tr><td valign="top">
					<input type='radio' name='exptype' id='csv' value="csv" /><label class="right" for='csv'><%= T("Comma Separated (CSV for Excel)") %></label>
					</td></tr>

					<tr><td valign="top">
					<input type='radio' name='exptype' id='yaml' value="yaml" /><label class="right" for='yaml'><%= T("YAML") %></label>
					</td></tr>

					<tr><td valign="top">
					<input type='radio' name='exptype' id='text' value="text" /><label class="right" for='text'><%= T("Plain Text (One record per line)") %></label>
					</td></tr>
					</tbody>
				</table>
			</fieldset>
			</td>

			<td align="left" valign="top" width="55%">
				<fieldset>
				<legend><%= T("Export Options") %></legend>
					<table border="0" cellspacing="10" cellpadding="5" width="100%">
					<tbody id="exp-options">
						<tr class="insert"><td valign="top">
							<input type='checkbox' name='fieldnames' id='fieldnames' checked="checked" /><label class="right" for='fieldnames'><%= T("Include field names in query") %></label>
						</td></tr>
						<tr class="insert"><td valign="top">
						<input type='checkbox' name='bulkinsert' id='bulkinsert' /><label class="right" for='bulkinsert'><%= T("Generate Bulk insert statements") %></label>
						</td></tr>
						<tr class="insert"><td valign="top">
						<input disabled="disabled" type='checkbox' name='bulklimit' id='bulklimit' /><label class="right" for='bulklimit'><%= T("Maximum size of SQL statement") %></label>
						&nbsp;<input disabled="disabled" type="text" name="bulksize" id="bulksize" style="height:12px;vertical-align:bottom;width:30px" />&nbsp;KB
						</td></tr>
						<tr class="csv ui-helper-hidden"><td valign="top">
							<input type='checkbox' name='fieldheader' id='fieldheader' checked="checked" /><label class="right" for='fieldheader'><%= T("Field names in first row") %></label>
						</td></tr>
						<tr class="text ui-helper-hidden"><td valign="top">
						<label><%= T("Fields separated by:") %></label><input type='text' size="4" name='separator' id='separator' value="\t" class="text" style="width:30px" />
						</td></tr>
						<tr class="xml xhtml yaml ui-helper-hidden"><td valign="top">
						<label><%= T("There are no options to configure for selected export type") %></label>
						</td></tr>
					<tbody>
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
window.title = "<%= T("Export Table") %>";
var exportType = 'exporttbl';
$(function() {
	$('#popup_overlay').remove();  // we do not want to show the popup overlay when form is submitted
	$('#btn_export').button().click(function() { exportData('tbl', "<%= TABLENAME %>"); });

	$("#exp-types input").click(function() {
		cls = $(this).val();
		tr = $("#exp-options").find("tr");
		tr.addClass("ui-helper-hidden").filter("."+cls).removeClass("ui-helper-hidden");
	});

	$("#bulkinsert").click(function() {
		if( $(this).prop("checked") )
			$("#bulklimit").removeAttr("disabled");
		else
		$("#bulklimit").attr("disabled","disabled");
	});
	$("#bulklimit").click(function() {
		if( $(this).prop("checked") )
			$("#bulksize").removeAttr("disabled");
		else
			$("#bulksize").attr("disabled","disabled");
	});
});
</script>
