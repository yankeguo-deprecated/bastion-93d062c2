export default {
  createToken: function ({login, password}) {
    return this.http.post('tokens/create', {login, password})
  },
  destroyToken: function ({id}) {
    return this.http.post(`tokens/${id}/destroy`)
  }
}
