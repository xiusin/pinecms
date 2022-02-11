import Vue from 'vue'
import VueRouter from 'vue-router'
Vue.use(VueRouter)
const router = new VueRouter({
    mode: "history",
    // routes: [{
    //     path: '/',
    //     name: 'FileManager',
    //     component: FileManager
    // }]
    routes: [{
            path: '/',
            redirect: '/file'
        },

        {
            path: '/file',
            name: 'FileManager',
            component: () => import('../views/FileManager.vue'),
            meta: {
                title: '文件管理'
            }
        },
        {
            path: '/login',
            name: 'Login',
            component: () => import('../views/Login.vue'),
            meta: {
                title: '登录'
            }
        },
        {
            path: '/register',
            name: 'Register',
            component: () => import('../views/Register.vue'),
            meta: {
                title: '注册'
            }
        }
    ]
})

export default router;