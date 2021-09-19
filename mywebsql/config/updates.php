<?php
/**
 * This file is a part of MyWebSQL package
 * configures various auto update settings
 *
 * @file:      config/updates.php
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

	// you can change how frequently a 'background' check is done by comparing local
	// and available online version of MyWebSQL

	// if true, new version will be checked for and notified on main interface
	$AUTOUPDATE_CHECK = false;

	// days of week on which auto update will run itself
	$AUTOUPDATE_DAYS = array('Mon');

?>