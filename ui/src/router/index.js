import Vue from 'vue'
import Router from 'vue-router'
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
  next()
  return
  /*
  if (to.name === 'index') {
    if (localStorage.getItem('access_token')) {
      next({name: 'dashboard'})
    } else {
      next()
    }
  } else {
    if (!localStorage.getItem('access_token')) {
      next({name: 'index'})
    } else {
      next()
    }
  }
  */
})

export default router
