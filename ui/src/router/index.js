import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import Dashboard from '@/components/Dashboard'
import Navbar from '@/components/Navbar'
import Sidebar from '@/components/Sidebar'
import Profile from '@/components/Profile'
import store from '../store'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      // hello, i.e. Login
      path: '/',
      name: 'index',
      component: Hello,
      meta: {
        hidesNavbar: true,
        hidesSidebar: true
      }
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      components: {
        default: Dashboard,
        navbar: Navbar,
        sidebar: Sidebar
      }
    },
    {
      path: '/profile',
      name: 'profile',
      components: {
        default: Profile,
        navbar: Navbar,
        sidebar: Sidebar
      }
    }
  ]
})

router.beforeEach((to, from, next) => {
  if (to.name === 'index') {
    if (store.getters.isSignedIn) {
      next({name: 'dashboard'})
    } else {
      next()
    }
  } else {
    if (!store.getters.isSignedIn) {
      next({name: 'index'})
    } else {
      next()
    }
  }
})

export default router
