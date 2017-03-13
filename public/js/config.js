require.config({
    baseUrl: "/js/components",
    paths: {
        "jquery": "jquery.min",
        "vue": "vue",
        "vue1": "vue1",
        "layui": "/js/components/layui/layui",
        "layui-private": "layui-private"
    },
    map: {
        "*": {
            "css": "css.min",
            "layui": "layui-private"
        },
        "layui-private": {
            "layui": "layui"
        }
    },
    shim: {
        "layui": {
            deps: ["css!/js/components/layui/css/layui.css"],
            exports: "layui"
        }
    }
})