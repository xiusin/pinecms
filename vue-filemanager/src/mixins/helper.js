export default {
    methods: {
        /**
         * Bytes to KB, MB, ..
         * @param bytes
         * @returns {string}
         */
        bytesToHuman(bytes) {
            const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];

            if (bytes === 0) return '0 Bytes';

            const i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)), 10);

            if (i === 0) return `${bytes} ${sizes[i]}`;

            return `${(bytes / (1024 ** i)).toFixed(1)} ${sizes[i]}`;
        },

        /**
         * 时间戳
         * @param timestamp
         * @returns {string}
         */
        timestampToDate(timestamp) {
            // 如果没有定义日期
            if (timestamp === undefined) return '-';
            const date = new Date(timestamp * 1000);
            // return date.toLocaleString(this.$store.state.fm.settings.lang);
            return `${date.getFullYear()}-${(date.getMonth()+1).toString().padStart(2,0)}-${(date.getDate()).toString().padStart(2,0)} ${(date.getHours()).toString().padStart(2,0)}:${(date.getMinutes()).toString().padStart(2,0)}`
        },

        /**
         * icon的 Mime 类型
         * @param mime
         * @returns {*}
         */
        mimeToIcon(mime) {
            const mimeTypes = {
                // image
                'image/gif': 'fa-file-image',
                'image/png': 'fa-file-image',
                'image/jpeg': 'fa-file-image',
                'image/bmp': 'fa-file-image',
                'image/webp': 'fa-file-image',
                'image/tiff': 'fa-file-image',
                'image/svg+xml': 'fa-file-image',

                // text
                'text/plain': 'fa-file-alt',

                // code
                'text/javascript': 'fa-file-code',
                'application/json': 'fa-file-code',
                'text/markdown': 'fa-file-code',
                'text/html': 'fa-file-code',
                'text/css': 'fa-file-code',

                // audio
                'audio/midi': 'fa-file-audio',
                'audio/mpeg': 'fa-file-audio',
                'audio/webm': 'fa-file-audio',
                'audio/ogg': 'fa-file-audio',
                'audio/wav': 'fa-file-audio',
                'audio/aac': 'fa-file-audio',
                'audio/x-wav': 'fa-file-audio',
                'audio/mp4': 'fa-file-audio',

                // video
                'video/webm': 'fa-file-video',
                'video/ogg': 'fa-file-video',
                'video/mpeg': 'fa-file-video',
                'video/3gpp': 'fa-file-video',
                'video/x-flv': 'fa-file-video',
                'video/mp4': 'fa-file-video',
                'video/quicktime': 'fa-file-video',
                'video/x-msvideo': 'fa-file-video',
                'video/vnd.dlna.mpeg-tts': 'fa-file-video',

                // archive
                'application/x-bzip': 'fa-file-archive',
                'application/x-bzip2': 'fa-file-archive',
                'application/x-tar': 'fa-file-archive',
                'application/gzip': 'fa-file-archive',
                'application/zip': 'fa-file-archive',
                'application/x-7z-compressed': 'fa-file-archive',
                'application/x-rar-compressed': 'fa-file-archive',

                // application
                'application/pdf': 'fa-file-pdf',
                'application/rtf': 'fa-file-word',
                'application/msword': 'fa-file-word',

                'application/vnd.ms-word': 'fa-file-word',
                'application/vnd.ms-excel': 'fa-file-excel',
                'application/vnd.ms-powerpoint': 'fa-file-powerpoint',

                'application/vnd.oasis.opendocument.text': 'fa-file-word',
                'application/vnd.oasis.opendocument.spreadsheet': 'fa-file-excel',
                'application/vnd.oasis.opendocument.presentation': 'fa-file-powerpoint',

                'application/vnd.openxmlformats-officedocument.wordprocessingml': 'fa-file-word',
                'application/vnd.openxmlformats-officedocument.spreadsheetml': 'fa-file-excel',
                'application/vnd.openxmlformats-officedocument.presentationml': 'fa-file-powerpoint',
            };

            if (mimeTypes[mime] !== undefined) {
                return mimeTypes[mime];
            }

            // 空白文件
            return 'fa-file';
        },

        /**
         * 文件扩展对应的 icon (font awesome)
         * @returns {*}
         * @param extension
         */
        extensionToIcon(extension) {
            // 文件的类型
            const extensionTypes = {
                // images color: #ff7743;
                gif: 'fa-file-image',
                png: 'fa-file-image',
                jpeg: 'fa-file-image',
                jpg: 'fa-file-image',
                bmp: 'fa-file-image',
                psd: 'fa-file-image',
                svg: 'fa-file-image',
                ico: 'fa-file-image',
                ai: 'fa-file-image',
                tif: 'fa-file-image',
                tiff: 'fa-file-image',

                // text color: #4d97ff;
                txt: 'fa-file-alt',
                json: 'fa-file-alt',
                log: 'fa-file-alt',
                ini: 'fa-file-alt',
                xml: 'fa-file-alt',
                md: 'fa-file-alt',
                env: 'fa-file-alt',

                // code color: #47ccab
                js: 'fa-file-code',
                php: 'fa-file-code',
                css: 'fa-file-code',
                cpp: 'fa-file-code',
                class: 'fa-file-code',
                h: 'fa-file-code',
                java: 'fa-file-code',
                sh: 'fa-file-code',
                swift: 'fa-file-code',

                // audio color: #8183f1;
                aif: 'fa-file-audio',
                cda: 'fa-file-audio',
                mid: 'fa-file-audio',
                mp3: 'fa-file-audio',
                mpa: 'fa-file-audio',
                ogg: 'fa-file-audio',
                wav: 'fa-file-audio',
                wma: 'fa-file-audio',

                // video color: #8183f1;
                wmv: 'fa-file-video',
                avi: 'fa-file-video',
                mpeg: 'fa-file-video',
                mpg: 'fa-file-video',
                flv: 'fa-file-video',
                mp4: 'fa-file-video',
                mkv: 'fa-file-video',
                mov: 'fa-file-video',
                ts: 'fa-file-video',
                '3gpp': 'fa-file-video',

                // archive color:#b54f13;
                zip: 'fa-file-archive',
                arj: 'fa-file-archive',
                deb: 'fa-file-archive',
                pkg: 'fa-file-archive',
                rar: 'fa-file-archive',
                rpm: 'fa-file-archive',
                '7z': 'fa-file-archive',
                'tar.gz': 'fa-file-archive',

                // application
                pdf: 'fa-file-pdf', //color: #ff5a5a;

                //color: #4d97ff;
                rtf: 'fa-file-word',
                doc: 'fa-file-word',
                docx: 'fa-file-word',
                odt: 'fa-file-word',

                // color: #63c422;
                xlr: 'fa-file-excel',
                xls: 'fa-file-excel',
                xlsx: 'fa-file-excel',

                //     color: #ff9743;
                ppt: 'fa-file-powerpoint',
                pptx: 'fa-file-powerpoint',
                pptm: 'fa-file-powerpoint',
                xps: 'fa-file-powerpoint',
                potx: 'fa-file-powerpoint',
            };

            if (extension && extensionTypes[extension.toLowerCase()] !== undefined) {
                return extensionTypes[extension.toLowerCase()];
            }

            // 空白文件
            return 'fa-file';
        },
        /**
         * 文件扩展对应的图标颜色
         * @param {Stirng} extension 
         */
        extensionToColor(extension) {
            // 文件的类型
            const extensionColors = {
                // images
                gif: '#ff7743',
                png: '#ff7743',
                jpeg: '#ff7743',
                jpg: '#ff7743',
                bmp: '#ff7743',
                psd: '#ff7743',
                svg: '#ff7743',
                ico: '#ff7743',
                ai: '#ff7743',
                tif: '#ff7743',
                tiff: '#ff7743',

                // text 
                txt: '#4d97ff',
                json: '#4d97ff',
                log: '#4d97ff',
                ini: '#4d97ff',
                xml: '#4d97ff',
                md: '#4d97ff',
                env: '#4d97ff',

                // code 
                js: '#47ccab',
                php: '#47ccab',
                css: '#47ccab',
                cpp: '#47ccab',
                class: '#47ccab',
                h: '#47ccab',
                java: '#47ccab',
                sh: '#47ccab',
                swift: '#47ccab',

                // audio 
                aif: '#8183f1',
                cda: '#8183f1',
                mid: '#8183f1',
                mp3: '#8183f1',
                mpa: '#8183f1',
                ogg: '#8183f1',
                wav: '#8183f1',
                wma: '#8183f1',

                // video 
                wmv: '#8183f1',
                avi: '#8183f1',
                mpeg: '#8183f1',
                mpg: '#8183f1',
                flv: '#8183f1',
                mp4: '#8183f1',
                mkv: '#8183f1',
                mov: '#8183f1',
                ts: '#8183f1',
                '3gpp': '#8183f1',

                // archive 
                zip: '#b54f13',
                arj: '#b54f13',
                deb: '#b54f13',
                pkg: '#b54f13',
                rar: '#b54f13',
                rpm: '#b54f13',
                '7z': '#b54f13',
                'tar.gz': '#b54f13',

                // application
                pdf: '#ff5a5a',

                rtf: '#4d97ff',
                doc: '#4d97ff',
                docx: '#4d97ff',
                odt: '#4d97ff',

                xlr: '#63c422',
                xls: '#63c422',
                xlsx: '#63c422',

                ppt: '#ff9743',
                pptx: '#ff9743',
                pptm: '#ff9743',
                xps: '#ff9743',
                potx: '#ff9743',
            };

            if (extension && extensionColors[extension.toLowerCase()] !== undefined) {
                return extensionColors[extension.toLowerCase()];
            }
            return '#b9c9d6';
        }
    },
};