<template>
  <b-row>
    <b-col md="4">
      <b-row>
        <b-col md="12">
          <h5 class="text-info">个人信息</h5>
        </b-col>
        <b-col md="6">
          <label><b>ID</b></label>
          <p>{{currentUser.id}}</p>
        </b-col>
        <b-col md="6">
          <label><b>登录名</b></label>
          <p>{{currentUser.login}}</p>
        </b-col>
        <b-col md="12">
          <label><b>昵称</b></label>
          <p>{{currentUser.nickname}}</p>
        </b-col>
        <b-col md="12">
          <label><b>创建时间</b></label>
          <p>{{currentUser.createdAt}}</p>
        </b-col>
      </b-row>
      </b-col>
      <b-col md="4">
        <h5 class="text-info">修改个人信息</h5>
        <b-form @submit="updateUser">
          <b-form-group label="昵称" label-for="nickname-input">
            <b-form-input id="nickname-input" v-model="form.nickname" required placeholder="昵称最长20位"></b-form-input>
          </b-form-group>
          <b-form-group v-if="form.error">
            <b-form-text text-variant="danger">{{form.error}}</b-form-text>
          </b-form-group>
          <b-form-group>
            <b-button type="submit" :block="true" :disabled="loading" variant="info">修改</b-button>
          </b-form-group>
        </b-form>
      </b-col>
  </b-row>
</template>

<script>
export default {
  name: 'profile',
  head: {
    title: {
      inner: '个人信息'
    }
  },
  computed: {
    currentUser () {
      return this.$store.state.user.currentUser
    }
  },
  watch: {
    currentUser (n, o) {
      this.form.nickname = n.nickname
    }
  },
  created () {
    this.form.nickname = this.currentUser.nickname
  },
  data () {
    return {
      loading: false,
      form: {
        nickname: null,
        error: null
      }
    }
  },
  methods: {
    updateUser () {
      this.loading = true
      this.form.error = null
      this.$api.updateUser({
        id: 'current',
        nickname: this.form.nickname
      }).then(({body}) => {
        this.$store.commit('setCurrentUser', body.user)
        this.form.error = null
        this.loading = false
      }, ({body}) => {
        this.form.error = body.message
        this.loading = false
      })
    }
  }
}
</script>
