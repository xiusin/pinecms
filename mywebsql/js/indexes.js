/**
 * This file is a part of MyWebSQL package
 *
 * @file:      js/indexes.js
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

var pendingChanges = false;
var bAddingIndex = false;
var newIndex = null;

function setError(o, s) {
	$('#grid-messages').html(s).addClass('error');
	if ($(selectedRow).length) {
		$(selectedRow).removeClass('sel');
		selectedRow = null;
	}
	$(o).each(function() { $(this).addClass('error'); });
	setTimeout(function() { $(o).each(function() { $(this).removeClass('error'); }); }, 2000);
}

function setMessage(s) {
	$('#grid-messages').html(s).removeClass('error');
}

function setupIndexes() {
	$('#grid-tabs').tabs();

	$('#indexlist option').remove();
	options = '';
	for (index in indexInfo)
		options += '<option value="' + index + '">' + index + '</option>';
	$('#indexlist').html(options).change(selectIndex);

	removeFieldList();

	$('#btn_addfield').button({ disabled: true }).click(addField).hide();
	$('#btn_delfield').button({ disabled: true }).click(deleteField).hide();
	$('#chk_primary,#chk_unique,#chk_fulltext').click(checkIndexOptions);

	$('#indextype').hide();

	$('#btn_add').button().click(addIndex);
	$('#btn_edit').button().click(editIndex).hide();
	$('#btn_save').button().click(saveIndex).hide();
	$('#btn_del').button().click(deleteIndex).hide();
	$('#btn_cancel').button().click(cancelChanges).hide();
	$('#btn_submit').button({ disabled: true }).click(validateIndexes);

	$("#dialog-list").dialog({
		autoOpen: false,
		width: 240,
		height: 240,
		modal: true,
		draggable: false,
		resizable: false,
		open: loadDialogValues,
		buttons: {
			'Add': addFieldToIndex,
			'Done': function () { $('#dialog-list').dialog('close'); }
		}
	});
}

function loadDialogValues(e, ui) {
	$('#list-items').html('');
	index_name = $('#indexlist').val();

	for(i=0; i<fieldInfo.length; i++) {
		fieldExists = false;
		fi = fieldInfo[i];
		// skip fields that are already part of index
		for(j=0; j<newIndex.length; j++) {
			if (newIndex[j].column == fi.fname) {
				fieldExists = true;
				break;
			}
		}
		if (fieldExists)
			continue;
		txt = fi.flen == '' ? (fi.fname + '  [ ' + fi.ftype + ' ]') : (fi.fname + '  [' + fi.ftype + ' (' + fi.flen + ')]');
		option = $('<option></option>').val(fi.fname).text(txt);
		$('#list-items').append(option);
	}
	setTimeout(function() { $('#list-item').focus(); }, 50 );
}

function addFieldToIndex() {
	if ($('#list-items option:selected').length == 0)
		return false;

	field = $('#list-items').val();
	flength = $('#flength').val();
	fieldObj = {"column":field, type:"BTREE", order:0, length: flength};
	index_name = $('#indexlist').val();
	newIndex.push(fieldObj);

	fi = getField(field);
	tr = $('<tr><td><input type="checkbox" /></td><td></td><td></td></tr>');
	td = tr.find('td');
	ftype = fi.flen == '' ? fi.ftype : fi.ftype + ' (' + fi.flen + ')';
	td.eq(1).text(fi.fname);
	td.eq(2).text(ftype);
	$('#table_grid tbody').append(tr);

	$('#list-items option:selected').remove();
	$('#flength').val('');

	$('#btn_save').button({ disabled: false });
}

function addIndex() {
	jPrompt(__('Enter new index name'), '', __('Indexes'), function(new_name) {
		if (new_name == null)
			return;
		else if (new_name && new_name != '' &&!indexExists(new_name)) {
			newIndex = [];
			bAddingIndex = true;
			option = $('<option value="' + new_name + '">' + new_name + '</option>');
			$('#indexlist').append(option);
			option.prop('selected', 'selected');
			selectIndex();
			editIndex();
		}
	});
}

function editIndex() {
	index = $('#indexlist').val();
	if(!bAddingIndex)
		newIndex = clone(indexInfo[index]);

	$('#btn_add,#btn_edit,#btn_del').hide();
	$('#table_grid').find('input').removeAttr('disabled');
	$('#indexlist').attr('disabled', 'disabled');
	$('#btn_addfield').button({ disabled: false }).click(addField).show();
	$('#btn_delfield').button({ disabled: false }).click(deleteField).show();
	$('#btn_cancel').show();
	$('#btn_save').button({ disabled: true }).show();
	$('#btn_submit').button({ disabled: true });

	$('#indextype').find('input').removeAttr('disabled');
	// if we have a primary key, we need to disable primary checkbox
	if (indexExists('PRIMARY'))
		$('#chk_primary').attr('disabled', 'disabled');
	if (index == 'PRIMARY')
		$('#chk_unique,#chk_fulltext').attr('disabled', 'disabled');
}

function deleteIndex()
{
	index = $('#indexlist').val();
	if (index == '')
		return;

	// are we discarding a newly created index or deleting existing one?
	if (!bAddingIndex) {
		delete indexInfo[index_name];
	}

	$('#indexlist option:selected').remove();
	removeFieldList();

	$('#indextype').hide();
	$('#btn_addfield').hide();
	$('#btn_delfield').hide();

	$('#btn_edit').hide();
	$('#btn_del').hide();
	$('#btn_submit').button({ disabled: false });
}

function saveIndex() {
	if(newIndex.length == 0) {
		jAlert(__('Index must have at least one field'));
		return false;
	}

	// if we want to create a primary key, then it's name must be changed also
	if (bAddingIndex && newIndex.primary == '1') {
		index_name = 'PRIMARY';
		delete newIndex.primary;
		$('#indexlist option:selected').val(index_name).text(index_name);
	}
	index_name = $('#indexlist').val();
	indexInfo[index_name] = newIndex;

	$('#btn_add').show();
	$('#btn_cancel').hide();
	$('#btn_save').hide();
	$('#indexlist').removeAttr('disabled');

	bAddingIndex = false;
	pendingChanges = true;
	$('#btn_submit').button({disabled:false});

	selectIndex();
}

function selectIndex() {
	index_name = $('#indexlist').val();
	if (index_name == '')
		return;

	$('#table_grid tr').not('#fhead').remove();

	primary = unique = fulltext = false;
	index = bAddingIndex ? newIndex : indexInfo[index_name];
	for(i=0; i<index.length; i++) {
		fi = getField(index[i].column);
		tr = $('<tr><td><input type="checkbox" /></td><td></td><td></td></tr>');
		td = tr.find('td');
		ftype = fi.flen == '' ? fi.ftype : fi.ftype + ' (' + fi.flen + ')';
		td.eq(1).text(fi.fname);
		td.eq(2).text(ftype);
		$('#table_grid tbody').append(tr);

		// set type of index
		if (index_name == 'PRIMARY')
			primary = true;
		if (index[i].unique == '1')
			unique = true;
		if (index[i].type == 'FULLTEXT')
			fulltext = true;
	}

	$('#chk_primary').prop('checked', primary);
	$('#chk_unique').prop('checked', unique);
	$('#chk_fulltext').prop('checked', fulltext);
	$('#table_grid').find('input').attr('disabled', 'disabled');
	$('#indextype').show().find('input').attr('disabled', 'disabled');
	$('#btn_addfield').button({ disabled: true }).show();
	$('#btn_delfield').button({ disabled: true }).show();
	$('#btn_edit').show();
	$('#btn_del').show();
}

function removeFieldList() {
	$('#table_grid tr').not('#fhead').remove();
	$('#table_grid tbody').append('<tr><td class="empty" colspan="3">' + __('Select an index to view / edit its details') + '</td></tr>');
}

function addField() {
	$('#dialog-list').dialog('open');
}

function deleteField() {
	checked = $('#table_grid input:checked');
	if (checked.length == 0)
		return false;

	index_name = $('#indexlist').val();
	delRows = [];
	checked.each(function() {
		field = $(this).parent().next().text();
		delRows.push($(this).parent().parent());
		for(i=0; i<newIndex.length; i++) {
			if (newIndex[i].column == field) {
				newIndex.splice(i, 1);
			}
		}
	});

	for(i=0;i<delRows.length;i++)
		delRows[i].remove();

	pendingChanges = true;
	$('#btn_save').button({ disabled: false });
}

function checkIndexOptions() {
	checked = $(this).prop('checked');
	id = $(this).attr('id');
	index = $('#indexlist').val();

	if (id == 'chk_primary' && checked) {
		if(indexExists('PRIMARY')) {
			jAlert(__('Primary key already exists'));
			$(this).removeAttr('checked');
			return false;
		}
	}

	if (id == 'chk_primary')
		newIndex.primary = checked ? '1' : '0';

	for(i=0; i<newIndex.length; i++) {
		if (id == 'chk_unique')
			newIndex[i].unique = checked ? '1' : '0';
		else if (id == 'chk_fulltext')
			newIndex[i].type = checked ? 'FULLTEXT' : 'BTREE';
	}
	$('#btn_save').button({ disabled: false });
	return true;
}

function cancelChanges() {
	$('#btn_add').show();
	$('#btn_cancel').hide();
	$('#btn_save').hide();
	$('#indexlist').removeAttr('disabled');

	if (bAddingIndex) {
		deleteIndex(true);
		bAddingIndex = false;
		removeFieldList();
	} else
		selectIndex();

	if (pendingChanges)
		$('#btn_submit').button({ disabled: false });
}

function getField(n) {
	for(j=0; j<fieldInfo.length; j++)
		if (fieldInfo[j].fname == n)
			return fieldInfo[j];

	return {};
}

function indexExists(name) {
	for(index_name in indexInfo) {
		if (index_name == name)
			return true;
	}
	return false;
}

function validateIndexes() {
	json = {};
	json.indexes = indexInfo;

	query = JSON.stringify(json);

	setMessage('Please wait...');
	$('#popup_overlay').removeClass('ui-helper-hidden');
	wrkfrmSubmit('indexes', 'alter', tableName, query, responseHandler);
}

function responseHandler(data) {
	result = $(data).find('#result').text();
	message = $(data).find('#message').html();
	if (result == '1') {
		setMessage(__('Indexes Updated'));
		$('#tab-messages').html(message);
		pendingChanges = false;
		$('#btn_submit').button({disabled:false});
	}
	else {
		setMessage(__('Error'));
		$('#tab-messages').html(message);
		$("#grid-tabs").tabs('select', 1);
	}
	div = $('#tab-messages div.sql_text').length > 0 ? $('#tab-messages div.sql_text') : $('#tab-messages div.sql_error');
	if (div.length) {
		code = div.html2txt();
		obj_lines = $('<div class="sql_lines"></div>');
		obj_out = $('<pre class="sql_output"></pre>');
		div.html('').append(obj_lines).append(obj_out);
		parent.commandEditor.win.highlightSql($('#tab-messages pre.sql_output'), $('#tab-messages div.sql_lines'), code);
	}

	$('#popup_overlay').addClass('ui-helper-hidden');
}

function clone(obj) {
	var temp = new obj.constructor();
	for(var key in obj)
		temp[key] = obj[key];
	return temp;
}