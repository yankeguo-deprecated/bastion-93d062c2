export default {
  fetchUser: function ({id}) {
    return this.http.get(`users/${id}`)
  }
}
