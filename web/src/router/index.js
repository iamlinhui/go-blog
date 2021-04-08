import Vue from 'vue'
import Router from 'vue-router'
import getPageTitle from '@/utils/title'

Vue.use(Router)

export const constantRoutes = [
    {
        path: '/',
        component: () => import('@/views/page/index'),
        meta: {
            title: '首页'
        }
    },
    {
        path: '/404',
        component: () => import('@/views/error/404'),
        meta: {
            title: '找不到页面'
        }
    },
    // 404 page must be placed at the end !!!
    {path: '*', redirect: '/404', hidden: true}
]

const createRouter = () => new Router({
    // mode: 'history', // require service support
    scrollBehavior: () => ({y: 0}),
    routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
    const newRouter = createRouter()
    router.matcher = newRouter.matcher // reset router
}

router.beforeEach((to, from, next) => {
    /* 路由发生变化修改页面title */
    document.title = getPageTitle(to.meta.title)
    next()
})

export default router
