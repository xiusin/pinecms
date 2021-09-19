<div style='float:left;margin-left:10px' id="blobedit_buttons">
<input type="file" name="blobdata" id="blobdata" size="12" value="<%= T("Import File") %>" />
<a id="btnSaveBlob" class="btn" href="javascript:blobSave()">Save</a>
</div>
<?php echo str_replace('<%= TYPE %>', '', __('Show blob data as: <%= TYPE %>')); ?>
<select name='blobtype' onchange="blobChangeType()"><%= BLOBOPTIONS %></select>