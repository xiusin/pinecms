/*
 * EmbeddedHelp v1.1.4 - jQuery help plugin
 * file: jquery.ehelp-1.1.4.js
 * Copyright (c) 2011 Josip Kalebic
 * josip.kalebic@gmail.com, www.embedded-help.net
 * improved by ckszabi@gmail.com
 *
 * Dual licensed under the MIT and GPL licenses:
 * 	http://www.opensource.org/licenses/mit-license.php
 * 	http://www.gnu.org/licenses/gpl.html
 *
 * Paths definition:
	[{
	'rel':'rel_name',
	'method':'static or animated',
	'path': [{
		'element': 'jquery selector',
		'desc': 'Help content (tooltip)',
		'duration': number in ms,
		"align": position of tooltip,
		'marker': 'class to mark element on page',
		'extf' : external function
		'ftriger': call external function at beginning or end 'B' or 'E'
		'value': value used by external function 
	}]
	}];
//---------------------------------------------------------
		extf: calls external function
		      two methods:
			- javascript function
				- examp: "diferentWay('#search','My phrase')"
					  whole string going throu $.globalEval
			- jquery function
				- examp: "userValueF"
					 - parameters are send in form of object
					 - whole path elements are send into function
//---------------------------------------------------------
		'ftriger': call external function at beginning or end 
			- options: 'B' or 'E'
			- default: 'B'
			- used for animated method only
//---------------------------------------------------------
		 align: the position of the tip. Possible values are
					LT  left top
					LB  left bottom
					RT  right top
					RB  right bottom
					R   right
					L   left
//---------------------------------------------------------
	OPTIONS: default values
		'animatedvp': true, -> viewport traction on animated method
		'staticvp': true,  -> viewport traction on static method
		'autoalign':true  -> align of tooltips 
					true -> positioned by script, 
					false -> poistioned by "align" value in paths 	              definition
		'autolinks':true  -> generates help links from json paths definitions
		//---------------------------------------------
*/
(function($){  
    $.fn.extend({   
        embeddedHelp: function(pathdefinition, options) {  

            var pathdefinition;

            var activeStaticPath;

            var Ghost;

            var ViewportOnMove;

            var Htimer;

            var options = $.extend({
                'animatedvp': true,
                'staticvp': true,
                'autoalign':true,
                'callextf':true,
					 'autolinks':true
            }, options);
            //--------------------------------------------------
            function getViewportSize() {
                var mode, domObject, size = {
                    height: window.innerHeight,
                    width: window.innerWidth
                };
		
                // if this is correct then return it. iPad has compat Mode, so will
                // go into check clientHeight/clientWidth (which has the wrong value).
                if (!size.height) {
                    mode = document.compatMode;
                    if (mode || !$.support.boxModel) { // IE, Gecko
                        domObject = mode == 'CSS1Compat' ?
                        document.documentElement : // Standards
                        document.body; // Quirks
                        size = {
                            height: domObject.clientHeight,
                            width:  domObject.clientWidth
                        };
                    }
                }
		
                return size;
            }
            //--------------------------------------------------
            function getViewportOffset() {
                return {
                    top:  window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop,
                    left: window.pageXOffset || document.documentElement.scrollLeft || document.body.scrollLeft
                };
            }
            //--------------------------------------------------
            function getDocumentSize() {
                return {
                    height: $(document).height(),
                    width:  $(document).width()
                };
            }
            //--------------------------------------------------
            function getElementSize($element) {
                return {
                    height: $element.height(),
                    width:  $element.width()
                };
            }
            //----------------------------------------------------
            function getAutoAlign(leftx, topy) {

                var doffset = getDocumentSize();

                var dhalfheight = (doffset.height / 2);
                var dhalfwidth = (doffset.width / 2);

                if(leftx > dhalfwidth) {
                    var algl = 'L';
                } else {
                    var algl = 'R';
                }

                if(topy > dhalfheight) {
                    var algt = 'T';
                } else {
                    var algt = 'B';
                }

                return (algl + algt);

            }
            //-----------------------------------
            function checkAllStaticBoxes() {

                $('.EHtoolgost').remove();
                $('.EHtooltipc').each(function(index) {
                    setView($(this));
                });
                return false;
            }
            //-----------------------------------
            function checkPointer() {
                doSetViewport($('#EHpointer'));
                return false;
            }
            //-----------------------------------------
            function compareOtherGhosts(bnumber, btop, bleft, bheight, bwidth) {

                var difftop = 0;
                var diffleft = 0;

                var vpSize = getViewportSize();
                var vpOffset = getViewportOffset();
                var wtop = vpOffset.top;
                var wleft = vpOffset.left;
                var wbottom = wtop + vpSize.height;
                var wright = wleft + vpSize.width;

                var conflict = false;
                var conflictnumbers = "";

                $('.EHtoolgost').each(function(index) {
                    var toffset = $(this).offset();
                    var twidth = $(this).outerWidth();
                    var theight = $(this).outerHeight();
                    var tnumber = $(this).find('span').html();

                    if(tnumber != bnumber && tnumber != "") {
			
                        if((bleft >= toffset.left && bleft <= (toffset.left + twidth)) || ((bleft + bwidth) >= toffset.left && (bleft + bwidth) <= (toffset.left + twidth))) {

                            if((btop >= toffset.top && btop <= (toffset.top + theight)) || ((btop + bheight) >= toffset.top && (btop + bheight) <= (toffset.top + theight))) {

                                $(this).find('span').html(tnumber + ", " + bnumber);
                                conflict = true;
                            }

	
                        }

                    }

                });


                return conflict;

            }
            //-----------------------------------
            function setView(that) {

                var poffset = that.offset();

                var pwidth = that.outerWidth();
                var pheight = that.outerHeight();

                var dSize = getDocumentSize();

                var vpSize = getViewportSize();
                var vpOffset = getViewportOffset();
 
                var wtop = vpOffset.top;
                var wleft = vpOffset.left;

                var wbottom = wtop + vpSize.height;
                var wright = wleft + vpSize.width;


                var number = that.find('span').html();
                var boxtext = "";

                var doGhost = false;

                //check if bottom scrollbar appears
                if(dSize.width > vpSize.width) {
                    var bootom_margin = 50;
                    var horizontalscrollbar = true;
                } else {
                    var bootom_margin = 35;
                    var horizontalscrollbar = false;
                }

                //check if right scrollbar appears
                if(dSize.height > vpSize.height) {
                    var right_margin = 50;
                    var verticalscrollbar = true;
                } else {
                    var right_margin = 35;
                    var verticalscrollbar = false;
                }

                if(poffset.top > wbottom) {
                    var newtop = (wbottom - bootom_margin);
                    doGhost = true;
                } else if(horizontalscrollbar == true && poffset.top > (wbottom - 20)) {
                    var newtop = (wbottom - bootom_margin);
                    doGhost = true;
                } else if((poffset.top + pheight) < wtop) {
                    var newtop = (wtop);
                    doGhost = true;
                } else {
                    var newtop = poffset.top;
                }

                if((poffset.left + pwidth) < wleft) {
                    var newleft = (wleft + 10);
                    doGhost = true;
                } else if(verticalscrollbar == true && (poffset.left + pwidth) < (wleft - 20)) {
                    var newleft = (wleft + 10);
                    doGhost = true;
                } else if(poffset.left > wright) {
                    var newleft = (wright - right_margin);
                    doGhost = true;
                } else {
                    var newleft = poffset.left;
                }

                if(doGhost == true) {


                    var GhostBox = $("<div>").addClass("EHtoolgost").html("<span>" + number + "</span>").appendTo("body");

                    var conflict = compareOtherGhosts(number, newtop, newleft, GhostBox.outerHeight(), GhostBox.outerWidth());

                    if(conflict != true) {

                        GhostBox.attr('id', 'ghost_' + number).css("top",(newtop + 5) + "px").css("left", (newleft - 5) + "px").appendTo("body").fadeIn("slow");

                    }
			
                }
		
                return false;
            }
            //----------------------------------------------------------------
            function doSetViewport(that) {

                var poffset = that.offset();

                var poHeight = that.height();
                var poWidth = that.width();

                var vpSize = getViewportSize();
                var vpOffset = getViewportOffset();
 
                var wtop = vpOffset.top;
                var wleft = vpOffset.left;

                var wbottom = wtop + vpSize.height;
                var wright = wleft + vpSize.width;

                var dSize = getDocumentSize();

                var doAnimateTop = false;
                var doAnimateLeft = false;

                if($.browser.opera) {
                    var movetag = "html";
                }
                else {
                    var movetag = "body,html";
                }

                if(poWidth > 0) {

                    if((poffset.top + poHeight) > wbottom) {
                        var newtop = (wtop + (vpSize.height));
                        if(newtop + vpSize.height > dSize.height) {
                            newtop = dSize.height - vpSize.height;
                        }
                        if(wtop != newtop) {
                            doAnimateTop = true;
                        }

                    } else if(poffset.top < wtop) {
                        var newtop = (wtop - (vpSize.height));
                        if(newtop < 0) {
                            newtop=0;
                        }
                        if(wtop != newtop) {
                            doAnimateTop = true;
                        }
                    }

                    if(poffset.left < wleft) {
                        var newleft = (wleft - (vpSize.width));
                        if(wleft < 0) {
                            wleft=0;
                        }
                        if(wleft != newleft) {
                            doAnimateLeft = true;
                        }
                    } else if((poffset.left + poWidth) > wright) {
                        var newleft = (wleft + (vpSize.width));
                        if(newleft + vpSize.width > dSize.width) {
                            newleft = dSize.width - vpSize.width;
                        }
                        if(wleft != newleft) {
                            doAnimateLeft = true;
                        }
                    }

                    if(doAnimateTop == true && ViewportOnMove != true) {
                        ViewportOnMove = true;

                        $(movetag).animate({
                            scrollTop : newtop
                        },'slow', function(event) {
                            ViewportOnMove = false;
                        });
                    }

                    if(doAnimateLeft == true && ViewportOnMove != true) {
                        ViewportOnMove = true;
                        $(movetag).animate({
                            scrollLeft : newleft
                        },'slow', function(event) {
                            ViewportOnMove = false;
                        });
                    }
                }
	
                return false;
	
            }
            //-----------------------------------
            function callExtFunction(elobj) {


                if(options.callextf == true && elobj.extf!="" && elobj.extf!=undefined) {

                    if($[elobj.extf]) {
                        var argsObj = elobj;
                        argsObj.object = $(elobj.element);
                        var ret = $[elobj.extf](argsObj);
                    } else {
                        var ret = jQuery.globalEval(elobj.extf);
                    }

                }

                return false;

            }
            //-----------------------------------
            function doClearAll() {

                clearInterval(Htimer);

                activeStaticPath = null;

                $("#EHpointer").stop(true);
                $("#EHtooltip").stop();

                $(".EHtooltmp").remove();
                $("#EHtooltip").remove();
                $("#EHpointer").remove();
                $(".EHtooltipc").remove();
                $(".EHtoolgost").remove();

                if(!$.isEmptyObject(pathdefinition)) {
                    $.each(pathdefinition, function(reli, pathvalue) {
                        $.each(pathvalue.path, function(key, value) {
                            $(value.element).removeClass(value.marker);
                        });
                    });
                }

                return false;
            }
            //-----------------------------------
            function doPathAnimation(that, pathvalue){

                var aoff = that.offset();

                var dSize = getDocumentSize();

                var elmax = pathvalue.length;

                if(!$.isEmptyObject(pathvalue)) {

                    $("body").append("<div id='EHtooltip'></div>");
                    $("body").append("<div id='EHpointer'></div>");
	
                    $("#EHpointer").css("top",(aoff.top + 10) + "px").css("left",(aoff.left + 10) + "px").fadeIn("fast");

                    var key = 0;
                    (function() {

                        var value = pathvalue[key++];
                        var func = arguments.callee;

                        var offset = $(value.element).offset();

                        var elHeight = $(value.element).height();
                        var elWidth = $(value.element).width();

                        var pointerx = offset.left + (elWidth/2);
                        var pointery = offset.top + (elHeight/2);

                        var pointerdim = getElementSize($("#EHpointer"));

                        $("#EHpointer").animate({
                            "left": pointerx + "px",
                            "top": pointery + "px"
                        }, 1000, function(event) {

                            $(value.element).addClass(value.marker);
                            var ttip = $("#EHtooltip").css("width", "auto");
                    
                            ttip.html(value.desc);
                    
			   if(value.ftriger != "E") {
                            var efrez = callExtFunction(value);
                    	   }
                    
                            var ttipwidth = ttip.width();
                            var ttipheight = ttip.height();
                    
                            if(options.autoalign != true && value.align != "") {
                                tbalign = value.align;
                            } else {
                                tbalign = getAutoAlign((pointerx + 5), (pointery + 5));
                            }
                    
                            switch(tbalign){
                                case "L":
                                    topset = pointery;
                                    leftset	= (pointerx - ttipwidth - 5);
                                    break;
                                case "LT":
                                    topset = (pointery - ttipheight - 5);
                                    leftset	= (pointerx - ttipwidth - 5);
                                    break;
                                case "LB":
                                    topset = (pointery + ttipheight + 5);
                                    leftset	= (pointerx - ttipwidth - 5);
                                    break;
                    
                                case "R":
                                    topset = pointery;
                                    leftset	= (pointerx + pointerdim.width + 5);
                                    break;
                                case "RT":
                                    topset = (pointery - 5);
                                    leftset	= (pointerx + pointerdim.width + 5);
                                    break;
                                case "RB":
                                    topset = (pointery + pointerdim.height + 5);
                                    leftset	= (pointerx + pointerdim.width + 5);
                                    break;
                            }
                    
                            ttip.css("top",topset + "px")
                            .css("left",leftset + "px")
                            .css("width", ttipwidth + "px")
                            .fadeIn("fast")
                            .delay(value.duration).queue(function () {
                                $(this).dequeue();
                            });
                    
                        }).delay(value.duration).queue(function () {
                            $(value.element).removeClass(value.marker);
                    
				if(value.ftriger == "E") {	
					var efrez = callExtFunction(value);
				}

                            if(key >= (elmax)) {
                                $("#EHtooltip").remove();
                                $("#EHpointer").remove();
                                clearInterval();
                    
                            }
                            $(this).dequeue();
                    
                          }).fadeTo(1, 1, func);
                    })();
                }

                if(options.animatedvp == true) {
                    Htimer = setInterval(checkPointer, 250);
                }

                return false;
            }
            //-----------------------------------
            function doPathStatic(pathvalue){

                $(".EHtooltipc").remove();
                $(".EHtooltmp").remove();


                var dSize = getDocumentSize();

                if(!$.isEmptyObject(pathvalue)) {
                    $.each(pathvalue, function(key, value) {
	
                        var offset = $(value.element).offset();
				
                        var elHeight = $(value.element).height();
                        var elWidth = $(value.element).width();
				
                        var pointerx = offset.left + (elWidth/2);
                        var pointery = offset.top + (elHeight/2);
                        var divhtml = "<span>" + (key + 1) +"</span>";

                        if(options.autoalign != true && value.align != "") {
                            tbalign = value.align;
                        } else {
                            tbalign = getAutoAlign((pointerx + 5), (pointery + 5));
                        }

                        var tipbox = $("<div>").addClass("EHtooltipc").html(divhtml).attr("rel", tbalign).attr("alt", value.desc);
					
                        tipbox.css("top",(pointery + 5) + "px").css("left",(pointerx + 5) + "px").appendTo("body").fadeIn("slow");

                        var efrez = callExtFunction(value);
		
			
                    });
                }
	
                if(options.staticvp == true) {
                    return checkAllStaticBoxes();
                }
		
                return false;
            }
            //-----------------------------------
	    function doBoxLInks(paths){

                    if(!$.isEmptyObject(paths)) {
                        $.each(paths, function(reli, pathvalue) {
                             $('#EHhelpBox').append("<a rel='" + pathvalue.rel + "' href=''>" + pathvalue.link + "</a><br/>");
                        });
                    }
	    }		
	    //-----------------------------------	
            return this.each(function() {  
                var p = pathdefinition;

                var obj = $(this);

                var items = $("a[rel]", obj);

		if(options.autolinks == true) { doBoxLInks(p); }

                items.live('click', function() {

                    var	aobj = $(this);
                    var rel = $(this).attr("rel");
                    var rez;
	
                    doClearAll();
	
                    if(!$.isEmptyObject(p)) {
                        $.each(p, function(reli, pathvalue) {
                            if(pathvalue.rel == rel) {
                                if(pathvalue.method == "animated") {
                                    rez = doPathAnimation(aobj, pathvalue.path);
                                } else {
                                    rez = doPathStatic(pathvalue.path);
                                    activeStaticPath = pathvalue.path;
                                }
                            }
                        });
                    }

                    return false;
  
                });  
                //-----------------------------------
                $('.EHclose').live('click', function() {
                    $('.EHtooltipc').fadeOut('slow').remove();
                    doClearAll();
                    return false;
                });
                //-----------------------------------
                $('.EHstopAll').live('click', function() {
                    doClearAll();
                    return false;
                });
                //-----------------------------------
                $('.EHtooltmp').live('mouseleave', function() {
                    $(this).remove();
                });
                //-----------------------------------

                $('.EHtooltipc').live('mouseover', function() {

                    $(".EHtooltmp").remove();
                    var align = $(this).attr("rel");
                    var toffset = $(this).offset();
                    var towidth = $(this).outerWidth();
                    var toheight = $(this).outerHeight();
                    var number = $(this).find('span').html();

                    var textdesc = activeStaticPath[(number-1)].desc;
			
			
                    switch(align){
                        case "L":
                        case "LT":
                        case "LB":
					
                            var tmpbox = $("<div>").addClass("EHtooltmp").html("<p>" + textdesc + "</p><a href='' class='EHclose'>X</a>").appendTo("body");

                            tmpboxWidth = tmpbox.width();
                            tmpboxHeight = tmpbox.height();
					
                            tmpbox.css("top",(toffset.top) + "px").css("left",(toffset.left - tmpboxWidth) + "px").css("width",tmpboxWidth + "px").css("border-right", "0px").fadeIn("slow");
	
                            break;

                        case "R":
                        case "RT":
                        case "RB":
                            var tmpbox = $("<div>").addClass("EHtooltmp").html("<p>" + textdesc + "</p><a href='' class='EHclose'>X</a>").appendTo("body");

                            tmpboxWidth = tmpbox.width();
                            tmpboxHeight = tmpbox.height();
					
                            tmpbox.css("top",(toffset.top) + "px").css("left",(toffset.left + towidth) + "px").css("width",tmpboxWidth + "px").css("border-left", "0px").fadeIn("slow");
                            break;
                    }
                });
                //--------------------------------------------------------------
                $(window).resize( function ()
                {
                    doPathStatic(activeStaticPath);
                    if(options.staticvp == true) {
                        checkAllStaticBoxes();
                    }
                });

                $(window).scroll( function ()
                {
                    if(options.staticvp == true) {
                        checkAllStaticBoxes();
                    }
                });

                $(window).keydown(function(e) {
                    if (e.keyCode == '27' || e.which == '27') {
                        doClearAll();
                    }
                });
                $(document).keydown(function(e) {
                    if (e.keyCode == '27' || e.which == '27') {
                        doClearAll();
                    }
                });
            //-----------------------------------
            });

        }

    });  
})(jQuery);


var ehelp_topics = [
	{
		"rel":"interface",
		"method":"animated",
		"path": [
			{
				"element": "#toolbarHolder",
				"desc": "Application main menu. Hover your mouse over options to see various commands",
				"duration": 4000,
				"align": "LT",
				"marker": "EHpointer"
			},
			{
				"element": "#dblist",
				"desc": "List of databases on the server. Selection shows the database you are currently working with.",
				"duration": 4000,
				"align": "LT",
				"marker": "EHpointer"
			},
			{
				"element": "#object_list",
				"desc": "Database objects e.g. tables, views, functions are displayed here",
				"duration": 4000,
				"align": "LT",
				"marker": "EHpointer"
			},
			{
				"element": "#screenContent",
				"desc": "Content area contains query results, success/error messages and other information",
				"duration": 4000,
				"align": "LT",
				"marker": "EHpointer"
			},
			{
				"element": "#sqlEditFrame",
				"desc": "Multiple Sql Editors. You can switch between the editors by shortcuts keys to work with multiple queries without overwriting them",
				"duration": 6000,
				"align": "LT",
				"marker": "EHpointer"
			},
			{
				"element": "#nav_bar",
				"desc": "Buttons for performing various result related operations.<br />As you work with results, additional buttons appear here based on the editing state of results",
				"duration": 8000,
				"align": "LT",
				"marker": "EHpointer"
			}
		]
	},
	{
		"rel":"queries",
		"method":"animated",
		"path": [
			{
				"element": "#sqlEditFrame",
				"desc": "Type sql query in one of the sql editors, or select a part of the text to be executed as a query",
				"duration": 4000,
				"align": "LT",
				"marker": "EHpointer"
			},
			{
				"element": "#nav_query",
				"desc": "Click on the [Query] button, or press Ctrl+Enter to run the query",
				"duration": 3000,
				"align": "LT",
				"marker": "EHpointer"
			},
			{
				"element": "#nav_queryall",
				"desc": "To execute multiple queries at once, Click [Query All] or press Ctrl+Shift+Enter",
				"duration": 3000,
				"align": "LT",
				"marker": "EHpointer"
			},
			{
				"element": "#screenContent",
				"desc": "Successful query results will show up in the result pane<br />If an error occurs, the error will be shown in the messages pane",
				"duration": 6000,
				"align": "LT",
				"marker": "EHpointer"
			}
		]
	},
	{
		"rel":"editing_results",
		"method":"animated",
		"path": [
			{
				"element": "#screenContent",
				"desc": "Once results appear in the results pane after a query or table selection, you can:<br />- Sort the results by clicking the column headers<br />- Use the checkboxes in second column to select records for deletion<br />- Double click a cell to edit its value<br />- Use the [Add Record] button to add a new record (which can be later saved by [Update records] button)",
				"duration": 10000,
				"align": "R",
				"marker": "EHpointer"
			},
			{
				"element": "#nav_bar",
				"desc": "Once you have selected/edited the records you want to update, use one of the buttons on the navigation bar to perform an operation",
				"duration": 5000,
				"align": "L",
				"marker": "EHpointer"
			},
			{
				"element": "#nav_bar",
				"desc": "Press [Update records] to generate and execute update queries for edited records",
				"duration": 4000,
				"align": "L",
				"marker": "EHpointer"
			},
			{
				"element": "#nav_bar",
				"desc": "Press [Delete records] to generate and execute queries to delete selected records",
				"duration": 4000,
				"align": "L",
				"marker": "EHpointer"
			},
			{
				"element": "#nav_bar",
				"desc": "Press [Generate SQL] to only generate update/delete queries. You can then inspect the generated sql text and execute as desired after any modification",
				"duration": 10000,
				"align": "L",
				"marker": "EHpointer"
			},
			{
				"element": "#screenContent",
				"desc": "Success/error messages and affected records will be shown in the messages pane",
				"duration": 5000,
				"align": "L",
				"marker": "EHpointer"
			}
		]
	},
	{
		"rel":"objects",
		"method":"animated",
		"path": [
			{
				"element": "#object_list",
				"desc": "Click on any of the objects in the database to perform default action.<br />Clicking on tables/views shows their data in the results pane.<br />For other objects, their creation command is shown by default",
				"duration": 10000,
				"align": "LT",
				"marker": "EHpointer"
			},
			{
				"element": "#object_list",
				"desc": "Right click on any object to see a list of commands that can be performed on that object",
				"duration": 8000,
				"align": "LT",
				"marker": "EHpointer"
			},
			{
				"element": "#object_list",
				"desc": "Some of the commands generate sql queries when executed, which are then automatically added to the first sql editor",
				"duration": 7000,
				"align": "LT",
				"marker": "EHpointer"
			}
		]
	}
];

var showEHelp = function() {
	var _temp_html = '<div id="EHhelpOverlay" class="ui-widget-overlay"></div><div id="EHhelpBox" class="ui-widget-header"><p class="title">Quickstart Tutorials</p>'
		+ '<a rel="interface" href="#">Introduction to interface elements</a>'
		+ '<a rel="queries" href="#">Performing queries</a>'
		+ '<a rel="editing_results" href="#">Editing and saving results</a>'
		+ '<a rel="objects" href="#">Working with database objects</a>'
		+ '<a id="EHstopAll" class="EHstopAll" href="#">Close Tutorial</a></div>';

	$('body').append(_temp_html);
	$('#EHhelpBox').embeddedHelp(ehelp_topics, {
		'animatedvp': true,
		'staticvp': true,
		'autoalign': true,
		'autolinks': false
	});
	$('#EHstopAll').click(function(e) {
		$('#EHhelpOverlay,#EHhelpBox').remove();
		$('#ehelp_css').remove();
		// todo: remove ehelp js from document nodes
	});
	
	$('#EHhelpBox').one('click', function() {
		$(this).animate( {
			'top': '3%',
			'right': '3%'
		});
	});
}