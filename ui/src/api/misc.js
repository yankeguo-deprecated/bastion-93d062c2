export default {
  check () {
    return this.http.get('')
  },
  adminCheck () {
    return this.http.get('admin')
  }
}
