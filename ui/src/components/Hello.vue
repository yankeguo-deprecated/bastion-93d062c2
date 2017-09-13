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
          <b-alert variant="danger" dismissible :show="alert.danger.show" @dismissed="alert.danger.show = false">
            {{ alert.danger.message }}
          </b-alert>
        </b-col>
      </b-row>

      <b-row>
        <b-col md="4"></b-col>
        <b-col md="4">
          <b-form @submit="submitForm">
            <b-form-group label="登录名:" label-for="login-input">
              <b-form-input id="login-input" type="text" v-model="form.data.login" required placeholder="输入登录名"></b-form-input>
            </b-form-group>
            <b-form-group label="密码:" label-for="password-input">
              <b-form-input id="password-input" type="password" v-model="form.data.password" required placeholder="输入密码"></b-form-input>
            </b-form-group>
            <b-form-group>
              <b-button type="submit" :block="true" :disabled="form.submitting" variant="primary">登录</b-button>
            </b-form-group>
          </b-form>
        </b-col>
      </b-row>

      <b-row class="text-center">
        <b-col>
          <small>{{ systemInfo }}</small>
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
      systemInfo: '-',
      alert: {
        danger: {
          message: '',
          show: false
        }
      },
      form: {
        data: {
          login: null,
          password: null
        },
        submitting: false
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
      this.form.submitting = true
      this.$api.createToken(this.form.data).then(({body}) => {
        this.$store.commit('setCurrentToken', body.token)
        this.$router.push({ name: 'dashboard' })
        this.form.submitting = false
      }, ({body}) => {
        this.alert.danger.message = body.message
        this.alert.danger.show = true
        this.form.submitting = false
      })
    }
  }
}
</script>
