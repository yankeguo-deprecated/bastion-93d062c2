export default {
  listAuditLogsByUser ({userId}) {
    return this.http.get(`users/${userId}/audit_logs`)
  }
}
