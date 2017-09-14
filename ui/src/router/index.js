import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import Dashboard from '@/components/Dashboard'
import Navbar from '@/components/Navbar'
import Profile from '@/components/Profile'
import Settings from '@/components/Settings'
import Security from '@/components/Security'
import SSHKeys from '@/components/SSHKeys'
import Tokens from '@/components/Tokens'
import store from '../store'

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
      components: {
        default: Dashboard,
        navbar: Navbar
      }
    },
    {
      path: '/settings',
      components: {
        default: Settings,
        navbar: Navbar
      },
      children: [
        {
          path: 'profile',
          name: 'profile',
          component: Profile
        },
        {
          path: 'security',
          name: 'security',
          component: Security
        },
        {
          path: 'ssh_keys',
          name: 'ssh_keys',
          component: SSHKeys
        },
        {
          path: 'tokens',
          name: 'tokens',
          component: Tokens
        }
      ]
    },
    {
      path: '/*',
      redirect: {
        name: 'dashboard'
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
