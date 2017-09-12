export default {
  createToken: function ({login, password}) {
    return this.http.post('tokens/create', {login, password})
  }
}
