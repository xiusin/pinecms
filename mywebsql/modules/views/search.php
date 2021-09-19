<link href='/mywebsql/cache?css=theme,default,alerts' rel="stylesheet" />

<style>
	div#db_objects { margin-top:5px;padding:3px;overflow:auto;height:330px;width:95%;border:3px double #efefef }
	div.objhead 	{ background-color:#ececec; padding: 5px; margin: 0 0 3px 0 }
	span.toggler 	{ display:inline-block; float:right; cursor: pointer; font-size:16px; margin: -5px 0 0 0 }
	div.obj 		{ padding:5px; margin:0 0 0 20px }
</style>

<div id="popup_wrapper">
	<div id="popup_contents">
		<table border="0" cellpadding="5" cellspacing="4" style="width: 100%;height:100%">
		<tr>
		<td align="left" valign="top" width="45%">
			<%= T("Select tables to search") %><br />
			<div id="db_objects">
				<%= T("Either the database is empty, or there was an error retrieving list of database objects") %>.<br/>
				<%= T("Please try closing and re-opening this dialog again") %>.
			<div>
		</td>

		<td align="left" valign="top" width="55%">
		<fieldset>
			<legend><%= T("Search Options") %></legend>
			<table border="0" cellspacing="10" cellpadding="5" width="100%">
			<tr><td valign="top">
			<label><%= T("Text to search") %>:</label><input type='text' name='keyword' id='keyword' />
			</td></tr>
				<tr><td valign="top"><%= T("Search in following field types") %>:</td></tr>
				<tr><td valign="top">
				<input class="ftype" type='checkbox' name='ftype_num' id='ftype_num' /><label class="right" for='ftype_num'><%= T("Numeric Fields") %>&nbsp;(INT, FLOAT...)</label>
				</td></tr>
				<tr><td valign="top">
				<input class="ftype" type='checkbox' name='ftype_char' id='ftype_char' checked="checked" /><label class="right" for='ftype_char'><%= T("Character Fields") %>&nbsp;(CHAR, VARCHAR...)</label>
				</td></tr>
				<tr><td valign="top">
				<input class="ftype" type='checkbox' name='ftype_text' id='ftype_text' checked="checked" /><label class="right" for='ftype_text'><%= T("Text Fields") %>&nbsp;(TEXT, LONGTEXT...)</label>
				</td></tr>
				<tr><td valign="top">
				<input class="ftype" type='checkbox' name='ftype_date' id='ftype_date' checked="checked" /><label class="right" for='ftype_date'><%= T("Date/Time Fields") %>&nbsp;(DATETIME, TIMESTAMP...)</label>
				</td></tr>
			</table>
		</fieldset>
		<fieldset>
			<legend><%= T("Comparison Type") %></legend>
			<table border="0" cellspacing="10" cellpadding="5" width="100%">
				<tr><td valign="top">
				<input type='radio' name='operator' value="equal" id='op_equal' checked="checked" /><label class="right" for='op_equal'><%= T("Equality Operator") %>&nbsp;(field = 'text')</label>
				</td></tr>
				<tr><td valign="top">
				<input type='radio' name='operator' value="like" id='op_like' /><label class="right" for='op_like'><%= T("Like Operator") %>&nbsp;(field like 'text')</label>
				</td></tr>
				<tr><td valign="top">
				<input type='radio' name='operator' value="wildcard" id='op_wildcard' /><label class="right" for='op_wildcard'><%= T("WildCard Operator") %>&nbsp;(field like '%text%')</label>
				</td></tr>
				<tr><td valign="top">
				<input type='radio' name='operator' value="greater" id='op_greater' /><label class="right" for='op_greater'><%= T("Greater Than") %>&nbsp;(field > 250)</label>
				</td></tr>
				<tr><td valign="top">
				<input type='radio' name='operator' value="lesser" id='op_lesser' /><label class="right" for='op_lesser'><%= T("Less Than") %>&nbsp;(field < 250)</label>
				</td></tr>
			</table>
		</fieldset>
		</td>
		</tr>
		</table>
	</div>
	<div id="popup_footer">
		<div id="popup_buttons">
			<input type='button' id="btn_search" value='<%= T("Search") %>' />
		</div>
	</div>
</div>

<script type="text/javascript" language='javascript' src="/mywebsql/cache?script=common,jquery,ui,query,options,alerts"></script>
<script type="text/javascript" language="javascript">
window.title = "<%= T("Search in Database") %>";
var tables = <%= TABLELIST %>;

function searchDatabase() {
	if ($('#keyword').val() == '') {
		jAlert(__('Enter the text to search in database'));
		return false;
	}
	if ($('#db_objects').find(':checked').length == 0) {
		jAlert(__('No Table selected'));
		return false;
	}
	if ($('.ftype:checked').length == 0) {
		jAlert(__('Select at least one field type for searching'));
		return false;
	}
	wrkfrmSubmit('search', '', '', '');
}

$(function() {
	if (tables.length == 0)
		return;

	$('#btn_search').button().click(searchDatabase);
	$('#db_objects').html('');
	uiShowObjectList(tables, 'tables', __('Tables'));

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

	$('.ftype').click(function() {
		if (this.id == 'ftype_num' && $(this).attr('checked')) {
			$('#ftype_char,#ftype_text,#ftype_date').removeAttr('checked');
		}
		if (this.id != 'ftype_num' && $(this).attr('checked')) {
			$('#ftype_num').removeAttr('checked');
		}
	});

	$('#keyword').focus();
});
</script>
<?php
	echo getGeneratedJS();
?>
