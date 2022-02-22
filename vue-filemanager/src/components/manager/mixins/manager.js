import EventBus from '@/eventBus.js';

export default {
    computed: {
        /**
         * 当前选择的磁盘
         * @returns {String}
         */
        selectedDisk() {
            return this.$store.state.fm[this.manager].selectedDisk;
        },

        /**
         * 该磁盘下当前路径中所有选择的文件夹
         * @returns {String}
         */
        selectedDirectory() {
            return this.$store.state.fm[this.manager].selectedDirectory;
        },

        /**
         * 该磁盘下当前路径中所有文件
         * @returns {*}
         */
        files() {
            return this.$store.getters[`fm/${this.manager}/files`];
        },

        /**
         * 该磁盘下当前路径中所有文件夹
         * @returns {*}
         */
        directories() {
            return this.$store.getters[`fm/${this.manager}/directories`];
        },

        /**
         * 该磁盘下当前路径中选择的文件和文件夹
         * @returns {Object}
         */
        selected() {
            return this.$store.state.fm[this.manager].selected;
        },

        /**
         * ACL on/off
         */
        acl() {
            return this.$store.state.fm.settings.acl;
        },

        /**
         * 检查当前路径是否为根目录
         * @return {Boolean}
         */
        isRootPath() {
            return this.$store.state.fm[this.manager].selectedDirectory === null;
        },
        isCheckedAll: {
            get: function () {
                // 是否全选
                return this.$store.getters[`fm/${this.manager}/isCheckedAll`];
            },
            // set: function (val) {
            // 设置全选或不全选
            // this.commit(`fm/${this.manager}/setIsCheckedAll`, val);
            // }
        },
        checkedFiles: {
            get: function () {
                // 选择的文件列表
                return this.$store.getters[`fm/${this.manager}/checkedFilesList`];
            },
            // 添加选择的文件
            set: function (item) {
                console.log(item);
                //     if (item.type === "dir") {
                //         let alreadySelected = this.checkSelect("directories", item.path);
                //         console.log("是否选择了该文件夹", alreadySelected);
                //         if (!alreadySelected) {
                //             this.$store.commit(`fm/${this.manager}/setSelected`, {
                //                 type: "directories",
                //                 path: item.path
                //             });
                //         }
                //     } else {
                //         let alreadySelected = this.checkSelect("files", item.path);
                //         console.log("是否选择了该文件", alreadySelected);
                //         if (!alreadySelected) {
                //             this.$store.commit(`fm/${this.manager}/setSelected`, {
                //                 type: "files",
                //                 path: item.path
                //             });
                //         }
                //     }
            }
        },
    },
    methods: {
        /**
         *  加载所选目录并显示文件
         * @param path
         */
        selectDirectory(path) {
            this.$store.dispatch(`fm/${this.manager}/selectDirectory`, {
                path,
                history: true
            });
        },

        /**
         * 返回上一级
         */
        levelUp() {
            if (this.selectedDirectory) {
                // 计算目录路径
                const pathUp = this.selectedDirectory.split('/').slice(0, -1).join('/');
                // 加载目录
                this.$store.dispatch(`fm/${this.manager}/selectDirectory`, {
                    path: pathUp || null,
                    history: true
                });
            }
        },

        /**
         * 检查该文件或文件夹是否已选择
         * @param type
         * @param path
         */
        checkSelect(type, path) {
            return this.selected[type].includes(path);
        },

        /**
         * 全选或全删
         * @param { Boolean } isAll
         * @returns
         */
        setAllSelected(isAll = false) {
            if (isAll) {
                this.$store.commit(`fm/${this.manager}/setAllSelected`, {
                    dir: this.directories,
                    file: this.files
                })
                return;
            }
            this.$store.commit(`fm/${this.manager}/removeAllSelected`);
        },
        /**
         *(grid模块)多选
         * @param {*} type
         * @param {*} item
         * @returns
         */
        mutliGridSelected(type, item) {
            // 在所选数组中搜索
            let alreadySelected = this.checkSelect(type, item.path);
            console.log("是否选择了该文件", alreadySelected);
            if (!alreadySelected) {
                // 添加新的选择项
                this.$store.commit(`fm/${this.manager}/setSelected`, {
                    type,
                    path: item.path
                });

            } else {
                // 删除所选项目
                this.$store.commit(`fm/${this.manager}/removeSelected`, {
                    type,
                    path: item.path
                });
            }
            let isCheckedAll =
                this.checkedFiles.length === this.files.length + this.directories.length;
            let isIndeterminate =
                this.checkedFiles.length > 0 &&
                this.checkedFiles.length < this.files.length + this.directories.length;
            this.$store.commit(`fm/${this.manager}/setChAndIn`, {
                isIndeterminate,
                isCheckedAll
            })
        },
        /**
         * (grid视图)中选择文件
         * @param {String} type
         * @param {String} item
         */
        selectGridItem(type, item) {
            // 在所选数组中搜索
            const alreadySelected = this.selected[type].includes(item.path);
            if (!alreadySelected) {
                this.$store.commit(`fm/${this.manager}/changeSelected`, {
                    type,
                    path: item.path,
                });
            } else {
                this.$store.commit(`fm/${this.manager}/removeAllSelected`);
                // 添加新的选择项
                this.$store.commit(`fm/${this.manager}/setSelected`, {
                    type,
                    path: item.path
                });
            }
            let isCheckedAll =
                this.checkedFiles.length === this.files.length + this.directories.length;
            let isIndeterminate =
                this.checkedFiles.length > 0 &&
                this.checkedFiles.length < this.files.length + this.directories.length;
            this.$store.commit(`fm/${this.manager}/setChAndIn`, {
                isIndeterminate,
                isCheckedAll
            })
        },
        /**
         * 文件夹和文件多选
         * @param {String} type
         * @param {Object} item
         */
        mutliSelected(type, item) {
            // 在所选数组中搜索
            let alreadySelected = this.checkSelect(type, item.path);
            console.log("是否选择了该文件", alreadySelected);
            if (!alreadySelected) {
                // 添加新的选择项
                this.$store.commit(`fm/${this.manager}/setSelected`, {
                    type,
                    path: item.path
                });

            } else {
                // 删除所选项目
                this.$store.commit(`fm/${this.manager}/removeSelected`, {
                    type,
                    path: item.path
                });
            }
            let isCheckedAll =
                this.checkedFiles.length === this.files.length + this.directories.length;
            let isIndeterminate =
                this.checkedFiles.length > 0 &&
                this.checkedFiles.length < this.files.length + this.directories.length;
            this.$store.commit(`fm/${this.manager}/setChAndIn`, {
                isIndeterminate,
                isCheckedAll
            })
        },
        /**
         *(table模块) 选择列表中的项目（ 文件 + 文件夹）
         * @param type
         * @param path
         * @param event
         */
        selectItem(type, item) {
            // 在所选数组中搜索
            const alreadySelected = this.selected[type].includes(item.path);
            if (!alreadySelected) {
                this.$store.commit(`fm/${this.manager}/changeSelected`, {
                    type,
                    path: item.path,
                });
            } else {
                this.$store.commit(`fm/${this.manager}/removeAllSelected`);
                // 添加新的选择项
                this.$store.commit(`fm/${this.manager}/setSelected`, {
                    type,
                    path: item.path
                });
            }

            // 如果按Ctrl则为多选
            // if (event.ctrlKey) {
            //     if (!alreadySelected) {
            //         console.log(111);
            //         // 添加新的选择项
            //         this.$store.commit(`fm/${this.manager}/setSelected`, {
            //             type,
            //             path: item.path
            //         });
            //     } else {
            //         console.log(222);
            //         // 删除所选项目
            //         this.$store.commit(`fm/${this.manager}/removeSelected`, {
            //             type,
            //             path: item.path
            //         });
            //     }
            // }

            // 单选
            // if (!event.ctrlKey && !alreadySelected) this.$store.commit(`fm/${this.manager}/changeSelected`, {
            //     type,
            //     path: item.path
            // });
        },

        /**
         * 显示上下文菜单
         * @param item
         * @param event
         * @param obj
         */
        contextMenu(item, event) {
            // 选中的项目类型
            const type = item.type === 'dir' ? 'directories' : 'files';
            // 在所选数组中搜索
            const alreadySelected = this.selected[type].includes(item.path);
            // 选择此项目
            if (!alreadySelected) {
                this.$store.commit(`fm/${this.manager}/changeSelected`, {
                    type,
                    path: item.path,
                });
            }
            // 触发上下文菜单事件
            EventBus.$emit('contextMenu', event);
        },

        /**
         * 选择要触发的事件
         * @param path
         * @param extension
         */
        selectAction(path, extension) {
            // 如果设置了fileCallback
            if (this.$store.state.fm.fileCallback) {
                this.$store.dispatch('fm/url', {
                    disk: this.selectedDisk,
                    path,
                }).then((response) => {
                    if (response.data.result.status === 'success') {
                        this.$store.state.fm.fileCallback(response.data.url);
                    }
                });
                return;
            }

            // 如果扩展未定义
            if (!extension) {
                return;
            }

            if (this.$store.state.fm.settings.officeExtensions
                .includes(extension.toLowerCase())) {
                // 显示图片预览
                this.$store.commit('fm/modal/setModalState', {
                    modalName: 'OfficeViewer',
                    show: true,
                });
            } else if (this.$store.state.fm.settings.imageExtensions
                .includes(extension.toLowerCase())) {
                // 显示图片预览
                this.$store.commit('fm/modal/setModalState', {
                    modalName: 'Preview',
                    show: true,
                });
            } else if (Object.keys(this.$store.state.fm.settings.textExtensions)
                .includes(extension.toLowerCase())) {
                // 显示文本文件
                this.$store.commit('fm/modal/setModalState', {
                    modalName: 'TextEdit',
                    show: true,
                });
            } else if (this.$store.state.fm.settings.audioExtensions
                .includes(extension.toLowerCase())) {
                // 显示音频播放器
                this.$store.commit('fm/modal/setModalState', {
                    modalName: 'AudioPlayer',
                    show: true,
                });
            } else if (this.$store.state.fm.settings.videoExtensions
                .includes(extension.toLowerCase())) {
                // 显示视频播放器
                this.$store.commit('fm/modal/setModalState', {
                    modalName: 'VideoPlayer',
                    show: true,
                });
            } else if (extension.toLowerCase() === 'pdf') {
                // 显示 pdf 文档
                this.$store.dispatch('fm/openPDF', {
                    disk: this.selectedDisk,
                    path,
                });
            }
        },
    },
};
