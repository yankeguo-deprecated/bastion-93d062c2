import tokenApi from './token'
import userApi from './user'
import miscApi from './misc'
import sshKeyApi from './sshKey'

function createApi (http) {
  return Object.assign({ http }, tokenApi, userApi, miscApi, sshKeyApi)
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
