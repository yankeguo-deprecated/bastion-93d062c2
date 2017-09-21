class State {
  constructor () {
    this.count = 0
  }
  begin () {
    this.count = this.count + 1
  }
  end () {
    this.count = this.count - 1
  }
  get isLoading () {
    return this.count > 0
  }
}

export default {
  install (Vue) {
    Object.defineProperties(Vue.prototype, {
      $state: {
        get () {
          if (!this.$_state) {
            this.$_state = new State()
          }
          return this.$_state
        }
      }
    })
  }
}
