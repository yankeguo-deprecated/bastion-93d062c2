import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import Dashboard from '@/components/Dashboard'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/',
      name: 'index',
      component: Hello
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: Dashboard
    }
  ]
})

router.beforeEach((to, from, next) => {
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
})

export default router
