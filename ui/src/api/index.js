import tokenApi from './token'
import userApi from './user'
import miscApi from './misc'

function createApi (http) {
  return Object.assign({ http }, tokenApi, userApi, miscApi)
}

export default {
  install (Vue) {
    Object.defineProperties(Vue.prototype, {
      $api: {
        get () {
          return createApi(this.$http)
        }
      }
    })
  }
}
