import Vue from 'vue'
var accessToken = localStorage.getItem('access_token')

export default class {
  static isSignedIn () {
    return !!accessToken
  }
  static setAccessToken (at) {
    accessToken = at
    Vue.http.headers.common['Authorization'] = `Token ${at}`
  }
}
