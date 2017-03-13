require.config({
    baseUrl: "/js/components",
    paths: {
        "jquery": "jquery.min",
        "vue": "vue",
        "vue1": "vue1",
        "layui": "/js/components/layui/layui"
    },
    map: {
        "*": {
            "css": 'css.min'
        }
    },
    shim: {
        "layui": ["css!/js/components/layui/css/layui.css"]
    }
})