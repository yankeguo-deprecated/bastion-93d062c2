export default {
  listServers () {
    return this.http.get('servers').then((resp) => {
      this.transformModelsDate(resp.body.servers, 'createdAt')
      return resp
    })
  },
  createServer ({name, address, port, tag, desc}) {
    return this.http.post('servers/create', {name, address, port, tag, desc})
  },
  updateServer ({id, tag, desc}) {
    return this.http.post(`servers/${id}/update`, {tag, desc})
  },
  destroyServer ({id}) {
    return this.http.post(`servers/${id}/destroy`)
  }
}
