import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import Dashboard from '@/components/Dashboard'
import Navbar from '@/components/Navbar'
import Profile from '@/components/Profile'
import Settings from '@/components/Settings'
import ChangePassword from '@/components/ChangePassword'
import SSHKeys from '@/components/SSHKeys'
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
          path: 'change-password',
          name: 'change-password',
          component: ChangePassword
        },
        {
          path: 'ssh-keys',
          name: 'ssh-keys',
          component: SSHKeys
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
