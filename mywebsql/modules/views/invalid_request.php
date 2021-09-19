<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset='utf-8';" />
<title>MyWebSQL</title>
<link rel="stylesheet" type="text/css" href="/mywebsql/cache?css=default" />
<link rel="SHORTCUT ICON" href="favicon.ico" />
</head>
<body style="background-color:white">
<div style="border:none;position:absolute;left:0px;top:0px;width:100%;height:100%;background-color:white;display:block;">
	<table border="0" width="100%" style="height:100%">
		<tr><td height="100%" valign="middle" align="center" style="text-align:center">
			<%= T("It appears that you attempted to submit an invalid request to the server") %>.<br />
			<%= T("The request has been denied. Reloading the page might solve the problem") %>.
			</td>
		</tr>
	</table>
</div>
<script type="text/javascript" language="javascript">
	window.title = "<%= T("Access Denied") %> !";
</script>
</body></html>
