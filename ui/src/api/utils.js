import dateformat from 'dateformat'

export default {

  transformDateTime (time, offset) {
    const d = new Date(time)
    d.setSeconds(d.getSeconds() - offset)
    return dateformat(d, 'yyyy-mm-dd HH:MM:ss')
  },

  transformModelDate (model, key) {
    const offset = new Date().getTimezoneOffset()
    model[key] = this.transformDateTime(model[key], offset)
  },

  transformModelsDate (models, key) {
    const offset = new Date().getTimezoneOffset()
    models.forEach((m) => {
      m[key] = this.transformDateTime(m[key], offset)
    })
  }

}
