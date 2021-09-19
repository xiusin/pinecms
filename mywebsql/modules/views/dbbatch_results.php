<link href='/mywebsql/cache?css=theme,default,grid,alerts' rel="stylesheet" />

<div id="popup_wrapper">
	<div id="popup_contents">
		<div class="ui-state-highlight padded"><%= T("Batch operation results") %></div>
		<table width="95%" border="0" cellspacing="1" cellpadding="2" id="table_grid"><tbody>
			<tr id='fhead'>
				<th style="width:30%"><%= T("Operation") %></th>
				<th style="width:70%"><%= T("Status") %></th>
			</tr>
			<?php if( isset($data['stats']['drop']) ) { ?>
				<tr>
					<td><?php echo  __('DROP selected database objects'); ?></td>
					<td><?php
						$txt = '<p><span class="ui-icon ui-icon-check"></span>' . str_replace('<%= NUM %>', $data['stats']['drop']['success'], __('<%= NUM %> queries successfully executed')) . '</p>';
						if ( $data['stats']['drop']['errors'] > 0 )
							$txt .= '<p><span class="ui-icon ui-icon-close"></span>' . str_replace('<%= NUM %>', $data['stats']['drop']['errors'], __('<%= NUM %> queries failed to execute')) . '</p>';
						echo $txt;
					?></td>
				</tr>
			<?php } ?>
			<?php if( isset($data['stats']['delprefix']) ) { ?>
				<tr>
					<td><?php echo  __('Delete prefix string from name'); ?></td>
					<td><?php
						$txt = '<p><span class="ui-icon ui-icon-check"></span>' . str_replace('<%= NUM %>', $data['stats']['delprefix']['success'], __('<%= NUM %> queries successfully executed')) . '</p>';
						if ( $data['stats']['delprefix']['errors'] > 0 )
							$txt .= '<p><span class="ui-icon ui-icon-close"></span>' . str_replace('<%= NUM %>', $data['stats']['delprefix']['errors'], __('<%= NUM %> queries failed to execute')) . '</p>';
						echo $txt;
					?></td>
				</tr>
			<?php } ?>
			<?php if( isset($data['stats']['addprefix']) ) { ?>
				<tr>
					<td><?php echo  __('Add prefix string to name'); ?></td>
					<td><?php
						$txt = '<p><span class="ui-icon ui-icon-check"></span>' . str_replace('<%= NUM %>', $data['stats']['addprefix']['success'], __('<%= NUM %> queries successfully executed')) . '</p>';
						if ( $data['stats']['addprefix']['errors'] > 0 )
							$txt .= '<p><span class="ui-icon ui-icon-close"></span>' . str_replace('<%= NUM %>', $data['stats']['addprefix']['errors'], __('<%= NUM %> queries failed to execute')) . '</p>';
						echo $txt;
					?></td>
				</tr>
			<?php } ?>
			<?php if( isset($data['queries']) && count($data['queries']) > 0 ) { ?>
				<tr>
					<td><?php echo  __('Generate SQL'); ?></td>
					<td><?php
						$txt = '<p><span class="ui-icon ui-icon-check"></span>' . str_replace('<%= NUM %>', count($data['queries']), __('<%= NUM %> queries generated')) . '</p>';
						echo $txt;
					?></td>
				</tr>
			<?php } ?>
		</tbody></table>

		<div style="display:none" id="command_list">
			<?php
				foreach($data['queries'] as $query)
					echo htmlspecialchars($query) . ";\n<br />";
			?>
		</div>
	</div>
</div>

<script type="text/javascript" language='javascript' src="/mywebsql/cache?script=common,jquery,ui,position,query"></script>
<script type="text/javascript" language="javascript">
window.title = "<%= T("Batch operation results") %>";
<?php
	$refresh = false;
	if ( isset($data['stats']['drop']) && $data['stats']['drop']['success'] > 0 )
		$refresh = true;
	else if ( isset($data['stats']['delprefix']) && $data['stats']['delprefix']['success'] > 0 )
		$refresh = true;
	else if ( isset($data['stats']['addprefix']) && $data['stats']['addprefix']['success'] > 0 )
		$refresh = true;

	if ($refresh)
		echo "parent.objectsRefresh();\n";

	if ( count($data['queries']) > 0)
		echo "parent.setSqlCode(parent.sql_delimiter + $('#command_list').text().trim(), 1);\n";
?>
</script>








































