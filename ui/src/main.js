import Vue from 'vue'
import ElementUI from 'element-ui'
import 'normalize.css/normalize.css'
import 'element-ui/lib/theme-default/index.css'
import App from './App'
import VueHead from 'vue-head'
import VueResource from 'vue-resource'
import router from './router'
import Auth from './lib/auth'

if (!window.localStorage) {
  alert('不支持该浏览器，请使用主流浏览器')
}

Vue.use(VueResource)
Vue.use(VueHead)
Vue.use(ElementUI)

Auth.init()

Vue.http.options.root = '/api'
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
