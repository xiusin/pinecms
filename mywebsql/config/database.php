<?php
/**
 * This file is a part of MyWebSQL package
 * defines per server database access list to restrict connection to a list of databases
 *
 * @file:      config/database.php
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

	// match the keys of this array to those defined in servers.php in the same folder
	// keys that do not exist in servers list will be simply ignored
	// if a server's database access is not defined here, then it will show all databases
	//     (given that the logged in user has access to all those databases)
	$DB_LIST = array(
		'Test Server' =>      array('test', 'wordpress'),
		//'Test Server 2'' =>    array('test')
	);
?>