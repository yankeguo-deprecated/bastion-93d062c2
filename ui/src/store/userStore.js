export default {
  state: {
    currentUser: null
  },
  mutations: {
    setCurrentUser (state, user) {
      state.currentUser = user
    }
  },
  getters: {
    hasCurrentUser (state) {
      return !!state.currentUser
    }
  }
}
