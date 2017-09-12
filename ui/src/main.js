import Vue from 'vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-default/index.css'
import 'normalize.css/normalize.css'
import App from './App'
import VueHead from 'vue-head'
import router from './router'

if (!window.localStorage) {
  alert('不支持该浏览器，请使用主流浏览器')
}

Vue.use(VueHead)
Vue.use(ElementUI)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
