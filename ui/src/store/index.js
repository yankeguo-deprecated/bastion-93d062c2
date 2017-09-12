import Vue from 'vue'
import Vuex from 'vuex'
import tokenStore from './tokenStore'
import userStore from './userStore'

Vue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    token: tokenStore,
    user: userStore
  }
})

export default store
