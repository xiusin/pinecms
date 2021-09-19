<?php
/**
 * This file is a part of MyWebSQL package
 *
 * @file:      config/config.php
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

	define('TRACE_MESSAGES', FALSE);            // logs verbose stuff in the error log file (only enable for debugging)
	define('TRACE_FILEPATH', "");               // if logs are to be directed to a separate file other than the default
	define('LOG_MESSAGES', FALSE);              // enabling this will send 'critical' messages to the default log file (including failed queries)
	define('MAX_RECORD_TO_DISPLAY', 100);       // only this much records will be shown in browser at one time to keep it responsive
	define('MAX_TEXT_LENGTH_DISPLAY', 80);      // blobs/text size larger than this is truncated in grid view format
	define('HOTKEYS_ENABLED', TRUE);            // enable hotkeys

	define('DEFAULT_EDITOR', "codemirror");     // if not set by the user, this editor will be used
	define('DEFAULT_THEME', 'default');         // if not set by the user, this theme will be used
	define('DEFAULT_LANGUAGE', 'en');           // if not set by the user, this langauge will be used for the interface
?>