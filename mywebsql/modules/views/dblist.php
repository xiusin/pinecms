<select name="dblist" id='dblist'  onchange='dbSelect()' style="width: 100%">
	<%= for (key, value) in dblist { %>
        <option value="<%= value %>" <%= if (curdb == value) { %> selected <% } %> ><%= value %> </option>
    <% } %>
</select>
