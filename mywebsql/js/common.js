/**
 * This file is a part of MyWebSQL package
 *
 * @file:      js/common.js
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

function wrkfrmSubmit(type, id, name, query) {
	if (!document.frmquery) {
		xfrm = getFrame();
		frm = xfrm.document.frmquery;
		xfrm.onerror = frameErrorHandler;
		setPageStatus(true);
	}
	else {	// function called from an iframe in the popup
		frm = document.frmquery;
		$('#popup_overlay').removeClass('ui-helper-hidden');
	}

	frm.type.value = type;
	frm.id.value = id;
	frm.name.value = name;
	frm.query.value = query;

	if (arguments.length <= 4)
		frm.submit();
	else {
		callback = arguments[4];
		data = 'q=wrkfrm&' + $(frm).serialize();
		$.ajax({ type: 'POST',
			url: '?',
			data: data,
			success: callback
		});
	}
}

function getFrame() {
	xfrm = null;
	if ($.browser.msie && $.browser.version < 9.0)
		xfrm = document.frames("wrkfrm");
	else
		xfrm = window.frames["wrkfrm"];
	return xfrm;
}

function resetFrame() {
	xfrm = getFrame();
	xfrm.src = "javascript:false";
}

function debugMsg(msg) {
	$('#messageContainer').innerHTML = msg;
}

function frameErrorHandler() {
	setPageStatus(false);
	$('#recordCounter').html('&nbsp;');
	$('#timeCounter').html('');
	$('#messageContainer').html('Navigation Error. Try reloading the page');
}

function setPageStatus(flg, msg) {
	if (flg) {
		$('#nav_bar').css('display', 'none');
		$('#loader').css('display', 'table');
		showNavBtns();		// hide all buttons
	}
	else {
		$('#loader').css('display', 'none');
		$('#nav_bar').css('display', 'table-row');
		showNavBtns('query', 'queryall');
	}

	if (msg)
		$('#messageContainer').html(msg);
}

function addCmdHistory(str) {
	d = new Date();
	h = d.getHours();
	m = d.getMinutes();
	if (h < 10) h = '0'+h;
	if (m < 10) m = '0'+m;
	$('#sql-history > tbody:last').append("<tr><td valign=\"top\" class=\"dt\">[" + h + ":" + m + "]</td><td class=\"hst\">" + str + ";</td></tr>");

	if (arguments.length > 1 && arguments[1] == true)
		currentQuery = str;
}

function str_replace(search, replace, subject) {
	var f = search, r = replace, s = "" + subject;
	var ra = is_array(r), sa = is_array(s), f = [].concat(f), r = [].concat(r), i = (s = [].concat(s)).length;

	while (j = 0, i--) {
		while (s[i] = s[i].split(f[j]).join(ra ? r[j] || "" : r[0]), ++j in f){};
	};

	return sa ? s : s[0];
}

function is_array( mixed_var ) {
	return ( mixed_var instanceof Array );
}

/* simple hack for IE :P */
if(!Array.indexOf) {
	Array.prototype.indexOf = function(obj) {
		for(var i=0; i<this.length; i++) {
			if(this[i]==obj){
				return i;
			}
		}
		return -1;
	}
}

function __(txt) {
	if (window.lang && window['lang'][txt])
		return window['lang'][txt];
	return txt;
}

function formatBytes(val) {
	if (val < 1024)
		return val+' B';
	size = (val < 1024*1024) ? number_format(val/1024)+' KB' : number_format(val/(1024*1024), 2)+' MB';
	return size;
}

/* number_format function by phpjs */
function number_format (number, decimals, dec_point, thousands_sep) {
    number = (number + '').replace(/[^0-9+\-Ee.]/g, '');
    var n = !isFinite(+number) ? 0 : +number,
        prec = !isFinite(+decimals) ? 0 : Math.abs(decimals), sep = (typeof thousands_sep === 'undefined') ? ',' : thousands_sep,
        dec = (typeof dec_point === 'undefined') ? '.' : dec_point,
        s = '', toFixedFix = function (n, prec) {
            var k = Math.pow(10, prec);            return '' + Math.round(n * k) / k;
        };
    // Fix for IE parseFloat(0.55).toFixed(0) = 0;
    s = (prec ? toFixedFix(n, prec) : '' + Math.round(n)).split('.');
    if (s[0].length > 3) {        s[0] = s[0].replace(/\B(?=(?:\d{3})+(?!\d))/g, sep);
    }
    if ((s[1] || '').length < prec) {
        s[1] = s[1] || '';
        s[1] += new Array(prec - s[1].length + 1).join('0');    }
    return s.join(dec);
}

/* htmlspecialchars function by phpjs (name changed) */
function htmlchars (string, quote_style, charset, double_encode) {
	var optTemp = 0,
	i = 0, noquotes = false;
	if (typeof quote_style === 'undefined' || quote_style === null) {
		quote_style = 2;
	}
	string = string.toString();
	if (double_encode !== false) { // Put this first to avoid double-encoding
		string = string.replace(/&/g, '&amp;');
	}
	string = string.replace(/</g, '&lt;').replace(/>/g, '&gt;');
	var OPTS = {
		'ENT_NOQUOTES': 0,
		'ENT_HTML_QUOTE_SINGLE': 1,
		'ENT_HTML_QUOTE_DOUBLE': 2,
		'ENT_COMPAT': 2,
		'ENT_QUOTES': 3,
		'ENT_IGNORE': 4
	};
	if (quote_style === 0) {
		noquotes = true;
	}
	if (typeof quote_style !== 'number') { // Allow for a single string or an array of string flags
		quote_style = [].concat(quote_style);
		for (i = 0; i < quote_style.length; i++) {
			// Resolve string input to bitwise e.g. 'PATHINFO_EXTENSION' becomes 4
			if (OPTS[quote_style[i]] === 0) {
				noquotes = true;
			} else if (OPTS[quote_style[i]]) {
				optTemp = optTemp | OPTS[quote_style[i]];
			}
		}
		quote_style = optTemp;
	}
	if (quote_style & OPTS.ENT_HTML_QUOTE_SINGLE) {
		string = string.replace(/'/g, '&#039;');
	}
	if (!noquotes) {
		string = string.replace(/"/g, '&quot;');
	}
	return string;
}

function uiShowStatus(progress, type, id, delay) {
	var status = $(progress).data("status");
	if (!status) {
		status = window.setTimeout( function() { uiShowStatus(progress, type, id, delay); }, delay );
		$(progress).data("status", status);
		return true;
	}

	$.ajax({ type: 'GET',
		url: 'status.php?type=' + type + '&id=' + id,
		success: function(res) {
			if(res && res.c) {
				$(progress).progressbar("value", res.c);
				if (res.c >= 100) {
					$(progress).progressbar("destory");
					$(progress).removeData("status");
				}
			}
			status = window.setTimeout( function() { uiShowStatus(progress, type, id, delay); }, delay );
			$(progress).data("status", status);
		},
		error: function() {
			$(progress).progressbar("destory");
			$(progress).removeData("status");
		},
		dataType: 'json'
	});
}

function uiShowObjectList(list, name, title, uncheck)
{
	// objects other than schema are enclosed inside their schema container, so we make a flat list
	// of objects prefixed by the schema name
	if (!list.length) {
		new_list = [];
		for (var i in list) {
			for(j=0; j<list[i].length; j++) {
				new_list.push( i + "." + list[i][j] );
			}
		}
		list = new_list;
	}

	html = '';
	for(i=0; i<list.length; i++)
	{
		table = list[i];
		id = str_replace(/[\s\"']/, '', table);
		value = str_replace(/[\"]/, '&quot', table);
		html += '<div class="obj"><input' + (!uncheck ? ' checked="checked"' : '') + ' type="checkbox" name="' + name + '[]" id="' + name + '_' + id + '" value="'
				+ value + '" /><label class="right" for="' + name + '_' + id + '">' + table + '</label></div>';
	}
	if (html != '')
	{
		html = '<div class="objhead ui-widget-header"><input' + (!uncheck ? ' checked="checked"' : '') + ' type="checkbox" class="selectall" id="h_' + title
				+ '" /><label class="right" for="h_' + title + '">' + title + '</label><span class="toggler">&#x25B4;</span></div><div>'
				+ html + '</div>';
		$('#db_objects').append(html);
	}
}