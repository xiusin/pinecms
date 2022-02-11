import HTTP from '../../../http/get';

/**
 * 上下文菜单 actions
 * {name}Action
 */
export default {
  methods: {
    /**
     * 打开文件夹
     */
    openAction() {
      // 选择文件夹
      this.$store.dispatch(`fm/${this.$store.state.fm.activeManager}/selectDirectory`, {
        path: this.selectedItems[0].path,
        history: true,
      });
    },

    /**
     * 播放音频
     */
    audioPlayAction() {
      // 显示播放模块
      this.$store.commit('fm/modal/setModalState', {
        modalName: 'AudioPlayer',
        show: true,
      });
    },

    /**
     * 播放视频
     */
    videoPlayAction() {
      // 显示播放模块
      this.$store.commit('fm/modal/setModalState', {
        modalName: 'VideoPlayer',
        show: true,
      });
    },

    /**
     * 浏览文件
     */
    viewAction() {
      // 显示预览图
      this.$store.commit('fm/modal/setModalState', {
        modalName: 'Preview',
        show: true,
      });
    },

    /**
     * 编辑文件
     */
    editAction() {
      // 显示文本文件
      this.$store.commit('fm/modal/setModalState', {
        modalName: 'TextEdit',
        show: true,
      });
    },

    /**
     * 选择文件
     */
    selectAction() {
      // 文件回调
      this.$store.dispatch('fm/url', {
        disk: this.selectedDisk,
        path: this.selectedItems[0].path,
      }).then((response) => {
        if (response.data.result.status === 'success') {
          this.$store.state.fm.fileCallback(response.data.url);
        }
      });
    },

    /**
     * 下载文件
     */
    downloadAction() {
      const tempLink = document.createElement('a');
      tempLink.style.display = 'none';
      tempLink.setAttribute('download', this.selectedItems[0].basename);

      // 有权限则下载文件
      if (this.$store.getters['fm/settings/authHeader']) {
        // HTTP.download(this.selectedDisk, this.selectedItems[0].path).then((response) => {
        //   tempLink.href = window.URL.createObjectURL(new Blob([response.data]));
        //   document.body.appendChild(tempLink);
        //   tempLink.click();
        //   document.body.removeChild(tempLink);
        // });
        HTTP.downloadFile(this.selectedDisk, this.selectedItems[0].path).then((response) => {
          if (typeof response.data === 'string' && /\.txt$/g.test(this.selectedItems[0].path)) {
            tempLink.href = window.URL.createObjectURL(new Blob([response.data]));
          } else {
            tempLink.href = response.data;
          }
          document.body.appendChild(tempLink);
          tempLink.click();
          document.body.removeChild(tempLink);
        });
      } else {
        tempLink.href = `${this.$store.getters['fm/settings/baseUrl']}download?disk=${this.selectedDisk}&path=${encodeURIComponent(this.selectedItems[0].path)}`;
        document.body.appendChild(tempLink);
        tempLink.click();
        document.body.removeChild(tempLink);
      }
    },

    /**
     * 复制选中的文件项
     */
    copyAction() {
      // 添加到剪切板
      this.$store.dispatch('fm/toClipboard', 'copy');
    },

    /**
     * 剪切选中的文件项
     */
    cutAction() {
      // 添加到剪切板
      this.$store.dispatch('fm/toClipboard', 'cut');
    },

    /**
     * 重命名选中的文件项
     */
    renameAction() {
      // 显示命名模块
      this.$store.commit('fm/modal/setModalState', {
        modalName: 'Rename',
        show: true,
      });
    },

    /**
     * 粘贴复制或剪切的文件项
     */
    pasteAction() {
      // 在选中的目录中粘贴选中的文件项
      this.$store.dispatch('fm/paste');
    },

    /**
     * 压缩选中的文件项
     */
    zipAction() {
      // 显示压缩模块
      this.$store.commit('fm/modal/setModalState', {
        modalName: 'Zip',
        show: true,
      });
    },

    /**
     * 解压选中的压缩包
     */
    unzipAction() {
      // 显示解压模块
      this.$store.commit('fm/modal/setModalState', {
        modalName: 'Unzip',
        show: true,
      });
    },

    /**
     * 删除选中的文件项
     */
    deleteAction() {
      // 显示删除模块
      this.$store.commit('fm/modal/setModalState', {
        modalName: 'Delete',
        show: true,
      });
    },

    /**
     * 显示选中的文件项属性
     */
    propertiesAction() {
      // 显示属性模块
      this.$store.commit('fm/modal/setModalState', {
        modalName: 'Properties',
        show: true,
      });
    },
  },
};