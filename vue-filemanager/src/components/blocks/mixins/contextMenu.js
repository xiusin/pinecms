export default {
  computed: {
    /**
     * 当前所在磁盘
     * @returns {*}
     */
    selectedDisk() {
      return this.$store.getters['fm/selectedDisk'];
    },

    /**
     * 选中的文件或文件夹
     * @returns {*}
     */
    selectedItems() {
      return this.$store.getters['fm/selectedItems'];
    },

    /**
     * 判断选中的文件的acl
     * @returns 
     */
    selectedItemAcl() {
      return this.selectedItems.every(function (item) {
        return item.acl === 1;
      })
    },
    /**
     * 判断当前选中文件的创建者是否与当前用户一致
     * @returns 
     */
    authorIsUser() {
      let username = this.$store.state.fm.username;
      return this.selectedItems.every(function (item) {
        return username === item.author;
      })
    },
    /**
     * 选择的磁盘盘符
     * @returns {*}
     */
    selectedDiskDriver() {
      return this.$store.state.fm.disks[this.selectedDisk].driver;
    },

    /**
     * 多选择
     * @returns {boolean}
     */
    multiSelect() {
      return this.$store.getters['fm/selectedItems'].length > 1;
    },

    /**
     * 选中的文件类型是目录还是文件
     * @returns {*}
     */
    firstItemType() {
      return this.$store.getters['fm/selectedItems'][0].type;
    },
  },
  methods: {
    /**
     * 能否查看该图片
     * @param extension
     * @returns {boolean}
     */
    canView(extension) {
      // 扩展未找到
      if (!extension) return false;

      return this.$store.state.fm.settings.imageExtensions.includes(extension.toLowerCase());
    },

    /**
     * 能否在代码编辑器中编辑该文件
     * @param extension
     * @returns {boolean}
     */
    canEdit(extension) {
      // 扩展未找到
      if (!extension) return false;

      return Object.keys(this.$store.state.fm.settings.textExtensions)
        .includes(extension.toLowerCase());
    },

    /**
     * 能否播放音频文件
     * @param extension
     * @returns {boolean}
     */
    canAudioPlay(extension) {
      // 扩展未找到
      if (!extension) return false;

      return this.$store.state.fm.settings.audioExtensions.includes(extension.toLowerCase());
    },

    /**
     * 能否播放视频文件
     * @param extension
     * @returns {boolean}
     */
    canVideoPlay(extension) {
      // 扩展未找到
      if (!extension) return false;

      return this.$store.state.fm.settings.videoExtensions.includes(extension.toLowerCase());
    },

    /**
     * 能否压缩
     * @param extension
     * @returns {boolean}
     */
    isZip(extension) {
      // 扩展未找到
      if (!extension) return false;

      return extension.toLowerCase() === 'zip';
    },
  },
};