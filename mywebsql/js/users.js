/**
 * This file is a part of MyWebSQL package
 *
 * @file:      js/users.js
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

function selectUser() {
	user = $('#userlist').val();

	if (user == '')
		return;

	setMessage(__('Please wait...'));
	$('#popup_overlay').removeClass('ui-helper-hidden');
	wrkfrmSubmit('usermanager', '', '', user);
}

function setError(o, s) {
	$('#grid-messages').html(s).addClass('error');
	if ($(selectedRow).length) {
		$(selectedRow).removeClass('sel');
		selectedRow = null;
	}
	$(o).each( function() {
		$(this).addClass('error');
	});
	setTimeout( function() {
		$(o).each( function() {
			$(this).removeClass('error');
		});
	}, 2000);
}

function selectAll() {
	check = $(this).prop('checked');
	n = $('#grid-tabs').tabs('option', 'selected');
	cls = '';
	if (n == 1)
		cls = '.prv';
	else if (n == 2)
		cls = '.dbprv';
	if (cls != '')
		$(cls).prop('checked', check);

	$(cls).each((n==1) ? updatePrivilege : updateDbPrivilege);
}

function setMessage(s) {
	$('#grid-messages').html(s).removeClass('error');
}

function cancelOperation() {
	$('#grid-tabs').tabs('enable', 1).tabs('enable', 2);
	$('#userlist').removeAttr('disabled');
	$('#button-list').animate({'margin-top':'0px'});
	$('#btn_cancel,#btn_add2').hide();
	$('#btn_submit').show();
	loadUserData();
}

function loadUserData() {
	$('#username').val(USER_INFO['username']);
	$('#hostname').val(USER_INFO['host']);

	currentDb = null;

	$('#tab-global').html('');
	html = '';
	for(var index in PRIVILEGE_NAMES) {
		html += '<div class="input float" style="width:160px"><input class="prv" type="checkbox" id="' + index + '" name="' + index + '" ';
		if ($.inArray(index, PRIVILEGES) != -1)
			html += ' checked="checked"';
		html += '/><label class="right" for="' + index + '">' + PRIVILEGE_NAMES[index] + '</label></div>';
	}
	$('#tab-global').html(html);

	$('#tab-global .prv').click(updatePrivilege);

	showDbList();

	$('#checkboxes').hide();

	$('#popup_overlay').hide();

	$('#popup_wrapper').css('display', 'block');

	setTimeout( function() {
		$('#username').focus();
	}, 50 );
}

function showDbList() {
	html = '<div class="padded float" style="width:190px"><select name="db_names" id="db_names" size="13" style="width:170px">';
	list = DATABASES;
	for(i=0; i<list.length; i++) {
		cls = '';
		db = list[i];
		db_privileges = DB_PRIVILEGES[db] || [];
		if (db_privileges.length > 0 )
			cls = ' class="used"';
		name = htmlchars(db);
		html += '<option name="' + name + '" value="'+ name + '" ' + cls + '>' + name + '</option>';
	}
	html += '</select></div><div class="padded float" id="db_privileges" style="width:330px">'+__('Select a database to view privileges for the user') + '</div>';
	$('#tab-db').html(html);

	$('#db_names').change(showDbPrivileges);
	$('#db_names .used').css('font-weight', 'bold');
	$('#db_names').children('option').each(function(i) {
		if ($(this).hasClass('used')) {
			$('#db_names').prop('selectedIndex', i);
			$('#db_names').trigger('change');
			return false;
		}
	});

}

function showDbPrivileges() {
	$('#checkboxes').show();
	db = $('#db_names').val();
	currentDbPrivileges = DB_PRIVILEGES[db] || [];
	html = '';
	for(var index in DB_PRIVILEGE_NAMES) {
		html += '<div class="input float" style="width:140px"><input class="dbprv" type="checkbox" id="db_' + index + '" name="db_' + index + '" ';
		if ($.inArray(index, currentDbPrivileges) != -1)
			html += ' checked="checked"';
		html += '/><label class="right" for="db_' + index + '">' + DB_PRIVILEGE_NAMES[index] + '</label></div>';
	}
	$('#db_privileges').html(html);

	$('#tab-db .dbprv').click(updateDbPrivilege);

	$('#selectall').prop('checked', $('#tab-db .dbprv').not(':checked').length == 0);
}

function updatePrivilege() {
	prv = $(this).attr('id');
	checked = $(this).prop('checked');
	if (checked) {
		if ($.inArray(prv, PRIVILEGES) == -1)
			PRIVILEGES.push(prv);
	}
	else {
		for(i=0; i<PRIVILEGES.length; i++) {
			if (PRIVILEGES[i] == prv) {
				PRIVILEGES.splice(i, 1);
				break;
			}
		}
	}
}

function updateDbPrivilege() {
	prv = $(this).attr('id').substr(3);
	db = $('#db_names').val();
	checked = $(this).prop('checked');
	if (checked) {
		if ($.inArray(prv, DB_PRIVILEGES[db]) == -1)
			DB_PRIVILEGES[db].push(prv);
	}
	else {
		for(i=0; i<DB_PRIVILEGES[db].length; i++) {
			if (DB_PRIVILEGES[db][i] == prv) {
				DB_PRIVILEGES[db].splice(i, 1);
				break;
			}
		}
	}

	fw = DB_PRIVILEGES[db].length > 0 ? 'bold' : 'normal';
	$('#db_names option:selected').css('font-weight', fw);
}

function addUser() {
	$('#grid-messages').html('');
	$('#grid-tabs').tabs('select', 0).tabs('disable', 1).tabs('disable', 2);
	$('#username,#hostname,#userpass,#userpass2').val('');
	$('#userlist').attr('disabled', 'disabled');
	height = '-' + $('#button-list').outerHeight() + 'px';
	$('#button-list').animate({'margin-top':height});
	$('#btn_submit').hide();
	$('#btn_cancel,#btn_add2').show();
}

function deleteUser() {
	user = $('#userlist').val();

	if (user == '') {
		jAlert(__("Select a User"), __("User Manager"));
		return;
	}

	optionsConfirm(__('Are you sure you want to delete this user account?'), 'users.delete', function(result, id, confirm_always) {
		if (result) {
			if (confirm_always)
				optionsConfirmSave(id);
			data = JSON.stringify(USER_INFO);
			wrkfrmSubmit("usermanager", "delete", "", data);
		}
	});
}

function addNewUser() {
	username = $('#username').val();
	hostname = $('#hostname').val();
	password = $('#userpass').val();
	password2 = $('#userpass2').val();

	if (username == '' || hostname == '' || password == '') {
		jAlert(__("User information is incomplete or invalid"), __("User Manager"));
		return false;
	}

	if (password != password2) {
		jAlert(__("Passwords do not match"), __("User Manager"));
		return false;
	}

	json = {'username':username, 'hostname':hostname, 'pwd': password};
	query = JSON.stringify(json);

	setMessage(__('Please wait...'));
	$('#popup_overlay').removeClass('ui-helper-hidden');
	wrkfrmSubmit('usermanager', 'add', '', query);
}

function updateUser() {
	user = $('#userlist').val();

	if (user == '') {
		jAlert(__("Select a User"), __("User Manager"));
		return false;
	}

	username = $('#username').val();
	hostname = $('#hostname').val();
	password = $('#userpass').val();
	password2 = $('#userpass2').val();
	removepass = $('#nopass').prop('checked') ? '1' : '0';

	if (removepass == '1')
		password = password2 = '';

	if (username == '' || hostname == '') {
		jAlert(__("User information is incomplete or invalid"), __("User Manager"));
		return false;
	}

	if (password != password2) {
		jAlert(__("Passwords do not match"), __("User Manager"));
		return false;
	}

	json = {
		'oldusername': USER_INFO.username,
		'oldhostname': USER_INFO.host,
		'username': username,
		'hostname': hostname,
		'password': password,
		'removepass': removepass,
		'privileges': PRIVILEGES,
		'db_privileges': DB_PRIVILEGES
	};
	query = JSON.stringify(json);

	$('#popup_wrapper').css('display', 'none');
	setMessage(__('Please wait...'));
	wrkfrmSubmit('usermanager', 'update', '', query);
}
