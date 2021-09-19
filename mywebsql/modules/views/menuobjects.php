<div id="object-context-menus" style="display:none">

<div id="db-menu">
	<ul class="dropdown context ui-state-default">
		<li><a href="javascript:dbSelect([name])" title="<%= T("Use this database") %>"><%= T("Use Database") %></a></li>
	</ul>
</div>

<div id="object-menu">
	<ul class="dropdown context ui-state-default">
		<li><a class="itable" href="javascript:tableCreate([name])" title="<%= T("Create a new table in the database") %>"><%= T("Create Table") %></a></li>
		<li class="option mysql5 mysqli pgsql sqlite sqlite3"><a class="iview" href="javascript:objCreate(1,
		[name])" title="<%= T("Create a new view in the database") %>"><%= T("Create View") %></a></li>
		<li class="option mysql5 mysqli"><a class="iproc" href="javascript:objCreate(2, [name])" title="<%= T("Create a new stored procedure in the database") %>"><%= T("Create Stored Procedure") %></a></li>
		<li class="option mysql5 mysqli pgsql"><a class="ifunc" href="javascript:objCreate(3, [name])" title="<%= T("Create a new user defined function in the database") %>"><%= T("Create Function") %></a></li>
		<li class="option mysql5 mysqli pgsql sqlite sqlite3"><a class="itrig" href="javascript:objCreate(4, [name])" title="<%= T("Create a new trigger in the database") %>"><%= T("Create Trigger") %></a></li>
		<li class="option mysql5 mysqli"><a class="ievt" href="javascript:objCreate(5, [name])" title="<%= T("Create a new event in the database") %>"><%= T("Create Event") %></a></li>
		<li class="option pgsql"><a class="ischm" href="javascript:objCreate(6, [name])" title="<%= T("Create a new schema in the database") %>"><%= T("Create Schema") %>...</a></li>
	</ul>
</div>

<div id="table-menu">
<ul class="dropdown context ui-state-default">
	<li><a href="tableSelect([name])"><%= T("Select statement") %></a></li>
	<li><a href="tableInsert([name])"><%= T("Insert statement") %></a></li>
	<li><a href="tableUpdate([name])"><%= T("Update statement") %></a></li>
	<li class="option mysql4 mysql5 mysqli pgsql"><a href="tableDescribe([name])"><%= T("Describe") %></a></li>
	<li class="option mysql4 mysql5 mysqli sqlite sqlite3"><a href="showCreateCmd('table',	[name])"><%= T("Show create command") %></a></li>
	<li><a href="tableViewData([name])"><%= T("View data") %></a></li>
	<li class="option mysql4 mysql5 mysqli sqlite sqlite3"><a href="javascript:void(0)"><?php echo __('Alter Table');?> &raquo;</a>
		<ul class="ui-state-default">
			<li class="option mysql4 mysql5 mysqli"><a href="javascript:tableAlter([name])"><%= T("Structure") %></a></li>
			<li class="option mysql4 mysql5 mysqli"><a href="tableIndexes([name])"><%= T("Indexes") %></a></li>
			<li class="option mysql4 mysql5 mysqli"><a href="tableEngine([name])"><%= T("Engine Type") %></a></li>
		</ul>
	</li>
	<li><a href="javascript:void(0)"><%= T("More operations") %> &raquo;</a>
		<ul class="ui-state-default">
			<li><a class="itrunc" href="objTruncate('table', [name])"><%= T("Truncate") %></a></li>
			<li><a class="idrop" href="objDrop('table', [name])"><%= T("Drop") %></a></li>
			<li><a class="iren" href="objRename('table', [name])"><%= T("Rename") %></a></li>
			<li><a class="icopy" href="objCopy('table', [name])"><%= T("Create Copy") %></a></li>
		</ul>
	</li>
	<li class="separator">-------------------------------------------------------</li>
	<li><a class="iexprt" href="tableExport([name])"><%= T("Export table data") %></a></li>
	<li class="option mysql4 mysql5 mysqli sqlite sqlite3"><a class="itable" href="tableCreate()"><%= T("Create Table") %></a></li>
</ul>
</div>

<div id="view-menu">
<ul class="dropdown context ui-state-default">
	<li><a href="tableSelect([name])"><%= T("Select statement") %></a></li>
	<li class="option mysql4 mysql5 mysqli"><a href="tableDescribe([name])"><%= T("Describe") %></a></li>
	<li><a href="showCreateCmd('view', [name])"><%= T("Show create command") %></a></li>
	<li><a href="tableViewData([name])"><%= T("View data") %></a></li>
	<li><a href="objCreate(1)"><%= T("Create View") %></a></li>
	<li><a href="javascript:void(0)"><%= T("More operations") %> &raquo;</a>
		<ul class="ui-state-default">
			<li><a href="objDrop('view', [name])"><%= T("Drop") %></a></li>
			<li class="option mysql4 mysql5 mysqli sqlite sqlite3"><a href="objRename('view', [name])"><%= T("Rename") %></a></li>
			<li class="option mysql4 mysql5 mysqli sqlite sqlite3"><a class="icopy" href="objCopy('view', [name])"><%= T("Create Copy") %></a></li>
		</ul>
	</li>
</ul>
</div>

<div id="proc-menu">
<ul class="dropdown context ui-state-default">
	<li><a href="showCreateCmd('procedure', [name])"><%= T("Show create command") %></a></li>
	<li><a href="objCreate(2)"><%= T("Create Procedure") %></a></li>
	<li><a href="javascript:void(0)"><%= T("More operations") %> &raquo;</a>
		<ul class="ui-state-default">
			<li><a href="objDrop('procedure', [name])"><%= T("Drop") %></a></li>
			<li><a href="objRename('procedure', [name])"><%= T("Rename") %></a></li>
			<li><a class="icopy" href="objCopy('procedure', [name])"><%= T("Create Copy") %></a></li>
		</ul>
	</li>
</ul>
</div>

<div id="func-menu">
<ul class="dropdown context ui-state-default">
	<li class="option mysql4 mysql5 mysqli sqlite sqlite3"><a href="showCreateCmd('function', [name])"><%= T("Show create command") %></a></li>
	<li><a href="objCreate(3)"><%= T("Create Function") %></a></li>
	<li class="option mysql4 mysql5 mysqli sqlite sqlite3"><a href="javascript:void(0)"><%= T("More operations") %> &raquo;</a>
		<ul class="ui-state-default">
			<li><a href="objDrop('function', [name])"><%= T("Drop") %></a></li>
			<li><a href="objRename('function', [name])"><%= T("Rename") %></a></li>
			<li><a class="icopy" href="objCopy('function', [name])"><%= T("Create Copy") %></a></li>
		</ul>
	</li>
</ul>
</div>

<div id="trig-menu">
<ul class="dropdown context ui-state-default">
	<li class="option mysql4 mysql5 mysqli sqlite sqlite3"><a href="showCreateCmd('trigger', [name])"><%= T("Show create command") %></a></li>
	<li><a href="objCreate(4)"><%= T("Create Trigger") %></a></li>
	<li class="option mysql4 mysql5 mysqli sqlite sqlite3"><a href="javascript:void(0)"><?php echo __('More
	operations'); ?> &raquo;</a>
		<ul class="ui-state-default">
			<li><a href="objDrop('trigger', [name])"><%= T("Drop") %></a></li>
		</ul>
	</li>
</ul>
</div>

<div id="evt-menu">
<ul class="dropdown context">
	<li><a href="showCreateCmd('event', [name])"><%= T("Show create command") %></a></li>
	<li><a href="objCreate(5)"><%= T("Create Event") %></a></li>
	<li><a href="javascript:void(0)"><%= T("More operations") %> &raquo;</a>
		<ul>
			<li><a href="objDrop('event', [name])"><%= T("Drop") %></a></li>
			<!--li><a href="objRename('event', [name])"><%= T("Rename") %></a></li-->
		</ul>
	</li>
</ul>
</div>

<div id="schm-menu">
<ul class="dropdown context">
	<li><a href="objCreate(6)"><%= T("Create Schema") %></a></li>
	<li><a href="javascript:void(0)"><%= T("More operations") %> &raquo;</a>
		<ul>
			<li><a href="objDrop('schema', [name])"><%= T("Drop") %></a></li>
		</ul>
	</li>
</ul>
</div>

<div id="panel-header">
<ul class="dropdown context">
	<li><a href="main_layout.toggle('north')"><%= T("Show/Hide Header") %></a></li>
</ul>
</div>

<div id="panel-menu-objects">
<ul class="dropdown context">
	<li><a href="main_layout.toggle('west')"><%= T("Show/Hide Panel") %></a></li>
</ul>
</div>

<div id="panel-menu-editor">
<ul class="dropdown context">
	<li><a href="data_layout.toggle('south')"><%= T("Show/Hide Panel") %></a></li>
</ul>
</div>

<div id="history-menu">
<ul class="dropdown context">
	<li class="clipboard single"><a href="javascript:void(0)" title="<%= T("Copy to clipboard") %>"><%= T("Copy to clipboard") %></a></li>
	<li class="clipboard"><a href="javascript:void(0)" title="<%= T("Copy all queries to clipboard") %>"><%= T("Copy all queries to clipboard") %></a></li>
	<li><a href="javascript:historyClear($(this))" title="<%= T("Clear all queries from history") %>"><%= T("Clear history") %></a></li>
</ul>
</div>

<div id="data-menu-th">
<ul class="dropdown context">
	<li><a href="copyColumn([name])"><%= T("Copy Column values") %></a></li>
</ul>
</div>

<div id="data-menu-td">
<ul class="dropdown context">
	<li><a href="copyText([name])"><%= T("Copy to clipboard") %></a></li>
	<li><a href="sqlFilterText([name])"><%= T("Generate SQL Filter") %></a></li>
</ul>
</div>

</div>