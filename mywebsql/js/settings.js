/**
 * This file is a part of MyWebSQL package
 *
 * @file:      settings.js
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

/* works for confirmation dialogs transparently */
function optionsSave() {
	for(i=0; i<arguments.length; i++) {
		arg = arguments[i];
		obj = document.getElementById(arg);
		// ! may warn user here that some options could not be saved
		if (obj) {
			if (obj.type == "checkbox")
				val = obj.checked ? "on" : "";
			else
				val = obj.value;
			$.cookies.set("prf_"+arg, val, {path: EXTERNAL_PATH, hoursToLive: COOKIE_LIFETIME});
		}
	}

	jAlert(__("New settings saved and applied."));
}

/* works for general purpose settings to be set */
function optionsSet(name, val) {
	$.cookies.set(name, val, {path: EXTERNAL_PATH, hoursToLive: COOKIE_LIFETIME});
}

function optionsGet(name) {
	return $.cookies.get(name);
}

function optionsConfirm(msg, id, callback) {
	ask = $.cookies.get("prf_cnf_"+id);
	if (ask == 'no')
		return callback(true, '', false);
	return jConfirm(msg, __('Confirm Action'), callback, id);
}

function optionsConfirmSave(id) {
	$.cookies.set("prf_cnf_"+id, 'no', {path: EXTERNAL_PATH, hoursToLive: COOKIE_LIFETIME});
}

// resets all confirmation dialogs
function optionsReset() {
	$.each($.cookies.filter("prf_.*"), function(c) {
		$.cookies.set( c, false );
	});
}