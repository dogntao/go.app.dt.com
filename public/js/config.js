require.config({
    baseUrl: "/js/components",
    paths: {
        "jquery": "jquery.min",
        "vue": "vue",
        "vue1": "vue1",
        "layui": "/js/components/layui/layui",
        "layui-private": "/js/private/layui-private",
        "vue-private": "/js/private/vue-private"
    },
    map: {
        "*": {
            "css": "css.min",
            "layui": "layui-private",
            "vue": "vue-private"
        },
        "layui-private": {
            "layui": "layui",
        },
        "vue-private": {
            "vue": "vue1",
        }
    },
    shim: {
        "layui": {
            deps: ["css!/js/components/layui/css/layui.css"],
            exports: "layui"
        }
    }
})