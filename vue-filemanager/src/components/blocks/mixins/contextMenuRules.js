/**
 * 上下文菜单项显示/隐藏的规则
 * {name}Rule
 * @returns {boolean}
 */
export default {
  methods: {
    /**
     * 打开 - 显示/隐藏
     * @returns {boolean}
     */
    openRule() {
      return !this.multiSelect && this.firstItemType === 'dir';
    },

    /**
     * 播放音频 - 显示/隐藏
     * @returns {boolean}
     */
    audioPlayRule() {
      return this.selectedItems.every((elem) => elem.type === 'file') &&
        this.selectedItems.every((elem) => this.canAudioPlay(elem.extension));
    },

    /**
     * 播放视频 - 显示/隐藏
     * @returns {boolean}
     */
    videoPlayRule() {
      return !this.multiSelect && this.canVideoPlay(this.selectedItems[0].extension);
    },

    /**
     * 查看- 显示/隐藏
     * @returns {boolean|*}
     */
    viewRule() {
      return !this.multiSelect &&
        this.firstItemType === 'file' &&
        this.canView(this.selectedItems[0].extension);
    },

    /**
     * 编辑 - 显示/隐藏
     * @returns {boolean|*}
     */
    editRule() {
      return !this.multiSelect &&
        this.firstItemType === 'file' &&
        this.canEdit(this.selectedItems[0].extension);
    },

    /**
     * 选择 - 显示/隐藏
     * @returns {boolean|null}
     */
    selectRule() {
      return !this.multiSelect && this.firstItemType === 'file' &&
        this.$store.state.fm.fileCallback;
    },

    /**
     * 下载- 显示/隐藏
     * @returns {boolean}
     */
    downloadRule() {
      return !this.multiSelect && this.firstItemType === 'file';
    },

    /**
     * 复制 - 显示/隐藏
     * @returns {boolean}
     */
    copyRule() {
      return true;
    },

    /**
     * 剪切 - 显示/隐藏
     * @returns {boolean}
     */
    cutRule() {
      return true;
    },

    /**
     * 重命名 - 显示/隐藏
     * @returns {boolean}
     */
    renameRule() {
      return !this.multiSelect;
    },

    /**
     * 粘贴 - 显示/隐藏
     * @returns {boolean}
     */
    pasteRule() {
      return !!this.$store.state.fm.clipboard.type;
    },

    /**
     * 压缩 - 显示/隐藏
     * @returns {boolean}
     */
    zipRule() {
      return this.selectedDiskDriver === 'local';
    },

    /**
     * 解压 - 显示/隐藏
     * @returns {boolean}
     */
    unzipRule() {
      return this.selectedDiskDriver === 'local' &&
        !this.multiSelect &&
        this.firstItemType === 'file' &&
        this.isZip(this.selectedItems[0].extension);
    },

    /**
     * 删除 - 显示/隐藏
     * @returns {boolean}
     */
    deleteRule() {
      return true;
    },

    /**
     * 文件属性 - 显示/隐藏
     * @returns {boolean}
     */
    propertiesRule() {
      return !this.multiSelect;
    },
  },
};