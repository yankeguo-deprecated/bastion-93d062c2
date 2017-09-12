import Vue from 'vue'
import Router from 'vue-router'
import Auth from '../lib/auth'
import Hello from '@/components/Hello'
import Dashboard from '@/components/Dashboard'
import Navbar from '@/components/Navbar'
import Sidebar from '@/components/Sidebar'

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
    }
  ]
})

router.beforeEach((to, from, next) => {
  if (to.name === 'index') {
    if (Auth.isSignedIn()) {
      next({name: 'dashboard'})
    } else {
      next()
    }
  } else {
    if (!Auth.isSignedIn()) {
      next({name: 'index'})
    } else {
      next()
    }
  }
})

export default router
