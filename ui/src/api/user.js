export default {
  fetchUser: function ({id}) {
    return this.http.get(`users/${id}`).then((resp) => {
      this.transformModelDate(resp.body.user, 'createdAt')
      return resp
    })
  },
  updateUser: function ({id, nickname}) {
    return this.http.post(`users/${id}/update`, {nickname}).then((resp) => {
      this.transformModelDate(resp.body.user, 'createdAt')
      return resp
    })
  },
  updateUserPassword: function ({id, password, newPassword}) {
    return this.http.post(`users/${id}/update_password`, {password, newPassword})
  },
  listUsers: function () {
    return this.http.get('users').then((resp) => {
      this.transformModelsDate(resp.body.users, 'updatedAt')
      this.transformModelsDate(resp.body.users, 'usedAt')
      this.transformModelsDate(resp.body.users, 'createdAt')
      return resp
    })
  },
  createUser: function ({login, password}) {
    return this.http.post('users/create', {login, password})
  },
  updateAuthority: function ({id, isAdmin, isBlocked}) {
    return this.http.post(`users/${id}/update_authority`, {isAdmin, isBlocked})
  }
}
