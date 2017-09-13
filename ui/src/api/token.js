import UAParser from 'ua-parser-js'

export default {
  createToken: function ({login, password}) {
    const ua = UAParser()
    const desc = `Browser: ${ua.browser.name} ${ua.browser.version} / ${ua.os.name} ${ua.os.version}`
    return this.http.post('tokens/create', {login, password, desc})
  },
  destroyToken: function ({id}) {
    return this.http.post(`tokens/${id}/destroy`)
  },
  listTokens: function ({userId}) {
    return this.http.get(`users/${userId}/tokens`)
  }
}
