<?php
/**
 * This file is a part of MyWebSQL package
 *
 * @file:      modules/views/templates/variables.php
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

	function parseVariables(&$db) {
		$vars = array();
		while($row = $db->fetchRow())
		{
			switch($row[0]) {
				case "version": $vars['SERVER_VERSION'] = $row[1]; break;
				case "version_comment": $vars['SERVER_COMMENT'] = $row[1]; break;
				case "character_set_server": $vars['SERVER_CHARSET'] = $row[1]; break;
				case "character_set_client": $vars['CLIENT_CHARSET'] = $row[1]; break;
				case "character_set_database": $vars['DATABASE_CHARSET'] = $row[1]; break;
				case "character_set_results": $vars['RESULT_CHARSET'] = $row[1]; break;
			}
		}
		$vars['SERVER_NAME'] = 'MySQL';
		return $vars;
	}
?>
