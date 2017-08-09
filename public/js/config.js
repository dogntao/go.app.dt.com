require.config({
    baseUrl: "/js/components",
    paths: {
        "jquery": "jquery.min",
        "vue": "vue",
        "vue1": "vue1",
        "layui": "/js/components/layui/layui",
        "select2": "/js/components/select2/select2.min",
        "layui-private": "/js/private/layui-private",
        "vue-private": "/js/private/vue-private"
    },
    map: {
        "*": {
            "css": "css.min",
            "layui": "layui-private",
            // "select2": "select2-private",
            "vue": "vue-private"
        },
        "layui-private": {
            "layui": "layui",
        },
        // "select2-private": {
        //     "select2": "select2",
        // },
        "vue-private": {
            "vue": "vue1",
        }
    },
    shim: {
        "layui": {
            deps: ["css!/js/components/layui/css/layui.css"],
            exports: "layui"
        },
        "select2": {
            deps: ["css!/js/components/select2/select2.min.css"],
            exports: "select2"
        },
    }
})