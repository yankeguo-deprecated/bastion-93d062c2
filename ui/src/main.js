import Vue from 'vue'
import VueHead from 'vue-head'
import VueResource from 'vue-resource'
import BootstrapVue from 'bootstrap-vue'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.use(BootstrapVue)
Vue.use(VueResource)
Vue.use(VueHead)

import Api from './api'

Vue.use(Api)

import App from './App'
import router from './router'
import store from './store'

/* Setup HTTP */
Vue.http.options.root = '/api'

Vue.http.interceptors.push(function (request, next) {
  if (store.getters.isSignedIn) {
    request.headers.set('Authorization', `Bearer ${store.getters.currentTokenSecret}`)
  }
  next()
})

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  store,
  router,
  render: h => h(App)
})
