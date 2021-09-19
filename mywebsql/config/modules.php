<?php
/**
 * This file is a part of MyWebSQL package
 * defines the list of modules to be allowed/disallowed from access
 *
 * @file:      config/modules.php
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 *
 * Notes:
 *  First, set the type of MODULE_ACCESS_MODE to 'allow' or 'deny'
 *  then define the include array, or exclude array based on above preference
 */

	// Module access mode, can be either 'allow' or 'deny'
	// allow = Only the modules defined in allowed list will be accessible
	// deny = all modules, except those defined in deny list will be accessible
	define('MODULE_ACCESS_MODE', 'deny');

	// list of modules that should be allowed access. All other modules will be
	//   denied access by default
	$ALLOW_MODULES = array(
	);

	// list of modules that should be denied access. All other modules will be
	//   accessible by default
	$DENY_MODULES = array(
		// example: to deny access to the 'Database Manager' tool, which is potentially dangerous,
		// uncomment the following line
		//'databases'
	);

?>