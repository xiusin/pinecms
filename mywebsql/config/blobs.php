<?php
/**
 * This file is a part of MyWebSQL package
 * outputs scripts and stylesheets for the application
 *
 * @file:      config/blobs.php
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

	// Descriptive name, byte header, content header, html tag replacement, function to decrypt/deserialize blob data
	$blobTypes = array(
				"txt" =>
					array("Text/Binary", "", ""),
				"jpg" =>
					array("Jpeg", array("\xFF\xD8\xFF\xE0", "\xFF\xD8\xFF\xE1"), "Content-type: Image/jpeg", "<img src='#link#'/ >"),
				"gif" =>
					array("GIF", "\x47\x49\x46\x38\x39\x61", "Content-type: Image/gif", "<img src='#link#'/ >"),
				"png" =>
					array("PNG", "\x89\x50\x4E\x47\x0D\x0A\x1A\x0A", "Content-type: Image/png", "<img src='#link#'/ >"),
				"bmp" =>
					array("Bitmap", "\x42\x4D\x36\x10", "Content-type: Image/bmp", "<img src='#link#'/ >"),
				//"ico" =>
				//	array("Icon", "", ""),
				"ser" =>
					array("Serialized data", "", "", "", "unserialize"),
			);

?>