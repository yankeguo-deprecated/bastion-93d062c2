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
  }
}
