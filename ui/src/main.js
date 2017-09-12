import Vue from 'vue'
import ElementUI from 'element-ui'
import 'normalize.css/normalize.css'
import 'element-ui/lib/theme-default/index.css'
import VueHead from 'vue-head'
import VueResource from 'vue-resource'
import Api from './api'

Vue.use(VueResource)
Vue.use(VueHead)
Vue.use(ElementUI)
Vue.use(Api)

import App from './App'
import router from './router'
import store from './store'

/* Setup HTTP */
Vue.http.options.root = '/api'

Vue.http.interceptors.push(function (request, next) {
  if (store.getters.isSignedIn) {
    request.headers.set('Authorization', `Bearer ${store.state.token.token}`)
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
