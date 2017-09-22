<template>
  <b-row>
    <b-col>

      <b-row class="text-center">
        <b-col>
          <p>
            <img src="../assets/logo.png">
          </p>
          <p>
            <h1>欢迎使用</h1>
          </p>
        </b-col>
      </b-row>

      <b-row>
        <b-col md="4"></b-col>
        <b-col md="4">
          <b-form @submit="submitForm">
            <b-form-group label="登录名:" label-for="login-input">
              <b-form-input autocomplete="off" autocorrect="off" autocapitalize="off" spellcheck="false" id="login-input" type="text" v-model="form.login" required placeholder="输入登录名"></b-form-input>
            </b-form-group>
            <b-form-group label="密码:" label-for="password-input">
              <b-form-input id="password-input" type="password" v-model="form.password" required placeholder="输入密码"></b-form-input>
            </b-form-group>
            <b-form-group v-if="form.error">
              <b-form-text text-variant="danger">{{form.error}}</b-form-text>
            </b-form-group>
            <b-form-group>
              <b-button type="submit" :block="true" :disabled="loading" variant="primary">登录</b-button>
            </b-form-group>
            <b-form-group>
              <b-form-text class="text-center">{{ systemInfo }}</b-form-text>
            </b-form-group>
          </b-form>
        </b-col>
      </b-row>

    </b-col>
  </b-row>
</template>

<script>
export default {
  name: 'index',
  head: {
    title: {
      inner: '欢迎使用'
    }
  },
  created () {
    this.updateSystemInfo()
  },
  data () {
    return {
      loading: false,
      systemInfo: '-',
      form: {
        login: null,
        password: null
      }
    }
  },
  methods: {
    updateSystemInfo () {
      this.$api.check().then(({body}) => {
        const {name, version} = body
        this.systemInfo = name + ' v' + version
      }, (response) => {})
    },
    submitForm () {
      this.loading = true
      this.$api.createToken({
        login: this.form.login,
        password: this.form.password
      }).then(({body}) => {
        this.$store.commit('setCurrentToken', body.token)
        this.$router.push({ name: 'dashboard' })
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
