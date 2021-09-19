<?php
/**
 * This file is a part of MyWebSQL package
 * defines settings related to server side backups
 *
 * @file:      config/backups.php
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

	// the folder on server where the backups should be saved ( default = './backups/' )
	// make sure it contains trailing slash
	define( 'BACKUP_FOLDER', BASE_PATH . '/backups/' );

	// back up filename format
	// (extension will be added automatically based on the type of output format)

	// you can use the following variables in the format which will be replaced at runtime by their values
	// <db> = name of the database
	// <date> = current datetime (see below for more)
	// <ext> = filename extension
	define( 'BACKUP_FILENAME_FORMAT', '<db>-<date><ext>' );

	// date format to be used in the backup filename (see php manual for format specifiers)
	define( 'BACKUP_DATE_FORMAT', 'Ymd-His' );
?>