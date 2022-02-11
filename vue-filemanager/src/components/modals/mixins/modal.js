export default {
    directives: {
        /**
         * input元素自动聚焦
         */
        focus: {
            inserted(el) {
                el.focus();
            },
        },
    },
    computed: {
        /**
         * 选择磁盘
         * @returns {String}
         */
        activeManager() {
            return this.$store.state.fm.activeManager;
        },
    },
    methods: {
        /**
         * 隐藏模块
         */
        hideModal() {
            this.$store.commit('fm/modal/setModalState', {
                modalName: null,
                show: false,
            });
        },
        handleClose(done) {
            done();
            this.hideModal();
        },
        handleCloseTip(done) {
            this.$confirm("确认关闭?", "提示", {
                    confirmButtonText: "确定",
                    cancelButtonText: "取消",
                    type: "warning"
                })
                .then(result => {
                    if (result == "confirm") {
                        done();
                        this.hideModal();
                    }
                })
                .catch(() => {});
        }
    },
}