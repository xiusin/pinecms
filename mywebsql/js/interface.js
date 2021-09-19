/**
 * This file is a part of MyWebSQL package
 *
 * @file:      js/interface.js
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

$.mywebsql = { "popup":false, "dialogs":[] };
var main_layout;
var data_layout;

var currentTreeItem = null;
var clipboard_helper = null;
var historyCurItem = null;

$(document).ready(function () {
	if ($.browser.msie && $.browser.version <= 7) {
		$('select').hide();
		$('#screen-wait').find('div:first').remove();
		$('#screen-wait').find('div.compat-notice').show();
		return true;
	}

	$('body').bind('contextmenu', function(e) {
		var otarget = e.originalTarget || e.target;
		// allow input areas to show default context menu
		if (otarget.type == "textarea" || otarget.type == "text")
			return true;
		return false;
	});
	$("ul#main-menu li ul li:has(ul)").find("a:first").append(" &raquo; ");
	$("ul#main-menu").find("a").click(function(event) {
		event.stopPropagation();
		$(this).blur();
	});
	$('ul#main-menu li').click(function() { eval($(this).find('a').attr('href')); });
	$('#toolbarHolder').mouseover(function() { if (_contextMenu) _contextMenu.hide(); });

	// make tabs first to avoid resize problems
	$(".ui-layout-data-center").tabs();
	$(".ui-layout-data-south").tabs({
		select: function(event, ui) { setTimeout(focusEditor, 200); }
	});

	main_layout_props = {
		spacing_open: 3, spacing_closed: 6
		,east__minSize: 200
		,west__minSize: 150, west__size: $('body').innerWidth() * 0.2
		,north__minSize: 76
		,north__resizable: false, north__closable: true, north__spacing_open: 0, north__spacing_closed: 6
		,north__onopen: function() { main_layout.allowOverflow('north'); }
		,south__resizable: false, south__closable: true, south__spacing_open: 0, south__spacing_closed: 0
		,center__onresize: function() { data_layout.resizeAll(); layoutState.save('main_layout'); }
		,enableCursorHotkey: false
	};

	main_layout = $('body').layout( $.extend(main_layout_props, layoutState.load('main_layout')) );

	data_layout_props = {
		spacing_open: 3, spacing_closed: 6
		,resizerClass: "ui-data-resizer"
		,togglerClass: "ui-data-toggler"
		,center__paneSelector: ".ui-layout-data-center"
		,south__paneSelector: ".ui-layout-data-south"
		,center__resizable: true
		,south__resizable: true, south__closable: true, south__minSize: 66, south__size: 160
		,center__onresize: function() { layoutState.save('data_layout'); }
		,enableCursorHotkey: false
		,onresizeall_end: function() {
			var n = $(".ui-layout-data-center").tabs('option', 'selected');
			if (n == 0) resizeTableHeader('data');
			else if (n == 2) resizeTableHeader('info'); 
		}
	};

	data_layout = $('div.ui-layout-center').layout( $.extend(data_layout_props, layoutState.load('data_layout')) );

	main_layout.close('south');
	main_layout.allowOverflow('north');

	$(".ui-layout-data-center").tabs('select', 2);

	$("#tablelist").treeview();

	contextHandler();

	// chrome selection issue fix when a resizer is used for resizing
	$('.ui-layout-resizer').bind('selectstart', function() { return false; });

	$('#nav_query').button({
		text: true,	icons: {primary: 'ui-icon-play'}
	}).click(function() { queryGo(0); });
	$('#nav_queryall').button({
		text: true,	icons: {primary: 'ui-icon-seek-end'}
	}).click(function() { queryGo(1); });
	$('#nav_delete').button({
		text: true,	icons: {primary: 'ui-icon-close'}
	}).click(function() { queryDelete(); });
	$('#nav_update').button({
		text: true,	icons: {primary: 'ui-icon-disk'}
	}).click(function() { querySave(); });
	$('#nav_gensql').button({
		text: true,	icons: {primary: 'ui-icon-script'}
	}).click(function() { queryGenerate(); });
	$('#nav_addrec').button({
		text: true,	icons: {primary: 'ui-icon-plusthick'}
	}).click(function() { queryAddRecord(); });
	$('#nav_copyrec').button({
		text: true,	icons: {primary: 'ui-icon-copy'}
	}).click(function() { queryCopyRecord(); });
	$('#nav_refresh').button({
		text: true,	icons: {}
	}).click(function() { queryRefresh(); });

	$('#sp-results-maximize').button({
		text: false, icons: {primary: 'ui-icon-newwin'}
	}).click(function() { resultsPaneToggle(); });

	initClipboard();

	$(window).unload(function(){
		layoutState.save('data_layout');
		layoutState.save('main_layout');
	});

	taskbar.init();
	$("#object-filter-text").quickText().bind('keyup', function() {
		// filter database list or table list when quick search filter is applied
		var li = $("#tablelist").hasClass("dblist") ? "span.odb a" : "span.file a";
		$("#object_list").setObjectFilter( $(this).val(), li, 'ul' );
	});

	$('#screen-wait').remove();
	$('#wrkfrm').attr('src', 'index.php?q=wrkfrm&type=info');

	loadUserPreferences();
	showNavBtns('query', 'queryall');
});

function contextHandler() {
	// setup context menus for everything first time
	if(arguments.length == 0) {
		// remove options that do not apply to this type of database (server)
		$('ul#main-menu .option').not("."+DB_DRIVER).remove();
		$('#object-context-menus .option').not("."+DB_DRIVER).remove();

		$(".ui-layout-north").contextMenu('#panel-header');
		$("#object_list").contextMenu('#panel-menu-objects');
		$("#sql-editor-pane").contextMenu('#panel-menu-editor');
		$("#sql-history").contextMenu('#history-menu');
	}

	// only update context menus for object list
	$('#tablelist .odb').contextMenu('#db-menu');
	$('#tablelist .otable').contextMenu('#table-menu');
	$('#tablelist .oview').contextMenu('#view-menu');
	$('#tablelist .oproc').contextMenu('#proc-menu');
	$('#tablelist .ofunc').contextMenu('#func-menu');
	$('#tablelist .otrig').contextMenu('#trig-menu');
	$('#tablelist .oevt').contextMenu('#evt-menu');
	$('#tablelist .schmf').contextMenu('#schm-menu');
	$('#tablelist span').filter('.tablef,.viewf,.procf,.funcf,.trigf,.evtf').contextMenu('#object-menu');
}

function initClipboard() {
	$('#sql-history tr').live('hover', function() { historyCurItem = $(this); } );

	// copying using zeroclipboard and context menu is a pain... but we have to do it ...
	ZeroClipboard.setMoviePath('js/jquery.clipboard.swf');
	clipboard_helper = new ZeroClipboard.Client();
	$('#history-menu li.clipboard').mouseover(function() {
		single = $(this).hasClass('single');  // copy single or all queries
		clipboard_helper.setText(getHistoryText(single));
		if (clipboard_helper.div) {
			clipboard_helper.receiveEvent('mouseout', null);
			clipboard_helper.reposition(this);
		}
		else {
			clipboard_helper.glue(this);
			$(clipboard_helper.div).click(function() {
				clipboard_helper.hide();
			});
		}
		clipboard_helper.receiveEvent('mouseover', null);
	});
}

function showNavBtns() {
	bn = new Array("addrec", "copyrec", "query", "queryall", "delete", "update", "gensql");
	for(i=0; i<bn.length; i++)
		$('#nav_' + bn[i]).css("display", "none");

	for(i=0; i<arguments.length; i++)
		$('#nav_' + arguments[i]).css("display", "block");
}

function showNavBtn() {
	for(i=0; i<arguments.length; i++)
		$('#nav_' + arguments[i]).css("display", "block");
}

function hideNavBtn(btn) {
	for(i=0; i<arguments.length; i++)
		$('#nav_' + arguments[i]).css("display", "none");
}

function switchEditor(n) {
	$(".ui-layout-data-south").tabs('select', n);
}

function editorTextSize(s) {
	ts = parseFloat($(currentEditor().editor.container).css('font-size'));
	if (ts) {
		if (s > 0 && ts <= 100)
			ts += s;
		else if (s < 0 && ts >= 11)
			ts += s;
		$(currentEditor().editor.container).css('font-size', ts + 'px');
	}
}

function editorClear() {
	$(currentEditor().editor.container).html('');
}

function currentEditor() {
	n = $(".ui-layout-data-south").tabs('option', 'selected');
	obj = commandEditor;
	switch(n) {
		case 1: obj = commandEditor2; break;
		case 2: obj = commandEditor3; break;
	}
	return obj;
}

function focusEditor() {
	ed = currentEditor();
	ed.focus();
}

function getDataMenu(m, t, e) {
	target = $(e.originalTarget || e.target);

	if (target.hasClass('tch'))
		return false;

	if (target.is('th'))
		return $('#data-menu-th').clone();
	else if (target.is('td'))
		return $('#data-menu-td').clone();

	return false; // no menu here
}

function objListHandler(data, state) {
	tree = $(data).find('#objlist').html();
	if (tree != '') {
		$('#object_list').html(tree);
		$("#tablelist").treeview();
		contextHandler(false);
		$("#object-filter-text").val("");
	}
	else
		jAlert(__('An error occured while refreshing the object list.'));

	// restore previous tree state
	for(i=0;i<state.length;i++)
		$('#'+state[i]+' span:first').trigger('click');

	setPageStatus(false);
}

function resultsPaneToggle() {
	var btn = $("#sp-results-maximize");
	var max = btn.data("max");
	if (max == 1) {
		main_layout.open('north');
		data_layout.open('south');
		main_layout.open('west');
		btn.removeData("max");
	} else {
		btn.data("max", 1);
		main_layout.close('north');
		data_layout.close('south');
		main_layout.close('west');
	}
}
function uiCreateDialog(id) {
	// dialog is already created for selected element
	if ($.inArray(id, $.mywebsql.dialogs) != -1)
		return true;
	if (id == 'dbcreate') {
		$("#dialog-dbcreate").dialog({
			autoOpen: false,
			width:320,
			height: 160,
			modal: true,
			buttons: [ {
						text: __('Cancel'),
						click: function() { $(this).dialog('close'); }
					}, {
						text: __('Create Database'),
						click: function() { dbCreate(1); }
					}
			]
		});
	}
	$.mywebsql.dialogs.push(id);
	return true;
}

function initEditor(n) {
	var editor = commandEditor;
	var ck = "sql_commandEditor";
	switch(n) {
		case 1: {
			editor = commandEditor2;
			ck = "sql_commandEditor2";
		} break;
		case 2: {
			editor = commandEditor3;
			ck = "sql_commandEditor3";
		} break;
	}

	var x = $.cookies.get(ck);
	if ( x )
		editor.setCode( x );
}

// quick search filter functionality for dom elements other than tables
$.fn.setObjectFilter = function(text, elem, container) {
	if (text == '') {
		$(elem, this).parentsUntil(container).removeClass('ui-helper-hidden');
		return this;
	}

	string = text.toUpperCase();
	$(elem, this).each(function(){
		var contents = $(this).text().toUpperCase();
		// check the string against that element text
		if ( contents.match(string) ) {
			$(this).parentsUntil(container).removeClass('ui-helper-hidden')
		} else {
			$(this).parentsUntil(container).addClass('ui-helper-hidden');
		}
	});
};

$.fn.quickText = function() {
	return this.each(function(s) {
		input = $(this);
		if (input.val() == '')
		{
			input.addClass('blur');
			input.val(input.attr('data-placeholder'));
		}
		input.focus(function()
		{
			if ( $.trim($(this).val()) ===  $(this).attr('data-placeholder') ) {
				$(this).val("").removeClass('blur');
			}
		}).blur(function(){
			if ( $.trim($(this).val()) === "" ) {
				$(this).addClass('blur').val($(this).attr('data-placeholder'));
			}
		});
		return $(this);
	});
}
