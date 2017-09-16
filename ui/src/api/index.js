import tokenApi from './token'
import userApi from './user'
import miscApi from './misc'
import sshKeyApi from './sshKey'
import auditLogApi from './auditLog'
import utils from './utils'

function createApi (http) {
  return Object.assign(
    { http },
    utils,
    tokenApi,
    userApi,
    miscApi,
    sshKeyApi,
    auditLogApi
  )
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
