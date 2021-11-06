<!doctype html>
<html>

<head>
    <meta charset="utf-8">
    <title>{{ .title }}</title>
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="description" content="{{ .description }}">
    <meta name="keywords" content="{{ .keywords }}">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="/assets/css/laravel.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/flexboxgrid/6.3.1/flexboxgrid.min.css">
    <link rel="stylesheet"
        href="https://fonts.googleapis.com/css?family=Nunito:200,200i,300,300i,400,400i,600,600i,800,800i,900,900i">
    <link rel="stylesheet" href="/assets/font-awesome/css/font-awesome.min.css">
    <link rel="stylesheet" href="/assets/gitalk/gitalk.css">
</head>

<body class="docs language-default scotchified">
    <nav class="main">
        <a href='{{url "/" }}' class="brand nav-block">
            <span>{{ .title }}</span>
        </a>

        <ul class="main-nav">
            <li class=""><a href='{{url "/docs"}}'>文档</a></li>

            <li class="dropdown">
                <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">菜单<span class="caret"></span></a>
            </li>

        </ul>

        <div class="search nav-block invisible">
            <img src="/assets/svg/search.svg" />
            <input placeholder="搜索" type="text" id="search-input" aria-label="搜索" />
        </div>

        <div class="switcher">
            <div class="dropdown">
                <button class="btn dropdown-toggle" type="button" id="dropdownMenu1" data-toggle="dropdown"
                    aria-expanded="true">
                    master
                    <span class="caret"></span>
                </button>
                <ul class="dropdown-menu" role="menu" aria-labelledby="dropdownMenu1">
                    <li role="presentation">
                        <a role="menuitem" tabindex="-1"
                            href="">master</a>
                    </li>
                </ul>
            </div>
        </div>


        <div class="responsive-sidebar-nav">
            <a href="#" class="toggle-slide menu-link btn">&#9776;</a>
        </div>
    </nav>

    <nav id="slide-menu" class="slide-menu" role="navigation">

        <div class="brand">
            <a href="{{ url "/" }}">
                Dcat Page
            </a>
        </div>

        <ul class="slide-main-nav">
            <li><a href="{{ url "/" }}">Home</a></li>

            <li class=""><a href="{{ url "/docs"}}">Documentation</a></li>
        </ul>

        <div class="slide-docs-nav">
            <h2>Documentation</h2>

            <ul class="slide-docs-version">
                <li>
                    <h2>Version</h2>

                    <ul>
                        <li>
                            @if (isset($currentVersion))
                            @include(\DcatPage\view_name('partials.switcher'))
                            @endif
                        </li>
                    </ul>
                </li>
            </ul>

            {{ .index }}
        </div>

    </nav>

    <div class="docs-wrapper container">
        <section class="sidebar">
            <small><a href="#" id="doc-expand" style="font-size: 11px; color: #B8B8B8;">EXPAND ALL</a></small>
            {{ .index }}
        </section>
    
        <article>
            {{ .content }}
    
            <div id="comment-container" style="margin-top: 90px"></div>
        </article>
    </div>

    <footer class="main">
        <p>Built by <b>xiusin</b>. Copyright &copy;
            2019</p>
        <p class="less-significant">
            <a href="http://jackmcdade.com">
                Designed by<br>
                <img src="/assets/svg/jack-mcdade.svg" />
            </a>
        </p>
    </footer>

    <div class="fixed-bottom-btn"><a class="waves-effect waves-light" id="go-top"><i class="fa fa-chevron-up"></i></a></div>
    <style>
        .docs-wrapper article img {
            max-width:  100%;
        }
        
    </style>
    <script>
        var DcatPage = {
            version: '{{.version}}',
            getDocUrl: function (doc) {
                var temp = '{{doc "{doc}.md"}}'; 
                if (location.pathname.indexOf(temp.replace('{doc}.md', '')) !== -1) {
                    return doc + '.md';
                }
                return temp.replace('{doc}', doc);
            },

            config: {{.config}},
        };

        DcatPage.config.comment.id = "{{.version}}/{{.doc}}";

        (function () {
            function indices() {
                var s = document.createElement('script');
                s.type = 'text/javascript';
                s.async = true;
                s.src = '/assets/indices/master.js';
                var x = document.getElementsByTagName('script')[0];
                x.parentNode.insertBefore(s, x);
            }
            setTimeout(indices, 1);
        })();

    </script>
    
    <script src="/assets/gitalk/gitalk.min.js"></script>
    <script src="/assets/js/laravel.js"></script>
    <script src="/assets/js/viewport-units-buggyfill.js"></script>
    <script>window.viewportUnitsBuggyfill.init();</script>

</body>

</html>