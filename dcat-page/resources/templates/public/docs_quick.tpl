<html>

<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title> QUICK UI 文档 - 快速上手 </title>
    <meta name="keywords" content="前端框架,UI框架,QUICK UI 文档,入门,开发,快速上手">
    <link rel="stylesheet" href='{{ url "/quickui/github-markdown.css"}}'>
    <!--<link rel="stylesheet" href="/quickui_doc/static/css/github.min.css">-->
    <link rel="stylesheet" href='{{ url "/quickui/index.css"}}'>
    <link rel="stylesheet" href='{{ url "/quickui/github-markdown-jun.css"}}'>
</head>

<body>
    <div class="doc-header animated">
        <h1 class="logo"> <a href="http://www.uileader.com/quickui_doc/rudiments/"> QUICK UI 文档 </a> </h1>
        <div class="doc-tab">
            <ul class="tabs">
                <li class="tab active"> <a href="http://www.uileader.com/quickui_doc/rudiments/"> 入门 </a> </li>
                <li class="tab "> <a href="http://www.uileader.com/quickui_doc/framework/"> 框架 </a> </li>
                <li class="tab "> <a href="http://www.uileader.com/quickui_doc/component/"> 组件 </a> </li>
            </ul>
        </div>
    </div>
    <div class="doc-content">
        <div class="doc_navtop"> <span class="nav_bar"><em></em>{{ .title }}</span> </div>
        <div class="doc-nav">
            <div class="menu_top"> <img src='{{ url "/quickui/logo.png"}}'> <span>点击进入文档主目录</span>
                <div class="clear"></div>
            </div>
            {{.index}}
        </div>
        <div id="markdown" class="doc-markdown markdown-body">
           {{.content}}
        </div>
    </div>
    <script>
        var currentDocmd = "{{.doc}}"
    </script>
    <script src='{{ url "/quickui/jquery-1.12.4.min.js"}}'></script>
    <script src='{{ url "/quickui/highlight.min.js"}}'></script>
    <script src='{{ url "/quickui/marked.min.js"}}'></script>
    <script src='{{ url "/quickui/index.js"}}'></script> 
</body>

</html>