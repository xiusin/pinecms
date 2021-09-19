/**
 * This file is a part of MyWebSQL package
 *
 * @file:      js/query.js
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

var resultInfo = "";          // saves info of edited result set during editing
var editTableName = "";       // name of editing table from where the result came up
var currentQuery = "";			// used for refresh command in the interface
var queryID = '';             // unique query id
var editKey = [];             // unique or primary key(s) for editing
var totalRecords = 0;         // number of records fetched (global query)
var totalPages = 1;           // number of pages to show
var currentPage = 1;          // currently displayed page
var numRecords = 0;           // number of records fetched on this page

var sql_delimiter = ";\n";

function getResults(i) {
	xfrm = getFrame();
	if (i == 0) {
		if (xfrm.document.getElementById('results').innerText)		// IE
			x = xfrm.document.getElementById('results').innerText;
		else if (xfrm.document.getElementById('results').textContent)	//firefox
			x = xfrm.document.getElementById('results').textContent;
		else
			x = xfrm.document.getElementById('results').innerHTML;
	}
	else if (i == 1)
		x = xfrm.document.getElementById('results').innerHTML;
	else if (i == 2)
		x = xfrm.document.getElementById('title').innerHTML;

	return x;
}

function queryGo(type) {
	strQuery = getSqlCodeSelection();
	strMsg = __("all selected");
	if (strQuery == "") {
		strQuery = getSqlCode();
		strMsg = __("all");
	}

	if ($.trim(strQuery) == "") {
		jAlert(__("Please type in one or more queries in the sql editor!"), __("Execute query"), function() { focusEditor(); });
		return;
	}
	querySaveCache(strQuery);

	if (type == 1) {
		msg = str_replace('{{SELECTED}}', strMsg, __('Are you sure you want to execute {{SELECTED}} queries?'));
		optionsConfirm(msg, 'query.all', function(result, id, confirm_always) {
			if (result)
			{
				if (confirm_always) optionsConfirmSave(id);
				wrkfrmSubmit("queryall", "", "", strQuery);
			}
		});
	}
	else
		wrkfrmSubmit("query", "", "", strQuery);

	focusEditor();
}

function queryDelete() {
	qBig = "";
	discardRows = []; // rows that are newly added and now being deleted
	checked = $('#result_form input:checked').not('.check-all');

	if (checked.length == 0)
		return "";

	checked.each(function() {
		row = $(this).parent().parent();
		// newly added rows can be simply discarded
		if(row.hasClass('n'))
			discardRows.push(row);
		else
			qBig += makeDeleteClause(row);
	});

	checked.prop('checked', false);
	$('#dataTable input.check-all').prop('checked', false);

	$(discardRows).each(function() {
		this.remove();
	});

	// return or execute queries
	if (arguments.length > 0 && arguments[0] == true)
		return qBig;

	showNavBtns('query', 'queryall');
	// do we have any live records to delete, or only discardables
	if (qBig != "")
		wrkfrmSubmit("queryall", "", "", qBig);
}

function querySave() {
	editRows = $('#dataTable tbody tr.x').add('#dataTable tbody tr.n');
	if (editRows.length == 0)
		return "";

	qBig = "";
	editRows.each(function() {
		newRecord = $(this).hasClass('n');
		// generate new and updated records queries separately
		var upd = "";
		var ins_fields = [];
		var ins_data = [];
		editCols = $(this).find('td.x');
		editCols.each(function() {
			data = $(this).data('edit');
			f = getFieldName($(this).index()-2);				
			if (newRecord) {
				if (typeof data != "object") {
					data = {'setNull':true, 'value':''};
				}
				ins_fields.push(DB.quote(f));
				ins_data.push( data.setNull ? "NULL" : DB.escape( data.value ) );
			} else if (typeof data == "object") {
				if (data.setNull)
					upd += DB.quote(f) + "=" + "NULL,";
				else
					upd += DB.quote(f) + "=" + DB.escape( data.value ) + ',';
			}
		});
		// only generate query if at least one field is added (for new records)
		if (newRecord && ins_fields.length > 0) {
			qBig += sql_delimiter + "INSERT INTO " + DB.quote(editTableName) + " (" + ins_fields.join(',') + ") VALUES (" + ins_data.join(',') + ")";
		}
		// only generate query if at least one field is updated (for existing records)
		if(upd !== "") {
			upd = "UPDATE " + DB.quote(editTableName) + " SET " + upd.substr(0, upd.length-1);
			qBig += sql_delimiter + upd + makeWhereClause($(this));
		}
	});

	$('#dataTable tbody .x').add('#dataTable tbody .n').removeClass('x n').data('edit', null);

	// return or execute queries
	if (arguments.length > 0 && arguments[0] == true)
		return qBig;

	showNavBtns('query', 'queryall');
	wrkfrmSubmit("queryall", "", "", qBig);
}

function queryGenerate() {
	del = queryDelete(true);  // we must get the delete queries first, then save queries
	sql = querySave(true) + del;
	setSqlCode(sql, 1);
	showNavBtns('query', 'queryall');
}

function queryAddRecord() {
	num = $('#dataTable tbody tr').length+1;
	row = '<tr class="row n"><td class="tj">'+num+'</td><td class="tch"><input type="checkbox" /></td>';
	for(i=0;i<fieldInfo.length;i++) {
		txt = fieldInfo[i].blob ? '<span class="data">NULL</span>' : 'NULL';
		cls = fieldInfo[i].blob ? 'edit blob tnl' : 'edit tnl';
		row += '<td class="'+cls+'" nowrap="wrap">'+txt+'</td>';
	}
	row += '</tr>';
	row = $(row);
	row.find('td.edit').bind(editOptions.editEvent, editOptions.editFunc);
	row.find('input').click(function() { showNavBtn('delete', 'gensql'); });
	$('#dataTable tbody').append(row);
	$(".ui-layout-data-center").tabs('select', 0);
	setTimeout(function() { bottom = $("#results-div").prop("scrollHeight"); $("#results-div").prop("scrollTop", bottom);}, 50);
}

function queryCopyRecord() {
	checked = $('#result_form input:checked').not('.check-all');

	if (checked.length == 0)
		return "";

	checked.each(function() {
		old_row = $(this).parent().parent();
		num = $('#dataTable tbody tr').length+1;
		row = $(old_row).clone(true, true);
		row.addClass("n").removeClass("ui-state-active");
		td = row.children("td");
		td.find("input").prop("checked", false);
		td.eq(0).text(num);
		for(i=0;i<fieldInfo.length;i++) {
			current = td.eq(i+2);
			current.addClass("x");
			// if field is blob, we set it to NULL
			is_null = fieldInfo[i].autoinc || current.find("span.blob").length;
			if (is_null)
				current.text("NULL").addClass("tnl");

			// if text data, copy it internally
			txt = current.find("span.d").length ? current.find("span.d").text() : current.html();
			data = {'setNull':is_null, 'value':txt};
			current.data('edit', data);
		}

		$('#dataTable tbody').append(row);
	});

	checked.prop('checked', false);
	$('#dataTable input.check-all').prop('checked', false);
	showNavBtn('update', 'gensql');
	// select the last row to have some visual indication of what happened
	row.trigger('click');
	$(".ui-layout-data-center").tabs('select', 0);
	setTimeout(function() { bottom = $("#results-div").prop("scrollHeight"); $("#results-div").prop("scrollTop", bottom);}, 50);
}

function queryRefresh() {
	if (currentQuery == "") {
		jAlert(__("Failed to refresh the results."), __("Refresh results"), function() { focusEditor(); });
		return false;
	}

	wrkfrmSubmit("query", "", "", currentQuery);
	focusEditor();
}

/*
 * @@TODO: finalize this feature
 * function queryFormat() {
	strQuery = getSqlCode();

	if ($.trim(strQuery) == "") {
		jAlert(__("Please type in one or more queries in the sql editor!"), __("Format SQL"), function() { focusEditor(); });
		return;
	}

	editor = currentEditor();
	range = { from: editor.getCursor(true), to: editor.getCursor(false) };
	editor.autoFormatRange(range.from, range.to);
}*/

function transferQuery() {
	htm = sql_delimiter + getResults(0);
	setSqlCode(htm, 1);

	setPageStatus(false);

	commandEditor.focus();
}

function transferResultMessage(num, tm, msg) {
	resultInfo = "";
	$("#messages-div").html(getResults(1));
	document.getElementById("timeCounter").innerHTML = tm;

	$(".ui-layout-data-center").tabs('select', 1);
	$("#messages-div").prop("scrollTop", 0).prop("scrollLeft", 0);

	if (commandEditor.canHighlight()) {
		div = $('#messages-div div.sql-text');
		if(div.length > 0) {
			code = div.html2txt();
			obj_lines = $('<div class="sql_lines"></div>');
			obj_out = $('<pre class="sql_output cm-s-default"></pre>');
			div.html('').append(obj_lines).append(obj_out);
			commandEditor.highlightSql($('#messages-div pre.sql_output'), $('#messages-div div.sql_lines'), code);
		}
	}

	setPageStatus(false);
	showNavBtns('query', 'queryall');
}

function transferInfoMessage() {
	resultInfo = "";
	$("#info-div").html(getResults(1));
	$("#tab-info > .message").remove();
	$("#info-div .message").clone().prependTo("#tab-info");
	$("#info-div .message").remove();

	$(".ui-layout-data-center").tabs('select', 2);
	$("#info-div").attr("scrollTop", 0).prop("scrollLeft", 0);

	if (commandEditor.canHighlight()) {
		div = $('#info-div div.sql-text');
		if(div.length > 0) {
			code = div.html2txt();
			obj_lines = $('<div class="sql_lines"></div>');
			obj_out = $('<pre class="sql_output cm-s-default"></pre>');
			div.html('').append(obj_lines).append(obj_out);
			commandEditor.highlightSql($('#info-div pre.sql_output'), $('#info-div div.sql_lines'), code);
		}
	}

	if ($('#infoTable').length > 0)
		setupTable('infoTable', {highlight:true,selectable:true,editable:false,sortable:'inline'});

	setPageStatus(false);
	showNavBtns('query', 'queryall');

	$("#quick-info-search").bind('keyup', function() {
		$("#infoTable").setSearchFilter( $(this).val() );
		resizeTableHeader('info');
	});
		
	resizeTableHeader('info');
}

function transferResultGrid(num, tm, msg) {
	resultInfo = "";
	// must add form here. FireFox removes form element if taken from div's innerHTML (no idea why)
	$("#results-div").html("<form id='result_form' name='resForm' method='post' action='#' onsubmit='return false;'>" + getResults(1) + "</form>");
	if (num == -1)
		document.getElementById("recordCounter").innerHTML = "&nbsp;";
	else
		document.getElementById("recordCounter").innerHTML = totalRecords + " record(s)";
	document.getElementById("timeCounter").innerHTML = tm;
	document.getElementById("messageContainer").innerHTML = msg;

	$("#headerResults").html(getResults(2));

	numRecords = num;

	if(num > 0)
		setupResults();

	$(".ui-layout-data-center").tabs('select', 0);
	$("#results-div").prop("scrollTop", 0).prop("scrollLeft", 0);
		


	if (totalPages > 1) {
		str = __('Results page:') + '&nbsp;';
		str += (1==currentPage) ? '<span class="page">' + __('Previous') + '</span>' : '<a href="javascript:goPage('+(currentPage-1)+')">' + __('Previous') + '</a>';
		str += '<select id="page_selector" name="page" onchange="javascript:goPage(this.value)">';
		for(i=1; i<=totalPages; i++)
			str += (i==currentPage) ? '<option selected="selected" value="'+i+'">'+i+'</option>' : '<option value="'+i+'">'+i+'</option>';
		str += '</select>';
		str += (totalPages==currentPage) ? '<span class="page">' + __('Next') + '</span>' : '<a href="javascript:goPage('+(currentPage+1)+')">' + __('Next') + '</a>';
		$('#pagingContainer').html(str);
	}
	else
		$('#pagingContainer').html('');

	setPageStatus(false);
	editTableName == "" ? showNavBtns('query', 'queryall') : showNavBtns('addrec', 'query', 'queryall');
		
	// sometimes the cloned header has width problem, so to make sure we always see nice header, we resize it just after creating it
	resizeTableHeader('data');
}

function getFieldInfo(num) {
	return fieldInfo[num];
}

function getFieldName(num) {
	return fieldInfo[num]['name'];
}

function makeWhereClause(row) {
	str = " WHERE ";
	if (editKey.length == 0) {
		for(i=0; i<fieldInfo.length; i++) {
			td = row.find('td').eq(i+2);
			if (td.find('span.blob').length)	// for text/blobs this is true
					continue;
			var val = '';
			if (td.data('defText'))
				val = "=" + DB.escape(td.data('defText'));
			else {
				if(td.hasClass('tnl')) {
					val = ' is NULL';
				} else {
					var span = td.find('span.d');
					if (span.length)
						val =  "=" + DB.escape(span.text());
					else
						val = "=" + DB.escape(td.text());
				}
			}
			str += DB.quote( getFieldName(i) ) + val + " and ";
		}
	}
	else {
		for(i=0; i<fieldInfo.length; i++) {
			if (editKey.indexOf(getFieldName(i)) != -1) {
				var val = '';
				td = row.find('td').eq(i+2);
				if(td.hasClass('tnl')) {
					val = ' is NULL';
				} else {
					if (td.data('defText'))
						val = "=" + DB.escape(td.data('defText'));
					else
						val = "=" + DB.escape(td.text());
				}
				str += DB.quote( getFieldName(i) ) + val + " and ";
			}
		}
	}

	return str.substr(0, str.length - 5);
}

function makeDeleteClause(row) {
	str = sql_delimiter + "delete from " + DB.quote( editTableName ) + " " + makeWhereClause(row);
	return str;
}

/* *************************************************** */
function querySaveCache() {
	var editor = "sql_commandEditor";
	n = $(".ui-layout-data-south").tabs('option', 'selected');
	switch(n) {
		case 1: editor = "sql_commandEditor2"; break;
		case 2: editor = "sql_commandEditor3"; break;
	}
	$.cookies.set(editor, getSqlCode(), {path: EXTERNAL_PATH});
}

function vwBlb(obj, num, btype) {
	span = $(obj);
	fi = getFieldInfo(span.parent('td').index() - (editTableName == "" ? 1 : 2));
	name = fi['name'];
	tr = span.parent().parent();
	taskbar.openModal("blob-editor", "?q=wrkfrm&type=viewblob&id="+num+"&name="+name+"&blobtype="+btype+"&query="+queryID, 500, 300);
}
function vwTxt(obj, len, btype) {
	msg = str_replace('{{SIZE}}', len, __('Text Data [{{SIZE}}]'));
	span = $(obj).siblings('span.d');
	fi = getFieldInfo(span.parent('td').index() - (editTableName == "" ? 1 : 2));
	name = fi['name'];
	jAlert(span.html(), msg);
}

function loadUserPreferences() {
}

/* ******************************** */
function resultSelectAll() {
	check = $('#dataHeader input.check-all').prop('checked');
	$('#dataTable input').prop('checked', check);
	check ? showNavBtn('delete', 'gensql') : hideNavBtn('delete', 'gensql');
}

function getSqlCode() {
	ed = currentEditor();
	return ed.getCode();
}

function setSqlCode(s, u, n) {
	commandEditor.setCode( ((u==1) ? commandEditor.getCode() + s : s) );
	if (u==1)
		commandEditor.jumpToLine(commandEditor.lastLine());
	switchEditor(0);
}

function getSqlCodeSelection() {
	ed = currentEditor();
	return ed.getSelection();
}

/* ********************************* */
function getSelRecCount() {
	checked = $('#dataTable input:checked').not('.check-all');
	return checked.length;
}

function getHistoryText(single) {
	txt = "";
	if (single && !historyCurItem)
		return txt;
	obj = single ? historyCurItem.find('td.hst') : $('#sql-history .hst');
	for(i=0;i<obj.length; i++) {
		txt += $(obj[i]).html().replace(/(<br\s*[\/]?>)+/gi, "\n") + "\n";
	}
	return txt;
}

function historyClear() {
	optionsConfirm(__("Clear command history?"), 'history.clear', function(result, id, confirm_always) {
		if (result) {
			if (confirm_always) optionsConfirmSave(id);
			$("#sql-history").html("<tbody></tbody>");
		}
	});
}

function setupResults() {
	setupTable('dataTable', {highlight:true,selectable:true,editable:(editTableName != ""),sortable:true});

	if (editTableName != "")
		$('#dataTable input').not('check-all').click(function() { showNavBtn('delete', 'copyrec', 'gensql'); });

	//$("#dataTable").contextMenu(getDataMenu);
}

function goPage(num) {
	if(queryType == "table") {
		wrkfrmSubmit("query", "table", num, editTableName);
	} else {
		wrkfrmSubmit("query", "query", num, "");
	}
}

function goSort(field) {
	wrkfrmSubmit("query", "sort", field, '');
}

$.fn.html2txt = function() {
	txt = $(this).html().replace(/<br>/ig, "\n").replace(/&amp;/g,'&').replace(/&lt;/g,'<').replace(/&gt;/g,'>').replace(/&quot;/g,'"');
	return $.trim(txt);
};

$.fn.outerHTML = function(s) {
	return (s) ? this.before(s).remove() : jQuery("<p>").append(this.eq(0).clone()).html();
}

$.getSelectedText = function() {
	if(window.getSelection) { return window.getSelection(); }
	else if(document.getSelection) { return document.getSelection(); }
	else {
		var selection = document.selection && document.selection.createRange();
		if(selection.text) { return selection.text; }
		return '';
	}
	return '';
}
