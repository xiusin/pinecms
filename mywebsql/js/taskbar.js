/**
 * This file is a part of MyWebSQL package
 *
 * @file:      js/taskbar.js
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

var taskbar = {
	_win: {},
	init: function() {
		$("#taskbar .min-all").button().click( function() {taskbar.minimizeAll(); } );
	},
	openDialog: function (id, url, w, h) {
		var dlg = null;
		id = 'dialog-' + id.replace(/ /g,'-');
		obj = this.findDialog(id);
		if (obj) { // && dlg.status == 1) {
			// confirm current operation abort.. if user cancels, return;
			if (this._win[id].state == 0) {	// restore if minimized, else do nothing
				$('#'+id).dialogExtend("restore");
				this._win[id].state = 1;
			}
			return false;
		} else {	// create a new dialog
			// state: 0 = minimized, 1 = open
			// status: 0: loading, 1 = loaded
			this._win[id] = { status: 0, url: url, state: 1 };
			dlg = this.createDialog(id, url, w, h, false);
		}

		dlg.find('.dialog_contents').attr('src', 'javascript:false');
		this.updateDialog(0, dlg.id);
		dlg.dialog('open');
		dlg.find('ui-dialog-title').html(__('Loading')+'...');
		dlg.find('.dialog_contents').attr("src", url);

		$('#taskbar').append('<input type="button" value="'+__('Loading') + '..." id="tb-button-' + id + '" />');
		$('#tb-button-' + id).button().click( function() { taskbar.handle(id) });
		main_layout.open('south');
	},

	openModal: function (id, url, w, h) {
		id = 'dialog-' + id.replace(/ /g,'-');
		dlg = this.createDialog(id, url, w, h, true);

		dlg.find('.dialog_contents').attr('src', 'javascript:false');
		this.updateDialog(0, dlg.id);
		dlg.dialog('open');
		dlg.find('ui-dialog-title').html(__('Loading')+'...');
		dlg.find('.dialog_contents').attr("src", url);
	},

	handle: function(id) {
		if (this._win[id].state == 0) {
			$('#'+id).dialogExtend("restore");
			this._win[id].state = 1;
		}
		/*else {
			$('#'+id).dialogExtend("minimize");
			this._win[id].state = 0;
		}*/
		$('#'+id).dialog("moveToTop");
	},

	findDialog: function(id) {
		for(dlg in this._win) {
			if (dlg == id)
				return this._win[dlg];
		}
		return false;
	},

	createDialog: function(id, url, w, h, modal) {
		var dlg = $("#dialog-template").clone();
		var dlg_id = id.replace(/ /g,'-');
		dlg.attr('id', dlg_id);
		dlg.find('.dialog_contents').attr('id', dlg_id + '-contents');
		dlg.dialog({
			modal: modal,
			autoOpen: false,
			width: w,
			height: h,
			minWidth: 460,
			minHeight: 260,
			open: function() {
				w = $('#' + dlg_id).parent('.ui-dialog').width();
				h = $('#' + dlg_id).parent('.ui-dialog').height();
				$('#' + dlg_id + '-contents').width(w).height(h);
			},
			close: function() {
				$('#taskbar').find('#tb-button-' + id).remove();
				if ($('#taskbar').find('input').length == 0)
					main_layout.close('south');
				$('#'+dlg_id).dialog('destory');
				$('#'+dlg_id + '-contents').remove();
				$('#'+dlg_id).remove();
				delete taskbar._win[dlg_id];
			}
		});

		if (!modal) {
			dlg.dialogExtend({
				"maximize" : false,
				"minimize" : true,
				"events" : {
					//"maximize" : function(evt, dlg){ alert(evt.type+"."+evt.handleObj.namespace); },
					"minimize" : function(evt, dlg){ taskbar.minimize(dlg.id); }
					//"restore" : function(evt, dlg){ }
				}
			});
		}

		dlg.bind('dialogresizestart dialogdragstart', function() {
			iframe = $('#' + dlg_id + '-contents');
			var d = $('<div></div>');
			$('#' + dlg_id).append(d[0]);
			d[0].id = dlg_id + '-div';
			d.css({position:'absolute'});
			d.css({top:0,left:0});
			d.height(iframe.height());
			d.width('100%');
		});
		dlg.bind('dialogresizestop dialogdragstop', function() {
			$('#' + dlg_id + '-div').remove();
			w = $('#' + dlg_id).parent('.ui-dialog').width();
			h = $('#' + dlg_id).parent('.ui-dialog').height();
			$('#' + dlg_id + '-contents').width(w).height(h);
		});

		$('#' + dlg_id + '-contents').bind('load', function() { taskbar.updateDialog(1, dlg_id); });

		return dlg;
	},

	updateDialog: function(n, id) {
		if (!id)
			return;

		if (n == 1) {
			$('#' + id + ' .dialog_msg').css("display", "none");
			$('#' + id + '-contents').css("display", "block");
			$('#' + id).parent('.ui-dialog').trigger('resize');

			try { // if window has external url, exception will be thrown
				title = document.getElementById(id + '-contents').contentWindow.title;
				if (this._win[id]) {
					this._win[id].status = 1;
					$('#tb-button-' + id).button({ label: title });
				}
				$('#'+id).siblings('.ui-dialog-titlebar').find('.ui-dialog-title').html(title);
				win = document.getElementById(id + '-contents').contentWindow;
				$(win.document).find('#popup_overlay').addClass('ui-helper-hidden');
			} catch(e) {
			}
		} else {
			$('#' + id + ' .dialog_msg').css("display", "block");
			$('#' + id + '-contents').css("display", "none");
		}
	},

	minimize: function(id) {
		this._win[id].state = 0;
	},

	minimizeAll: function() {
		for(dlg in this._win) {
			if (this._win[dlg].state == 1) {
				$('#'+dlg).dialogExtend("minimize");
				this._win[dlg].state = 0;
			}
		}
	}
};

