function createEmptyUser () {
  return {
    nickname: '正在加载',
    login: '正在加载',
    fingerprint: '正在加载',
    publicKey: '正在加载'
  }
}

export default {
  state: {
    currentUser: createEmptyUser()
  },
  mutations: {
    setCurrentUser (state, user) {
      if (user) {
        state.currentUser = user
      } else {
        state.currentUser = createEmptyUser()
      }
    }
  },
  getters: {
    hasCurrentUser (state) {
      return !!state.currentUser.id
    }
  }
}
