<template>
  <div>
    <el-row class="banner">
      <img src="../assets/logo.png">
      <h1>{{ title }}</h1>
      <h2>{{ subTitle }}</h2>
    </el-row>
    <el-row>
      <el-col :md="{span:6, offset:9}">
        <el-form ref="form" :model="form" label-width="80px">
          <el-form-item label="登录名">
            <el-input placeholder="输入用户名" v-model="form.login" :disabled="submitting"></el-input>
          </el-form-item>
          <el-form-item label="密码">
            <el-input placeholder="输入密码" type="password" v-model="form.password" :disabled="submitting"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" size="large" :loading="submitting" @click="submitForm">登录</el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
    <el-row>
      <el-col class="info">
        <small>{{ info }}</small>
      </el-col>
    </el-row>
  </div>
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
    this.$api.check().then(({body}) => {
      this.info = body.name + ' v' + body.version
    }, (response) => {})
  },
  data () {
    return {
      submitting: false,
      title: '欢迎使用',
      subTitle: '请登录',
      info: null,
      form: {
        login: null,
        password: null
      }
    }
  },
  methods: {
    submitForm () {
      this.submitting = true
      this.$api.createToken(this.form).then(({body}) => {
        this.$store.commit('setToken', body.token.secret)
        this.$message.success('登录成功')
        this.$router.push({ name: 'dashboard' })
        this.submitting = false
      }, ({body}) => {
        this.$message.error(body.message)
        this.submitting = false
      })
    }
  }
}
</script>

<style scoped>
.banner {
  text-align: center;
  margin-bottom: 1rem;
}

.info {
  color: #999;
  text-align: center;
}
</style>
