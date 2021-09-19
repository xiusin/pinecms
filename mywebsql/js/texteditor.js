/**
 * This file is a part of MyWebSQL package
 *
 * @file:      js/texteditor.js
 * @author     Samnan ur Rehman
 * @copyright  (c) 2008-2014 Samnan ur Rehman
 * @web        http://mywebsql.net
 * @license    http://mywebsql.net/license
 */

textEditor = function(id) {
	this.textarea = $(id);
};

textEditor.prototype.focus = function() {
	return this.textarea.focus();
};

textEditor.prototype.getCode = function(s) {
	return this.textarea.val();
};

textEditor.prototype.setCode = function(s) {
	this.textarea.val(s);
};

textEditor.prototype.canHighlight = function() {
	return false;
}

textEditor.prototype.highlightSql = function() {
	return false;
}

textEditor.prototype.getSelection = function(s) {
	return this.textarea.val();
};

// @@TODO: this just goes to the last line, need to fix that
textEditor.prototype.jumpToLine = function(n) {
	h = this.textarea.prop("scrollHeight");
	this.textarea.prop("scrollTop", h);
}

textEditor.prototype.lastLine = function() {
	return -1;
}
