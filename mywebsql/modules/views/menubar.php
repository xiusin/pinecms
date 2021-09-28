<div class="ui-state-default ui-corner-all ui-helper-clearfix">
	<ul id="main-menu" class="dropdown">
		<li>
			<a href="javascript:void(0)"><%= T("Database") %></a>
			<ul class="ui-state-default">
				<li class="db"><a class="irfrsh" href="javascript:objectsRefresh()" title="<%= T("Refresh database object list") %>"><%= T("Refresh") %></a></li>
				<li><a href="javascript:dbCreate()" title="<%= T("Create a new database") %>"><%= T("Create new") %>...</a></li>
				<li class="db option mysql4 mysql5 mysqli mysql sqlite sqlite3 pgsql"><a class="ibatch" href="javascript:dbBatch()" title="<%= T("Perform one or more batch operations on database") %>"><%= T("Batch operations") %>...</a></li>
				<li class="db option mysql4 mysql5 pgsql mysql mysqli sqlite sqlite3"><a class="iexpdb" href="javascript:dataExport()" title="<%= T("Export database to external file") %>"><%= T("Export") %>...</a></li>
			</ul>
		</li>
		<li class="db">
			<a href="javascript:void(0)"><%= T("Objects") %></a>
			<ul class="ui-state-default">
				<li class="option mysql5 mysql mysqli"><a class="itable" href="javascript:tableCreate()" title="<%= T("Create a new table in the database") %>"><%= T("Create Table") %>...</a></li>
				<li class="option sqlite sqlite3 pgsql"><a class="itable" href="javascript:javascript:objCreate(0)" title="<%= T("Create a new table in the database") %>"><%= T("Create Table") %>...</a></li>
				<li class="option mysql5 mysql mysqli pgsql sqlite sqlite3"><a class="iview" href="javascript:objCreate(1)" title="<%= T("Create a new view in the database") %>"><%= T("Create View") %>...</a></li>
				<li class="option mysql5 mysql mysqli"><a class="iproc" href="javascript:objCreate(2)" title="<%= T("Create a new stored procedure in the database") %>"><%= T("Create Stored Procedure") %>...</a></li>
				<li class="option mysql5 mysql mysqli pgsql"><a class="ifunc" href="javascript:objCreate(3)" title="<%= T("Create a new user defined function in the database") %>"><%= T("Create Function") %>...</a></li>
				<li class="option mysql5 mysql mysqli pgsql sqlite sqlite3"><a class="itrig" href="javascript:objCreate(4)" title="<%= T("Create a new trigger in the database") %>"><%= T("Create Trigger") %>...</a></li>
				<li class="option mysql5 mysql mysqli"><a class="ievt" href="javascript:objCreate(5)" title="<%= T("Create a new event in the database") %>"><%= T("Create Event") %>...</a></li>
				<li class="option pgsql"><a class="ischm" href="javascript:objCreate(6)" title="<%= T("Create a new schema in the database") %>"><%= T("Create Schema") %>...</a></li>
			</ul>
		</li>

		<li style="display:none">
			<a href="javascript:void(0)"><%= T("Data") %></a>
			<ul class="ui-state-default">
				<li class="option mysql4 mysql mysql5 mysqli pgsql sqlite sqlite3"><a class="iimprt" href="javascript:dataImport()" title="<%= T("Import multiple queries from batch file") %>"><%= T("Import batch file") %>...</a></li>
				<li class="option mysql4 mysql mysql5 mysqli pgsql sqlite sqlite3"><a class="iimprt" href="javascript:tableImport()" title="<%= T("Import table data from external file") %>"><%= T("Import table data") %>...</a></li>
				<li class="db option mysql4 mysql mysql5 pgsql mysqli sqlite sqlite3"><a class="iexpdb" href="javascript:dataExport()" title="<%= T("Export database to batch file as sql dump") %>"><%= T("Export database") %>...</a></li>
				<li class="db"><a class="iexprt" href="javascript:resultsExport()" title="<%= T("Export query results to clipboard or files") %>"><%= T("Export current results") %>...</a></li>
				<li class="db option mysql4 mysql5 mysql mysqli"><a class="iexpdb" href="javascript:dataBackup()" title="<%= T("Backup database on the server as sql dump") %>"><%= T("Backup Database") %>...</a></li>
			</ul>
		</li>

		<li class="option mysql4 mysql5 mysql mysqli pgsql sqlite sqlite3">
			<a href="javascript:void(0)"><%= T("Tools") %></a>
			<ul class="ui-state-default">
				<li class="option mysql4 mysql5 mysql mysqli pgsql"><a class="itopts" href="javascript:toolsDbManager()" title="<%= T("Manage databases") %>"><%= T("Database Manager") %></a></li>
				<li class="option mysql4 mysql5 mysql mysqli"><a class="itprc" href="javascript:toolsProcManager()" title="<%= T("View and manage database processes") %>"><%= T("Process Manager") %></a></li>
				<!-- <li class="option mysql4 mysql5 mysql mysqli"><a class="itusr" href="javascript:toolsUsers()" title="<%= T("Manage database users") %>"><%= T("User Manager") %></a></li> -->
				<li class="db option mysql4 mysql5 mysql pgsql mysqli"><a class="itchk" href="javascript:toolsDbCheck()" title="<%= T("Analyze and repair database tables") %>"><%= T("Repair Tables") %></a></li>
				<!-- <li class="db option mysql4 mysql5 mysql mysqli pgsql sqlite sqlite3"><a class="itsrch" href="javascript:toolsDbSearch()" title="<%= T("Search for text in the database") %>"><%= T("Search in Database") %></a></li> -->
				<!-- <li><a class="itopts" href="javascript:toolsOptions()" title="<%= T("Configure application options") %>"><%= T("Options") %></a></li> -->
			</ul>
		</li>

		<li>
			<a href="javascript:void(0)"><%= T("Information") %></a>
			<ul class="ui-state-default">
				<li class="option mysql4 mysql5 mysql mysqli pgsql"><a href="javascript:infoServer()" title="<%= T("View server and connection details") %>"><%= T("Server/Connection Details") %></a></li>
				<li class="option mysql4 mysql5 mysql mysqli pgsql"><a href="javascript:infoVariables()" title="<%= T("View server configuration") %>"><%= T("Server Variables") %></a></li>
				<li class="db"><a href="javascript:infoDatabase()" title="<%= T("View current database summary stats") %>"><%= T("Database Summary") %></a></li>
			</ul>
		</li>

		<li>
			<a href="javascript:void(0)"><%= T("Interface") %></a>
			<ul>
				<li><a href="javascript:void(0)"><%= T("UI Theme") %></a>
					<ul class="ui-state-default">
						<%= THEMES_MENU %>
					</ul>
				</li>

				<li><a href="javascript:void(0)"><%= T("SQL Editor") %></a>
					<ul class="ui-state-default">
						<%= EDITOR_MENU %>
					</ul>
				</li>
				<li><a href="javascript:void(0)"><%= T("Show/Hide Panel") %></a>
					<ul class="ui-state-default">
						<li><a href="javascript:main_layout.toggle('west')" title="<%= T("Toggle Object Viewer") %>"><%= T("Database Objects") %></a></li>
						<li><a href="javascript:data_layout.toggle('south')" title="<%= T("Toggle SQL Editor") %>"><%= T("SQL Editor") %></a></li>
					</ul>
				</li>
			</ul>
		</li>

		<li>
			<a href="javascript:void(0)"><%= T("Help") %></a>
			<ul class="ui-state-default">
				<li><a class="ihlp" href="javascript:helpShowAll()" title="<%= T("Learn the basics of using MyWebSQL") %>"><%= T("Help contents") %></a></li>
				<li class="db"><a class="itutor" href="javascript:helpQuickTutorial()" title="<%= T("See quick hands-on tutorial of MyWebSQL interface") %>"><%= T("QuickStart Tutorials") %></a></li>
				<!-- <li><a class="idocs" href="javascript:helpOnlineDocs()" title="<%= T("View online documentation on project website") %>"><%= T("Online documentation") %></a></li> -->
				<!-- <li><a class="iftr" href="javascript:helpRequestFeature()" title="<%= T("If you would like your most favourite feature to be part of MyWebSQL, please click here to inform about it") %>"><%= T("Request a Feature") %>...</a></li> -->
				<!-- <li><a class="ibug" href="javascript:helpReportBug()" title="<%= T("If you have found a problem, or having trouble using the application, please click here to report the problem") %>"><%= T("Report a Problem") %></a></li> -->
				<!-- <li><a href="javascript:helpCheckUpdates()" title="<%= T("Check for updated versions of the application online") %>"><%= T("Check for updates") %></a></li> -->
			</ul>
		</li>

		<li class="right"><a class="ilgout" href="javascript:logout()" title="<%= T("Logout from this session") %>"><%= T("Logout") %></a></li>
	</ul>
</div>