<ul id="tablelist" class="filetree">
<%= if(len(tables) > 0) { %>
    <li id="tables"><span class="tablef"><%= T("Tables") %></span><span class="count"><%= len(tables) %></span>
		<%= for (key, v) in tables { %>
		    <ul>
                <li>
                    <span class="file otable" id="t_<%= v.Name %>">
                    <a href='javascript:objDefault("table", "t_<%= v.Name %>")'><%= v.Name %></a></span><!-- <span class="count"></span> --> <!--不要换行-->
                </li>
		    </ul>
		<% } %>
	</li>
<% } %>

<%= if(len(views) > 0) { %>
    <li id="views"><span class="viewf"><%= T("Views") %></span><span class="count"><%= len(views) %></span>
		<%= for (key, v) in views { %>
		    <ul>
                <li><span class="file oview" id="t_<%= v.Name %>">
                    <a href='javascript:objDefault("view", "t_<%= v.Name %>")'><%= v.Name %></a></span>
                </li>
		    </ul>
		<% } %>
	</li>
<% } %>
</ul>

