export default {
  listAuditLogsByUser ({userId}) {
    return this.http.get(`users/${userId}/audit_logs`).then((resp) => {
      this.transformModelsDate(resp.body.auditLogs, 'createdAt')
      return resp
    })
  },
  listAuditLogs ({offset}) {
    return this.http.get(`audit_logs?offset=${offset}`).then((resp) => {
      this.transformModelsDate(resp.body.auditLogs, 'createdAt')
      this.transformModelsDate(resp.body.auditLogs, 'updatedAt')
      return resp
    })
  }
}
