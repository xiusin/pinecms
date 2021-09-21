<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset='utf-8';" />
<title>MyWebSQL</title>
	<link rel="stylesheet" type="text/css" href="/mywebsql/cache?css=theme,default" />
	<link rel="SHORTCUT ICON" href="favicon.ico" />
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, minimum-scale=1, user-scalable=no, minimal-ui">
	<meta name="apple-mobile-web-app-capable" content="yes">
	<meta name="apple-mobile-web-app-status-bar-style" content="black">
	<link rel="stylesheet" type="text/css" href="/mywebsql/cache?css=menu,treeview,results,context,alerts" />
	<script type="text/javascript" language="javascript" src="/mywebsql/cache?script=jquery,mobile,mysql"></script>
	<!--[if lt IE 8]>
		<script type="text/javascript" language="javascript" src="/mywebsql/cache?script=json2"></script>
	<![endif]-->
</head>
<body class="mainbody">

	<div id="editToolbar">
		<div class="tb-header ui-widget-header"><span class="fname"></span></div>
		<div class="tb-row">Type: <span class="ftype"></span></div>
		<div class="tb-row">[ <%= KEYCODE_SETNULL %> ]</div>
	</div>

	<div id="inplace-text">
		<div class="tb-row"><textarea rows="4" cols="20"></textarea></div>
	</div>

	<div class="ui-layout-north">
		<div id="main_header">
			<a target="_blank" href="/mywebsql/index/index"><img src="/mywebsql/img/logo.png" class="logo" alt="MyWebSQL" width="45" height="38" border="0" /></a>
			<div class="title">
				<div class="main">MyWebSQL</div>
				<div class="version"><%= T("version") + APP_VERSION %></div>
			</div>
			<div class="info">
				<span class="server"><%= auth.ServerName %></span> - <%= version_comment %>&nbsp;
				<%= version_full %><br />
				<%= LoginUser %>
			</div>
			<div class="updates ui-state-active"></div>
		</div>
		<div id="toolbarHolder">
        <%= MenuBarHTML %>
	</div>
</div>

<div class="ui-layout-west">

	<div id="db_combo" class="ui-state-default">
		<%= dbListHtml %>
	</div>

	<div id="object_list" class="ui-state-default">
		<%= treeHtml %>
	</div>

	<div id="object-filter" class="ui-state-default">
		<input type="text" id="object-filter-text" size="5" placeholder="<%= T("Type to filter object list") %>" data-placeholder="<%= T("Type to filter object list") %>" />
	</div>

</div>

<div class="ui-layout-center">
	<div id="screenContent" class="ui-layout-data-center">
		<ul>
			<li><a href="#tab-results" id="headerResults"><%= T("Results") %></a></li>
			<li><a href="#tab-messages" id="headerMessages"><%= T("Messages") %></a></li>
			<li><a href="#tab-info" id="headerInfo"><%= T("Information") %></a></li>
			<li><a href="#tab-history" id="headerHistory"><%= T("History") %></a></li>
		</ul>

		<div id="screen-pane-buttons">
			<button id="sp-results-maximize" title="<%= T("Maximize/Restore Results Pane") %>"></button>
		</div>

		<div class="ui-layout-content ui-corner-bottom">
			<div id="tab-results">
				<div id="results-div"><div class="message" style="text-align:center"><%= T("There are no results to show in this view") %></div></div>
				<div id="rec_pager">
					<table width="100%" border="0" cellpadding="0" cellspacing="0" class="ui-state-default">
					<tr>
						<td id="recordCounter" class="footer" nowrap="nowrap">&nbsp;</td>
						<td id="timeCounter" class="footer" nowrap="nowrap">&nbsp;</td>
						<td id="modifyFlag" class="footer" nowrap="nowrap"><!--button id="nav_refresh">Refresh results</button--></td>
						<td id="messageContainer" class="footer" nowrap="nowrap"><%= T("Please wait") %> ...</td>
						<td id="pagingContainer" class="footer" nowrap="nowrap"></td>
					</tr>
					</table>
				</div>
			</div>
			<div id="tab-messages">
				<div id="messages-div"></div>
			</div>
			<div id="tab-info" class="ui-widget">
				<div id="info-div"></div>
			</div>
			<div id="tab-history">
				<table id="sql-history" width="100%"><tbody><tr><td></td></tr></tbody></table>
			</div>
		</div>
	</div>
	<div id="sql-editor-pane" class="ui-layout-data-south">

		<ul>
			<li><a href="#editor_container"><%= T("SQL Editor") %></a></li>
			<li><a href="#editor_container2"><%= T("SQL Editor") %> 2</a></li>
			<li><a href="#editor_container3"><%= T("SQL Editor") %> 3</a></li>
		</ul>

		<div id="editor_container">
			<textarea class="sql-editor" id="commandEditor" name="commandEditor" rows="5" cols="40"></textarea>
		</div>
		<div id="editor_container2">
			<textarea class="sql-editor" id="commandEditor2" name="commandEditor2" rows="5" cols="40"></textarea>
		</div>
		<div id="editor_container3">
			<textarea class="sql-editor" id="commandEditor3" name="commandEditor3" rows="5" cols="40"></textarea>
		</div>

		<div id="nav_bar">
			<button id="nav_queryall"><%= T("Query All") %></button>
			<button id="nav_query"><%= T("Query") %></button>
			<button id="nav_addrec"><%= T("Add Record") %></button>
			<button id="nav_copyrec"><%= T("Copy Record(s)") %></button>
			<button id="nav_delete"><%= T("Delete Record(s)") %></button>
			<button id="nav_update"><%= T("Update Record(s)") %></button>
			<button id="nav_gensql"><%= T("Generate SQL") %></button>
		</div>

		<div id="loader">
			<img width="60" height="12" id="loaderImg" src="themes/<%= THEME_PATH %>/images/loading.gif" alt="<%= T("Loading") %>..." />
		</div>

	</div>
</div>

<div class="ui-layout-south">
	<div id="taskbar" class="ui-state-default">
		<button class="min-all" title="<%= T("Minimize All") %>"><%= T("Minimize All") %></button>
	</div>
</div>
<%= dialogs %>
<iframe src="javascript:false" name="wrkfrm" id="wrkfrm" frameborder="0" width="0" height="0"></iframe>
<div id="screen-wait" class="ui-widget-overlay">
	<div><span><%= T("Loading") %>...</span><img src="themes/<%= THEME_PATH %>/images/loading.gif" alt="" /></div>
	<div class="compat-notice" style="display:none;margin:200px auto;width:530px;color: #222222;font-family:segoe ui;font-size:13pt;font-weight:bold">
		<%= T("Your browser appears to be very old and does not support all features required to run MyWebSQL.") %><br /><br />
		<%= T("Try using a newer version of the browser to run this application.") %>
	</div>
</div>
<script type="text/javascript" language="javascript">
	var THEME_PATH = "<%= THEME_PATH %>";
	var EXTERNAL_PATH = "<%= EXTERNAL_PATH %>";
	var COOKIE_LIFETIME = 1440; // hours
	var MAX_TEXT_LENGTH_DISPLAY = <%= MAX_TEXT_LENGTH_DISPLAY %>;
	var APP_LANGUAGE = "zh";
	var APP_VERSION = "<%= APP_VERSION %>";
	var DB_DRIVER = "mysql";
	var DB_VERSION = 8; // 5 - 8 根据实际连接数据库版本
	var DB = db_mysql;
	var commandEditor = null;
	var commandEditor2 = null;
	var commandEditor3 = null;
// <?php
//
// 	if (Session::get('db', 'changed')) {
// 		echo 'document.getElementById("messageContainer").innerHTML = "Database changed to: '.htmlspecialchars(Session::get('db', 'name')).'";';
// 		Session::del('db', 'changed');
// 	}
// 	else
// 		echo 'document.getElementById("messageContainer").innerHTML = "Connected to: '.htmlspecialchars(DB_HOST).' as '.htmlspecialchars(DB_USER).'";';
// ?>
</script>
<script type="text/javascript" language="javascript" src="/mywebsql/cache?script=layout,ui,dialogs,context,alerts,cookies,select,interface,options,treeview,common,taskbar,settings,query,tables,clipboard"></script>
<%= contextMenusHTML %>
<%= UpdateSqlEditor %>
<%= HotkeysHTML %>
</body></html>
