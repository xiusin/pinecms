<?php
/**
 * This file is a part of MyWebSQL package
 * outputs scripts and stylesheets for the application
 * @file:      cache.php
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */
  	define('BASE_PATH', dirname(__FILE__));

	$useCache = file_exists('js/min/minify.txt');
	include(BASE_PATH . '/modules/configuration.php');
	initConfiguration(false);

	$fileList = v($_REQUEST["script"]);
	// concat theme path to make etags unique per theme
	if ($fileList == '')	$fileList = THEME_PATH . v($_REQUEST["css"]);
	if ($fileList == '')	exit();
		
	header("X-Frame-Options: SAMEORIGIN"); 

	// cache scripts and css per version, if not in development mode
	if ($useCache) {
		$versionTag = md5($fileList.APP_VERSION);
		$eTag = v($_SERVER['HTTP_IF_NONE_MATCH']);
		if ($eTag != '' && $versionTag == $eTag) {
			header($_SERVER['SERVER_PROTOCOL'].' 304 Not Modified');
			header('Content-Length: 0');
			exit();
		}
		header('Etag: '.$versionTag);
	}

	include(BASE_PATH . "/lib/functions.php");
	include(BASE_PATH . "/lib/output.php");

	Output::buffer();

	$regex = '#^(\w+/){0,2}\w+$#';

	if (v($_REQUEST["script"]) != "")
	{
		$script_path = BASE_PATH . ($useCache ? "/js/min" : "/js");
		$scripts = explode(",", $_REQUEST["script"]);
		header("mime-type: text/javascript");
		header("content-type: text/javascript");
		echo "/**\n * This file is a part of MyWebSQL package\n * @web        http://mywebsql.net\n * @license    http://mywebsql.net/license\n */\n\n";
		foreach($scripts as $script)
			if ( preg_match($regex, $script) == 1 )
				if(file_exists("$script_path/$script".".js"))
					echo file_get_contents("$script_path/$script".".js") . "\n\n";
	}
	else if (v($_REQUEST["css"]) != "")
	{
		$styles = explode(",", $_REQUEST["css"]);
		header("mime-type: text/css");
		header("content-type: text/css");
		echo "/**\n * This file is a part of MyWebSQL package\n * @web        http://mywebsql.net\n * @license    http://mywebsql.net/license\n */\n\n";
		$code = '';
		foreach($styles as $css) {
			if ( preg_match($regex, $css) == 1 ) {
				if(file_exists(BASE_PATH . "/themes/_base/$css".".css"))
					$code .= file_get_contents(BASE_PATH . "/themes/_base/$css".".css");
				if(file_exists(BASE_PATH . "/themes/".THEME_PATH."/$css".".css"))
					$code .= file_get_contents(BASE_PATH . "/themes/".THEME_PATH."/$css".".css") . "\n\n";
			}
			echo $code;
		}
	}

	Output::flush();
?>