export default {
  fetchCurrentUser: function () {
    return this.http.get('users/current')
  }
}
