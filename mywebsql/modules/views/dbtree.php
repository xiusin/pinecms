<ul id="tablelist" class="dblist">
	<%= for (key, value) in dblist { %>
       <li><span class="odb"><a href="javascript:dbSelect('<%= value %>')"><%= value %></a></span>
    <% } %>
</ul>
