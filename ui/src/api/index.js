import tokenApi from './token'
import userApi from './user'
import miscApi from './misc'
import sshKeyApi from './sshKey'
import auditLogApi from './auditLog'
import serverApi from './server'
import utils from './utils'

function createApi (http) {
  return Object.assign(
    { http },
    utils,
    tokenApi,
    userApi,
    miscApi,
    sshKeyApi,
    auditLogApi,
    serverApi
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
