export default {
  listGrants ({tag}) {
    return this.http.get(`tags/${tag}/grants`).then((resp) => {
      this.transformModelsDate(resp.body.grants, 'createdAt')
      this.transformModelsDate(resp.body.grants, 'updatedAt')
      this.transformModelsDate(resp.body.grants, 'expiresAt')
      return resp
    })
  },
  createGrant ({userLogin, canSudo, expiresIn, tag}) {
    return this.http.post(`grants/create`, {userLogin, canSudo, expiresIn, tag})
  },
  destroyGrant ({id}) {
    return this.http.post(`grants/${id}/destroy`)
  }
}
