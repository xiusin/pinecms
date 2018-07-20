"use strict";

// Theme Options
var aceOptions = {
    rtl: false,

    animations: true,

    nav: {
        height: 'auto', // use 'auto' and some fixed value with px or em 480
        arrow: false
    },

    slider: {
        obj: '',
        speed: '',
        dots: true
    }

};

// Theme Variables
var ace = {
    html: '',
    body: '',
    mobile: '',
    themeColor: '',

    header: {
        head: '',
        col1: '',
        col2: '',
        col3: ''
    },

    sidebar: {
        obj: '',
        btn: ''
    },

    nav: {
        obj: '',
        tooltip: $('<div class="ace-tooltip"></div>')
    },

    overlay: {
        obj: $('<div id="ace-overlay"></div>')
    },

    progress: {
        lines: '',
        charts: '',
        bullets: ''
    }
}

