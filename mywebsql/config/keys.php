<?php
/**
 * This file is a part of MyWebSQL package
 * definition of various keycodes used in browser
 *
 * @file:      config/keys.php
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

	// Change the setting below if you like to change shortcut key behaviour

	$KEY_CODES = array(
		'KEYCODE_SETNULL'        => array('shift+del', "Shift + Del"),    	// sets value to NULL during edit
		'KEYCODE_QUERY'          => array('ctrl+return', "Ctrl + Enter"), 	// single query
		'KEYCODE_QUERYALL'       => array('ctrl+shift+return', "Ctrl + Shift + Enter"),	// query all
		'KEYCODE_SWITCH_EDITOR1' => array('alt+1', "Alt + 1"),
		'KEYCODE_SWITCH_EDITOR2' => array('alt+2', "Alt + 2"),
		'KEYCODE_SWITCH_EDITOR3' => array('alt+3', "Alt + 3"),
		'KEYCODE_EDITOR_TEXTSIZE_INC' => array('ctrl+up', 'Ctrl + Up Arrow'),
		'KEYCODE_EDITOR_TEXTSIZE_DEC' => array('ctrl+down', 'Ctrl + Down Arrow'),
		'KEYCODE_EDITOR_CLEARTEXT' => array('ctrl+shift+del', 'Ctrl + Shift + Del')
	);

	// You should not change anything below unless you know what you are doing!

	$DOCUMENT_KEYS = array(
		'KEYCODE_SETNULL'       => 'closeEditor(true, null)',
		'KEYCODE_SWITCH_EDITOR1' => 'switchEditor(0)',
		'KEYCODE_SWITCH_EDITOR2' => 'switchEditor(1)',
		'KEYCODE_SWITCH_EDITOR3' => 'switchEditor(2)'

	);

	$SIMPLE_KEYS = array(
		'KEYCODE_QUERY'     => 'queryGo(0)',
		'KEYCODE_QUERYALL'  => 'queryGo(1)'
	);

	$CODEMIRROR_KEYS = array(
		'KEYCODE_QUERY'     => 'queryGo(0)',
		'KEYCODE_QUERYALL'  => 'queryGo(1)',
		'KEYCODE_SWITCH_EDITOR1' => 'switchEditor(0)',
		'KEYCODE_SWITCH_EDITOR2' => 'switchEditor(1)',
		'KEYCODE_SWITCH_EDITOR3' => 'switchEditor(2)',
		'KEYCODE_EDITOR_TEXTSIZE_INC' => 'editorTextSize(0.2)',
		'KEYCODE_EDITOR_TEXTSIZE_DEC' => 'editorTextSize(-0.2)',
		'KEYCODE_EDITOR_CLEARTEXT' => 'editorClear()'
	);

	$CODEMIRROR2_KEYS = array(
		'KEYCODE_QUERY'     => 'queryGo(0)',
		'KEYCODE_QUERYALL'  => 'queryGo(1)'
	);
?>