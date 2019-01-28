import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/Home'

/*import Index from '~/pages/Index'
import Login from '~/pages/Login'
import Register from '~/pages/Register'
import Log from '~/pages/Log'
import MessageBoard from '~/pages/MessageBoard'
import BlogWrite from '~/pages/blog/BlogWrite'
import BlogView from '~/pages/blog/BlogView'
import BlogAllCategoryTag from '~/pages/blog/BlogAllCategoryTag'
import BlogCategoryTag from '~/pages/blog/BlogCategoryTag'*/

import {Message} from 'element-ui';
import {getToken} from '@/request/token'
import store from '@/store'
import {removeToken} from "../request/token";

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/write/:id?',
      component: r => require.ensure([], () => r(require('~/pages/blog/BlogWrite')), 'blogwrite'),
      meta: {
        requireLogin: true
      },
    },
    {
      path: '/verify/:token',
      component: r => require.ensure([], () => r(require('~/pages/Verify')), 'verify')
    },
    {
      path: '',
      name: 'Home',
      component: Home,
      children: [
        {
          path: '/about',
          component: r => require.ensure([], () => r(require('~/pages/About')), 'about')
        },
        {
          path: '/user/center',
          component: r => require.ensure([], () => r(require('~/pages/usercenter/Index')), 'usercenter'),
          meta: {
            requireLogin: true
          },
        },
        {
          path: '/user/modipwd',
          component: r => require.ensure([], () => r(require('~/pages/usercenter/Password')), 'usercenter'),
          meta: {
            requireLogin: true
          },
        },
        {
          path: '/user/money',
          component: r => require.ensure([], () => r(require('~/pages/usercenter/Money')), 'usercenter'),
          meta: {
            requireLogin: true
          },
        },
        {
          path: '/user/task',
          component: r => require.ensure([], () => r(require('~/pages/usercenter/Task')), 'usercenter'),
          meta: {
            requireLogin: true
          },
        },
        {
          path: '/user/cart',
          component: r => require.ensure([], () => r(require('~/pages/usercenter/Cart')), 'usercenter'),
          meta: {
            requireLogin: true
          },
        },
        {
          path: '/user/ref',
          component: r => require.ensure([], () => r(require('~/pages/usercenter/Ref')), 'usercenter'),
          meta: {
            requireLogin: true
          },
        },

        {
          path: '/buy/:id/:paytype',
          component: r => require.ensure([], () => r(require('~/pages/Buy')), 'buy'),
          meta: {
            requireLogin: true
          },
        },

        {
          path: '/',
          component: r => require.ensure([], () => r(require('~/pages/Index')), 'index')
        },
        {
          path: '/log',
          component: r => require.ensure([], () => r(require('~/pages/Log')), 'log')
        },
        {
          path: '/free/video/:id?',
          component: r => require.ensure([], () => r(require('~/pages/blog/FreeVideoCategory')), 'freeVideoCategory')
        },
        {
          path: '/free/book/:id?',
          component: r => require.ensure([], () => r(require('~/pages/blog/FreeBookCategory')), 'freeBookCategory')
        },
        {
          path: '/paid/video/:id?',
          component: r => require.ensure([], () => r(require('~/pages/blog/PaidVideoCategory')), 'paidVideoCategory')
        },
        {
          path: '/paid/book/:id?',
          component: r => require.ensure([], () => r(require('~/pages/blog/PaidBookCategory')), 'paidBookCategory')
        },
        {
          path: '/soft/:id?',
          component: r => require.ensure([], () => r(require('~/pages/blog/SoftCategory')), 'softCategory')
        },
        {
          path: '/archives/:year?/:month?',
          component: r => require.ensure([], () => r(require('~/pages/blog/BlogArchive')), 'archives')
        },
        {
          path: '/archives/:year?/:month?',
          component: r => require.ensure([], () => r(require('~/pages/blog/BlogArchive')), 'archives')
        },
        {
          path: '/feedback',
          component: r => require.ensure([], () => r(require('~/pages/MessageBoard')), 'messageboard')
        },
        {
          path: '/view/:id',
          component: r => require.ensure([], () => r(require('~/pages/blog/BlogView')), 'blogview')
        },
        {
          path: '/:type/all',
          component: r => require.ensure([], () => r(require('~/pages/blog/BlogAllCategoryTag')), 'blogallcategorytag')
        },
        {
          path: '/:type/:id',
          component: r => require.ensure([], () => r(require('~/pages/blog/BlogCategoryTag')), 'blogcategorytag')
        }
      ]
    },
    {
      path: '/login',
      component: r => require.ensure([], () => r(require('~/pages/Login')), 'login')
    },
    {
      path: '/register',
      component: r => require.ensure([], () => r(require('~/pages/Register')), 'register')
    }
  ],
  mode: 'history',
  scrollBehavior(to, from, savedPosition) {
    return {x: 0, y: 0}
  }
})
router.beforeEach((to, from, next) => {
  if (getToken()) { // 有token
    if (to.path === '/login') {
      next({path: '/'})
    } else {
      if (store.state.account.length === 0) {
        store.dispatch('getUserInfo').then(data => { //获取用户信息
          next()
        }).catch(() => {
          removeToken() //todo 移除token防止死循环一直跳
          next({path: '/'})
        })
      } else {
        next()
      }
    }
  } else { // 无token
    if (to.matched.some(r => r.meta.requireLogin)) {
      Message({
        type: 'warning',
        showClose: true,
        message: '请先登录账户哟!',
        onClose: () => {
          next({path: '/login'})
        }
      })
    } else {
      next();
    }
  }
})


export default router
