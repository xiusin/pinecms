module.exports = {
    title: 'PineCMS内容管理系统',
    description: '一个GO + Sqlite/Mysql开发的CMS管理系统',
    themeConfig: {
        nav: [
            {text: '快速上手', link: '/guide/'},
            {text: '二次开发', link: '/develop/'},
            {text: '标签系统', link: '/taglibs/'},
            {
                text: 'github', items: [
                    {text: "PineCMS", link: 'https://github.com/xiusin/pinecms.git'},
                    {text: "Pine", link: 'https://github.com/xiusin/pine.git'},
                ]
            },
        ],
        displayAllHeaders: true,
        sidebar: [
            {
                title: "安装部署",
                collapsable: true,
                children: [
                    '/guide/using_unix',
                    '/guide/using_windows',
                ]
            },
            ["/guide/directory-structure", "目录结构"],
            {
                title: "文档系统",
                collapsable: true,
                children: [
                    '/models/',
                ]
            },
            {
                title: "二次开发",
                collapsable: true,
                children: [
                    '/develop/',
                ]
            },
            {
                title: "特殊页面",
                collapsable: true,
                children: [
                    '/pages/',
                    '/pages/search',
                ]
            },
            {
                title: "标签系统",
                collapsable: true,
                children: [
                    '/taglibs/',
                    '/taglibs/tags_adgroup',
                    '/taglibs/tags_ad',
                    '/taglibs/tags_myad',
                    '/taglibs/tags_query',
                    '/taglibs/tags_type',
                    '/taglibs/tags_flink',
                    '/taglibs/tags_artlist',
                ]
            },
            [ 'https://github.com/CloudyKit/jet/wiki', "JetTemplate Wiki"],
            [ 'https://github.com/xiusin/pine', "Pine"],
        ]
    }
}