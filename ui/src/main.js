import Vue from 'vue'
import ElementUI from 'element-ui'
import './theme/index.css'
import App from './App'
import router from './router'

if (!window.localStorage) {
  alert('不支持该浏览器，请使用主流浏览器')
}

Vue.use(ElementUI)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
