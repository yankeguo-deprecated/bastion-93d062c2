export default {
  check () {
    return this.http.get('')
  },
  listTags () {
    return this.http.get('tags')
  }
}
