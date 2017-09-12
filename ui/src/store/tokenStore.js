export default {
  state: {
    token: localStorage.getItem('access_token')
  },
  mutations: {
    setToken (state, token) {
      if (token) {
        localStorage.setItem('access_token', token)
      } else {
        localStorage.removeItem('access_token')
      }
      state.token = token
    }
  },
  getters: {
    isSignedIn (state) {
      return !!state.token
    }
  }
}
