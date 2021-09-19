<link href='/mywebsql/cache?css=theme,default,grid,alerts,editor' rel="stylesheet" />

<div id="popup_wrapper">
	<div id="popup_contents">
		<div class="ui-state-highlight padded"><%= MESSAGE %></div>
		<table width="95%" border="0" cellspacing="1" cellpadding="2" id="table_grid"><tbody>
			<tr id='fhead'>
				<th style="width:40%"><%= T("Table") %></th>
				<th style="width:60%"><%= T("Number of matches") %></th>
			</tr>
			<?php foreach($data['results'] as $table => $info) { ?>
				<tr>
					<td><?php echo $table; ?></td>
					<td><?php
						if ($info['matches'] == 0)
							echo $info['matches'];
						else {
							$txt = str_replace('<%= NUM %>', $info['matches'], __('<%= NUM %> match(es)'));
							$txt .= '&nbsp;[' . __('Copy query to editor') . ']';
							echo '<a href="#">'.$txt.'</a><div style="display:none">'.htmlspecialchars($data['queries'][$table]).'</div>';
						}
					?></td>
				</tr>
			<?php } ?>
		</tbody></table>
	</div>
</div>

<script type="text/javascript" language='javascript' src="/mywebsql/cache?script=common,jquery,ui,position,query"></script>
<script type="text/javascript" language="javascript">
window.title = "<%= T("Search Results") %>";
$(function() {
	$('#table_grid a').click(function() {
		sql = $(this).siblings('div').eq(0).text();
		parent.setSqlCode(parent.sql_delimiter + sql, 1);
	});
});
</script>
