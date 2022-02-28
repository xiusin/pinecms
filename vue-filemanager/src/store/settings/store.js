import mutations from './mutations';
import getters from './getters';

/* eslint camelcase: 0 */
import zh_CN from '../../lang/zh_CN';
export default {
  namespaced: true,
  state() {
    return {
      // windowsConfig: 2,

      // ACL
      acl: null,


      // axios headers
      headers: {
        "X-Requested-With": "XMLHttpRequest",
        Authorization: window.localStorage.getItem("user-token") ? `cfk ${window.localStorage.getItem("user-token")}` : null,
      },
      // axios default URL
      baseUrl: "",

      lang: 'zh-CN',

      translations: {
        'zh-CN': Object.freeze(zh_CN),
      },

      //显示隐藏文件
      hiddenFiles: false,

      // 上下文菜单项
      contextMenu: [
        [{
            name: "aa",
            icon: 'far fa-folder-open',
          },
          {
            name: 'open',
            icon: 'far fa-folder-open',
          },
          {
            name: 'audioPlay',
            icon: 'fas fa-play',
          },
          {
            name: 'videoPlay',
            icon: 'fas fa-play',
          },
          {
            name: 'view',
            icon: 'fas fa-eye',
          },
          {
            name: 'edit',
            icon: 'fas fa-file-signature',
          },
          {
            name: 'select',
            icon: 'fas fa-check',
          },
          {
            name: 'download',
            icon: 'fas fa-download',
          },
        ],
        [{
            name: 'copy',
            icon: 'far fa-copy',
          },
          {
            name: 'cut',
            icon: 'fas fa-cut',
          },
          {
            name: 'rename',
            icon: 'far fa-edit',
          },
          {
            name: 'paste',
            icon: 'far fa-clipboard',
          },
          {
            name: 'zip',
            icon: 'far fa-file-archive',
          },
          {
            name: 'unzip',
            icon: 'far fa-file-archive',
          },
        ],
        [{
          name: 'delete',
          icon: 'far fa-trash-alt text-danger',
        }, ],
        [{
          name: 'properties',
          icon: 'far fa-list-alt',
        }, ],
      ],

      // 用户预览文档
      officeExtensions: ['xls', 'xlsx', 'doc', 'docx', 'ppt', 'pptx'],

      // 用于查看和预览的图像扩展
      imageExtensions: ['png', 'jpg', 'jpeg', 'gif'],

      // 裁剪的image扩展
      cropExtensions: ['png', 'jpg', 'jpeg'],

      // 可播放音频的扩展
      audioExtensions: ['ogg', 'mp3', 'aac', 'wav'],

      // 可播放视频的扩展
      videoExtensions: ['webm', 'mp4'],

      // 代码编辑器的文件扩展名
      textExtensions: {
        sh: 'text/x-sh',
        // styles
        css: 'text/css',
        less: 'text/x-less',
        sass: 'text/x-sass',
        scss: 'text/x-scss',
        html: 'text/html',
        // js
        js: 'text/javascript',
        ts: 'text/typescript',
        vue: 'text/x-vue',
        // text
        htaccess: 'text/plain',
        env: 'text/plain',
        txt: 'text/plain',
        log: 'text/plain',
        ini: 'text/x-ini',
        xml: 'application/xml',
        md: 'text/x-markdown',
        // c-like
        java: 'text/x-java',
        c: 'text/x-csrc',
        cpp: 'text/x-c++src',
        cs: 'text/x-csharp',
        scl: 'text/x-scala',
        php: 'application/x-httpd-php',
        // DB
        sql: 'text/x-sql',
        // other
        pl: 'text/x-perl',
        py: 'text/x-python',
        lua: 'text/x-lua',
        swift: 'text/x-swift',
        rb: 'text/x-ruby',
        go: 'text/x-go',
        yaml: 'text/x-yaml',
        json: 'application/json',
      },
    };
  },
  mutations,
  getters,
};
