export default {
  listAuditLogsByUser ({userId}) {
    return this.http.get(`users/${userId}/audit_logs`).then((resp) => {
      this.transformModelsDate(resp.body.auditLogs, 'createdAt')
      return resp
    })
  }
}
