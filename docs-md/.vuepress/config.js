module.exports = {
    title: 'PineCMS内容管理系统',
    description: '一个GO语言和Sqlite开发的内容管理系统',
    markdown: {
        lineNumbers: true,
    },
    themeConfig: {
        nav: [
            {text: '入门', link: '/guide/'},
            {text: '模型系统', link: '/models/'},
            {text: '标签系统', link: '/taglibs/'},
            {text: '程序开发', link: '/develop/'},
        ],
        sidebar: {
            "/guide": [
                "/guide/",
                "/guide/installation",
                "/guide/directory-structure",
            ],
            "/models": [
                "/models/",
                "/models/import_dede",
                "/models/import_dede_tpl",
            ],

            '/taglibs': [
                '/taglibs/',
                '/taglibs/adlist',
                '/taglibs/myad',
                '/taglibs/query',
                '/taglibs/type',
                '/taglibs/flink',
                '/taglibs/artlist',
                '/taglibs/prenext',
                '/taglibs/pagelist',
                '/taglibs/likearticle',
                '/taglibs/hotwords',
                '/taglibs/tags',
                '/taglibs/position',
                '/taglibs/toptype',
            ],
            "/develop": [
                "/develop/",
                "/develop/controller",
                "/develop/view",
                "/develop/model",
                "/develop/example",
            ],
        },
        sidebarDepth: 0,
        smoothScroll: true,
        repo: 'https://github.com/xiusin/pinecms',
    },
    theme: 'antdocs',
    plugins: [
        [
            "homebadge", {
            selector: '.hero',
            repoLink: 'https://github.com/xiusin/pinecms',
            badgeLink: 'https://img.shields.io/github/stars/xiusin/pinecms?style=social',
            badgeGroup: [
                'https://img.shields.io/badge/build-passing-brightgreen?style=flat-square',
                'https://img.shields.io/github/license/xiusin/pinecms?style=flat-square&color=blue'
            ]
        }],
        ['@vuepress/search', {searchMaxSuggestions: 10}],
        "@vuepress/back-to-top",
        ["@vuepress/medium-zoom", true],
        "@vuepress/active-header-links",
    ],
}