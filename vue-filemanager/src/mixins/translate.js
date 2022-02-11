export default {
    computed: {
        /**
         * 选择的译文
         * @returns {*}
         */
        lang() {
            // 选择的译文存在
            // if (Object.prototype.hasOwnProperty.call(
            //         this.$store.state.fm.settings.translations,
            //         this.$store.state.fm.settings.lang,
            //     )) {
            //     return this.$store.state.fm.settings.translations[
            //         this.$store.state.fm.settings.lang
            //     ];
            // }
            // 默认中文
            return this.$store.state.fm.settings.translations['zh-CN'];
        },
    },
};