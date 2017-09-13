export default {
  state: {
    currentUser: {
      nickname: '正在加载',
      login: '正在加载',
      fingerprint: '正在加载',
      publicKey: '正在加载'
    }
  },
  mutations: {
    setCurrentUser (state, user) {
      state.currentUser = user
    }
  },
  getters: {
    hasCurrentUser (state) {
      return !!state.currentUser.id
    }
  }
}
