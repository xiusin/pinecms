export default {
    // 可用磁盘
    disks: [],

    /**
     * 管理器
     * 左(默认) 或 右
     */
    activeManager: 'left',

    /**
     * 剪贴板
     * 操作类型：复制 或 剪切
     */
    clipboard: {
        type: null,
        disk: null,
        directories: [],
        files: [],
    },
    fileCallback: null,

    /**
     * 登录状态
     */
    isLogin: false,

    /**
     * 自动登录
     */
    autoLogin: false,
    /**
     * 用户名
     */
    username: "",
    /**
     * 昵称
     */
    nickname: "",

    // 全屏模式
    // fullScreen: false,
};