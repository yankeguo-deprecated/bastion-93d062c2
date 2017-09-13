function saveToken (token) {
  if (token && token.id && token.secret) {
    localStorage.setItem('token_secret', token.secret)
    localStorage.setItem('token_id', token.id)
    return true
  } else {
    localStorage.removeItem('token_secret')
    localStorage.removeItem('token_id')
    return false
  }
}

function loadToken (token) {
  const id = parseInt(localStorage.getItem('token_id'))
  const secret = localStorage.getItem('token_secret')
  if (id && secret) {
    return {id, secret}
  } else {
    return null
  }
}

export default {
  state: {
    currentToken: loadToken(),
    tokens: []
  },
  mutations: {
    setCurrentToken (state, token) {
      if (saveToken(token)) {
        state.currentToken = token
      } else {
        state.currentToken = null
      }
    },
    setTokens (state, tokens) {
      state.tokens = tokens
    }
  },
  getters: {
    isSignedIn (state) {
      return !!state.currentToken
    },
    currentTokenId (state) {
      if (state.currentToken) {
        return state.currentToken.id
      } else {
        return null
      }
    },
    currentTokenSecret (state) {
      if (state.currentToken) {
        return state.currentToken.secret
      } else {
        return null
      }
    }
  }
}
