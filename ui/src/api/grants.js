export default {
  listGrants ({tag}) {
    return this.http.get(`tags/${tag}/grants`).then((resp) => {
      this.transformModelsDate(resp.body.grants, 'createdAt')
      this.transformModelsDate(resp.body.grants, 'updatedAt')
      return resp
    })
  }
}
