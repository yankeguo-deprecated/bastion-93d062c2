import Vue from 'vue'

export default class {

  static isSignedIn () {
    return !!this.token
  }

  static setToken (token) {
    this.token = token
    if (token) {
      localStorage.setItem('access_token', this.token)
    } else {
      localStorage.removeItem('access_token')
    }
    this.configureVue()
  }

  static init () {
    this.token = localStorage.getItem('access_token')
    this.configureVue()
  }

  static configureVue () {
    if (this.token) {
      Vue.http.headers.common['Authorization'] = `Token ${this.token}`
    } else {
      delete Vue.http.headers.common['Authorization']
    }
  }

}
