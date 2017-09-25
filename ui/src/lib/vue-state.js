class VueState {
  constructor () {
    this.count = 0
    this.isLoading = false
  }
  begin () {
    this.count = this.count + 1
    this.isLoading = this.count > 0
  }
  end () {
    this.count = this.count - 1
    this.isLoading = this.count > 0
  }
}

export default VueState
