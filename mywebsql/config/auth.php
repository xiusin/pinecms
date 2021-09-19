<?php
/**
 * This file is a part of MyWebSQL package
 * defines authentication mechanism for the application
 *
 * @file:      config/auth.php
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 *
 * Notes:
 *  Changing this file manually might break the application
 *  or create security issues.
 *  Please edit only if you know what you are doing !!!
 */

	// AUTH_TYPE defines the login/startup behaviour of the application
	// NONE    = No userid/password is asked for (NOT recommended)
	// BASIC   = browser requests authentication dialog
	// LOGIN   = User enters userid and password manually
	// CUSTOM  = Use a custom authentication scheme (see docs for details)
	define('AUTH_TYPE', 'LOGIN');

	// if required extensions are available, secure login will be available
	$secure_login_available = (extension_loaded('openssl') && extension_loaded('gmp') && extension_loaded('bcmath'));
	// avoid sending plain text login info for additional security (disabled for HTTPS automatically)
	define('SECURE_LOGIN', $secure_login_available);

	// AUTH_SERVER defines the name of mysql server for connections and authenticating users
	// if AUTH_TYPE is set to LOGIN and there is no server defined in configuration (config/servers.php),
	// then this will be used as default
	define('AUTH_SERVER', 'localhost|mysql5');
	// other examples
	//define('AUTH_SERVER', 'localhost:4040|mysql4');
	//define('AUTH_SERVER', 'localhost|mysqli');
	//define('AUTH_SERVER', 'c:/sqlitedb/|sqlite');

	// for AUTH_TYPE NONE only
	// use the following userid and password to connect to mysql server
	define('AUTH_LOGIN', 'test');
	define('AUTH_PASSWORD', 'test');

?>