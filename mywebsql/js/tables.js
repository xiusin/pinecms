/**
 * This file is a part of MyWebSQL package
 *
 * @file:      js/tables.js
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

var curEditField = null;			// current edited field
var curEditType = null;
var fieldInfo = null;				// information about fields in record set
var editOptions = { sortable:true, highlight:true, selectable:true, editEvent:'dblclick', editFunc:editTableCell };

var selectedRow = -1;
var res_modified = false;			// is the result modified?

var editHorizontal = false;

var editListOpen = false;

// options can include
// highlight: boolean: highlights row on mouse over
// selectable: boolean: makes row selectable
// editable: boolean: makes grid editable
// sortable: boolean: makes table sorting possible using its header

function setupTable(id, opt) {
	res_modified = false;

	if (!opt.editEvent) opt.editEvent = 'dblclick';
	if (!opt.editFunc) opt.editFunc = editTableCell;

	// only sort if there is more than one row
	if (opt.sortable && $('#'+id+' tbody tr').length > 2 ) {
		createTableHeader(id.substr(0,4));

		if (opt.sortable == true) {
			$('#dataHeader thead th').live('click', function() {
				if ($(this).attr('class').match(/tch|th_nosort/)) {
					return true;
				}
				var sort_column = $(this).index();
				if (editTableName && editTableName != '') {
					sort_column--;
				}
				goSort(sort_column);
			});
		} else if (opt.sortable == 'inline') {
			var DATE_RE = /^(\d\d?)[\/\.-](\d\d?)[\/\.-]((\d\d)?\d\d)$/;
			$("#"+id).sorttable();
		}
	}

	if (opt.highlight) {
		$('#'+id+' > tbody > tr').live('mouseenter', function() {
			$(this).addClass("ui-state-hover");
		});
		$('#'+id+' > tbody > tr').live('mouseleave', function() {
			$(this).removeClass('ui-state-hover');
		});
	}

	if (opt.selectable) {
		$('#'+id+' > tbody > tr').live('click', function() {
			if (selectedRow != null)
				$(selectedRow).removeClass('ui-state-active');
			$(this).addClass("ui-state-active");
			selectedRow = this;
		});
	}

	if (opt.editable) {
		editOptions = opt;
		$('#'+id+' td.edit').bind(opt.editEvent, opt.editFunc);
		// for mobile devices
		if(jQuery.support.touch){
			$('#'+id+' td.edit').bind("taphold", opt.editFunc);
		}
		$('#inplace-text textarea').unbind('keydown').bind('keydown', checkEditField);
	}
}

function editTableCell() {
	td = $(this);
	if (curEditField != null)
		closeEditor(true);

	isBlob = td.find('span.i').length;
	isText = td.find('span.d').length;
	txt =  isText ? td.find('span.d').text() : (isBlob ? td.find('span.i').text() : td.text());
	tstyle = td.hasClass('tr') ? "right" : "left";

	td.data('defText', txt);

	curEditField = this;
	index = td.index()-2;
	fi = getFieldInfo(index);
	w = td.width() - (isBlob ? 22 : 0);
	h = td.height();
	td.attr('width', w);

	input = createCellEditor(td, fi, txt, w, h, tstyle);

	if(input) {
		setTimeout( function() {
			input.focus();
			td.ensureVisible($("#results-div"), editHorizontal);
		}, 50 );
	}
}

function closeEditor(upd, value) {
	if (!curEditField)
		return;

	obj = $(curEditField);
	txt = '';
	var xt = new Object();
	if (upd) {
		if (arguments.length > 1 && value == null) {
			xt.value = "NULL";
			xt.setNull = true;
		}
		else {
			txt = xt.value = (curEditType == 'simple') ? obj.find('input').val() : $('#inplace-text textarea').val();
			xt.setNull = false;
		}

		// if not modified, don't bother
		if ( (xt.value != obj.data('defText')) || (xt.setNull && !obj.hasClass("tnl"))
				|| (!xt.setNull && obj.hasClass('tnl')) ) {
			if (!obj.parent().hasClass('n'))
				obj.parent().addClass('x');
			obj.data('edit', xt).addClass('x');
			res_modified = true;

			if (typeof showNavBtn == "function")
				showNavBtn('update', 'gensql');

			if(xt.setNull)
				obj.removeClass('tl').addClass('tnl');
			else
				obj.removeClass('tnl').addClass('tl');

			txt = xt.value;
			//txt = str_replace("<", "&lt;", xt.value);
			//txt = str_replace(">", "&gt;", txt);
		}
	}
	else
		txt = obj.data('defText');

	if (curEditType == 'text') {
		if (xt.setNull)
			obj.find('span.i').text('NULL').removeClass('tl').addClass('tnl');
		else {
			obj.find('span.i')
				.text(txt.length == 0 ? '' :
					( txt.length <= MAX_TEXT_LENGTH_DISPLAY ? txt : 'Text Data [' + formatBytes(txt.length) + ']') )
				.removeClass('tnl');
		}
		obj.find('span.d').text(txt);
	}
	else {
		if (obj.find('span.i').length == 0)
			obj.text(txt);
		else
			obj.find('span.i').text(txt);
	}

	obj.removeAttr('width');
	curEditField = null;

	if (curEditType == 'text')
		$('#inplace-text').hide();
		
	editListOpen = false;
	resizeTableHeader('data');
}

function checkEditField(event) {
	editHorizontal = false;
	// enter, tab, up arrow, down arrow
	keys = (curEditType == 'text') ? [9] : [13,9,38,40];
	if (keys.indexOf(event.keyCode) != -1) {
		event.preventDefault();
		elem = false;
		if (event.keyCode == 9) {
			elem = event.shiftKey ? $(curEditField).prev('.edit') : $(curEditField).next('.edit');
			// move to next/previous record if possible
			if (!elem.length) {
				tr = event.shiftKey ? $(curEditField).parent().prev() : $(curEditField).parent().next();
				if (tr.length)
					elem = event.shiftKey ? tr.find('td.edit:last') : tr.find('td.edit:first');
			}
			editHorizontal = true;
		}
		else if (event.keyCode == 38 || event.keyCode == 40) {
			tr = event.keyCode == 38 ? $(curEditField).parent().prev() : $(curEditField).parent().next();
			if (tr.length)
				elem = tr.find('td').eq($(curEditField).index());
		}
		$('#inplace-text textarea').unbind('blur');
		closeEditor(true);
		if (elem && elem.length)    // edit next or previous element
			elem.trigger(editOptions.editEvent);
	}
	else if (event.keyCode == 27)
		closeEditor(false);
	/*else if ($(this).attr('readonly') != '' && event.keyCode == 32 ) {
	   // focus is on a blob editor, need to open dialog for blob editing
		oldEditField = curEditField;
		closeEditor(false);
		$(oldEditField).find('span.blob').click();
	}*/
}

function createCellEditor(td, fi, txt, w, h, align) {
	curEditType = 'simple';
	keyEvent = 'keydown';
	input = null;
	code = '<form name="cell_editor_form" class="cell_editor_form" action="javascript:void(0);">';
	if (fi['blob'] == 1) {
		if (fi['type'] == 'binary') {
			code += '<input type="text" readonly="readonly" name="cell_editor" class="cell_editor" style="text-align:' + align + ';width: ' + (w-2) + 'px;" />';
			code += '</form>';
			td.find('span.i').html(code);
			input = td.find('input');
			input.val(txt).bind(keyEvent, checkEditField ).blur( function() { closeEditor(true); } );
		}
		else {
			span = $(td).find('span.d');
			txt = span.text();

			if ( optionsGet('ui-edit-popup') ) {
				// show dialog text editor to make large text editing easier
				$("#dialog-text-editor").dialog({
					autoOpen: true,
					width: 500,
					height: 300,
					modal: true,
					open: function() {
						$("#text-editor").val( txt );
						// suspend hotkeys while the popup editor is active
						window.hotkeys.suspend(true);
					},
					buttons: [ {
								text: __('Cancel'),
								click: function() { window.hotkeys.suspend(false); $(this).dialog('close'); }
							}, {
								text: __('Save'),
								click: function() {
									txt = $("#text-editor").val();
									span.text( txt );
									td.find('span.i').text(txt.length == 0 ? '' :
										( txt.length <= MAX_TEXT_LENGTH_DISPLAY ? txt : 'Text Data [' + formatBytes(txt.length) + ']') )
										.removeClass('tnl');
									var xt = new Object();
									xt.value = txt;
									xt.setNull = false;
									if (!td.parent().hasClass('n'))
										td.parent().addClass('x');
									td.data('edit', xt).addClass('x');
									if (typeof showNavBtn == "function")
										showNavBtn('update', 'gensql');
									curEditField = null;
									window.hotkeys.suspend(false);
									$(this).dialog('close');
								}
							}
					]
				});
				input = null;
			} else {
				w = td.width()-20;
				if (w < 200) w = 200;
				textarea = $('#inplace-text textarea');
				textarea.width(w).val(txt);

				$('#inplace-text').show().position({ of: td, my: "left top", at: "left top", offset: 0 });
				$('#inplace-text textarea').blur( function() { closeEditor(true); } );
				curEditType = 'text';
				input = textarea;
			}
		}
	}
	else {
		switch(fi['type']) {
			default:
				code += '<input type="text" name="cell_editor" class="cell_editor" style="text-align:' + align + ';width: ' + (w-2) + 'px;" />';
				if( fi['list'] && fi['list'].length > 0 ) {
					code += '<a href="javascript:void(0)" class="cell_editlist">&#x25BE;</a>';
				} else if ( fi['type']  && ( fi['type'] == "datetime" || fi['type'] == "date" ) ) {
					code += '<div class="dp"></div><a href="javascript:void(0)" class="cell_editlist">&#x25BE;</a>;';
				}
				code += '</form>';
				td.html(code);
				input = td.find('input');
				input.val(txt).select().bind(keyEvent, checkEditField ).blur( function(e) { if (!editListOpen) closeEditor(true); } );
				if( fi['list'] && fi['list'].length > 0 ) {
					$(".cell_editor").css({width:(w-20)+'px'}).autocomplete({
						minLength: 0,
						source: fi['list'],
						open: function( event, ui ) { $(".cell_editor").autocomplete( "widget" ).css({width:(w>160?w-2:160)+"px"}); },
						close: function( event, ui ) { $(".cell_editor").focus(); editListOpen = false; }
					});
					$(".cell_editlist").mousedown(function(e) {
						e.preventDefault();
						editListOpen = true;
						$(".cell_editor").focus().autocomplete("search", "");
						return false;
					});
				} else if ( fi['type']  && ( fi['type'] == "datetime" || fi['type'] == "date" ) ) {
					$(".cell_editor").css({width:(w-20)+'px'});
					$(".dp").datepicker({
						dateFormat:"yy-mm-dd",
						changeMonth: true,
						changeYear: true,
						onSelect: function(d, o) {
							$(".dp").fadeOut(300);
							d += $(".cell_editor").data("datetime");
							$(".cell_editor").val(d).focus();
							editListOpen = false;
						}
					});
					$(".cell_editlist").mousedown(function(e) {
						e.preventDefault();
						editListOpen = true;
						var d = $(".cell_editor").val();
						$(".cell_editor").data("datetime", d.substr(10));
						$(".dp").datepicker("setDate", d.substr(0, 10));
						$(".dp").fadeIn(300);
						$(".cell_editor").focus();
						return false;
					});
				}
				break;
		}
	}
	return input;
}

function createTableHeader(name) {
	var table = "#" + name + "Table";
	if ($(table + ' tbody tr').length <= 2)
		return;

	var header = "#" + name + "Header";
	var div = name == "data" ? '#results-div' : "#info-div";
	$(header).remove(); // just in case we have it created and not yet destroyed
	var tableHeader = $(table).clone();
	tableHeader.find('tbody').remove();
	tableHeader.attr('id', name + "Header").appendTo(div);

	$(div).scroll(function () {
		var t = parseInt($(this).scrollTop());
		tableHeader.css({top: t + 'px'})
	});
	
}

function resizeTableHeader(name) {
	var table = "#" + name + "Table";
	if ($(table + ' tbody tr').length <= 2)
		return;
		
	var header = "#" + name + "Header";
	var tableHeader = $(header);
	tableHeader.width($(table).width());
	var ths = $(table + " thead th");
	var l = ths.length;
	for (i = 0; i < l; i++) {
		var w = $(ths[i]).width();
		tableHeader.find("thead th").eq(i).width(w);
	}
}

$.fn.ensureVisible = function(el, horiz) {
	if (horiz) {
		pl = el.prop("scrollLeft");
		pw = el.width();
		p = this.position();
		w = this.width();
		if( pw < (p.left+w)  )
			el.prop("scrollLeft", p.left + w);
		else if( p.left < 0 )
			el.prop("scrollLeft", p.left);
	} else {
		pt = el.prop("scrollTop");
		ph = el.height();
		p = this.position();
		h = this.height();
		if( ph < (p.top+h) )
			el.prop("scrollTop", p.top + h);
		else if( p.top < 0 )
			el.prop("scrollTop", p.top);
	}
};

// quick table search filter functionality
$.fn.setSearchFilter = function(text) {
	if (text == '')
		$('tr', this).removeClass('ui-helper-hidden');
	else {
		string = text.toUpperCase();
		$('tbody tr', this).each(function(){
			var found = false;
			$('td', this).each(function(){
				var contents = $(this).text().toUpperCase();
				// check the string against that cell
				if ( contents.match(string) ) {
					found = true;
					return true;
				}
			});

			if (found)
				$(this).removeClass('ui-helper-hidden')
			else
				$(this).addClass('ui-helper-hidden');
		});
	}
};

/* Quick inline sorttable
	based on the 'stupid-table' jQuery plugin by joequery
*/

(function($) {

  $.fn.sorttable = function(sortFns) {
    return this.each(function() {
      var $table = $(this);
		var $header = $ ("#" + $table.attr('id').replace("Table", "Header"));
      sortFns = sortFns || {};

      // Merge sort functions with some default sort functions.
      sortFns = $.extend({}, $.fn.sorttable.default_sort_fns, sortFns);

      $header.on("click.sorttable", "thead th", function() {
        var $this = $(this);
        var th_index = 0;
        var dir = $.fn.sorttable.dir;

        // Account for colspans
        $this.parents("tr").find("th").slice(0, $this.index()).each(function() {
          var cols = $(this).attr("colspan") || 1;
          th_index += parseInt(cols,10);
        });

        // Determine (and/or reverse) sorting direction, default `asc`
        var sort_dir = $this.data("sort-default") || dir.ASC;
        if ($this.data("sort-dir"))
           sort_dir = $this.data("sort-dir") === dir.ASC ? dir.DESC : dir.ASC;

        // Choose appropriate sorting function.
        var type = $this.data("sort") || null;

        // Prevent sorting if no type defined
        if (type === null) {
          return;
        }

        // Trigger `beforetablesort` event that calling scripts can hook into;
        // pass parameters for sorted column index and sorting direction
        //$table.trigger("beforetablesort", {column: th_index, direction: sort_dir});
        // More reliable method of forcing a redraw
        $table.css("display");

        // Run sorting asynchronously on a timout to force browser redraw after
        // `beforetablesort` callback. Also avoids locking up the browser too much.
        setTimeout(function() {
          // Gather the elements for this column
          var column = [];
          var sortMethod = sortFns[type];
          var trs = $table.children("tbody").children("tr");

          // Extract the data for the column that needs to be sorted and pair it up
          // with the TR itself into a tuple
          trs.each(function(index,tr) {
            var $e = $(tr).children().eq(th_index);
            var sort_val = $e.data("sort-value");
            var order_by = typeof(sort_val) !== "undefined" ? sort_val : $e.text();
            column.push([order_by, tr]);
          });

          // Sort by the data-order-by value
          column.sort(function(a, b) { return sortMethod(a[0], b[0]); });
          if (sort_dir != dir.ASC)
            column.reverse();

          // Replace the content of tbody with the sorted rows. Strangely (and
          // conveniently!) enough, .append accomplishes this for us.
          trs = $.map(column, function(kv) { return kv[1]; });
          $table.children("tbody").append(trs);

          // Reset siblings
          $table.find("th").data("sort-dir", null).removeClass("sorting-desc sorting-asc");
          $this.data("sort-dir", sort_dir).addClass("sorting-"+sort_dir);
			 $this.parent().find("span").remove();
			 $table.find("thead span").remove();
			 $this.find("div").append( sort_dir == dir.ASC ? "<span>&nbsp;&#x25B4;<span>" : "<span>&nbsp;&#x25BE;<span>");
			 $table.find("thead th:eq("+th_index+") div").append( sort_dir == dir.ASC ? "<span>&nbsp;&#x25B4;<span>" : "<span>&nbsp;&#x25BE;<span>");
			 resizeTableHeader('info');

			 // reset first index numbering
			 $("tbody tr", $table).each(function(i) { $("td:eq(0)", this).html(i+1); });
          // Trigger `aftertablesort` event. Similar to `beforetablesort`
          //$table.trigger("aftertablesort", {column: th_index, direction: sort_dir});
          // More reliable method of forcing a redraw
          $table.css("display");
        }, 10);
      });
    });
  };

  // Enum containing sorting directions
  $.fn.sorttable.dir = {ASC: "asc", DESC: "desc"};

  $.fn.sorttable.default_sort_fns = {
    "numeric": function(a, b) {
		if (a != "NULL" && b != "NULL")
			return parseInt(a, 10) - parseInt(b, 10);
		if (a == "NULL")
			return -1;
		return 1;
    },
    "float": function(a, b) {
		if (a != "NULL" && b != "NULL")
			return parseFloat(a) - parseFloat(b);
		if (a == "NULL")
			return -1;
		return 1;
    },
    "text": function(a, b) {
		if (a != "NULL" && b != "NULL")
			return a.localeCompare(b);
		if (a == "NULL")
			return -1;
		return 1;
    },
    "text-ins": function(a, b) {
      a = a.toLocaleLowerCase();
      b = b.toLocaleLowerCase();
      return a.localeCompare(b);
    }
  };

})(jQuery);
