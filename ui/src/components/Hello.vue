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
  </div>
</template>

<script>
// import Auth from '../lib/auth'

export default {
  name: 'index',
  head: {
    title: {
      inner: '欢迎使用'
    }
  },
  data () {
    return {
      title: '欢迎使用',
      subTitle: '请登录',
      submitting: false,
      form: {
      }
    }
  },
  methods: {
    submitForm () {
      this.submitting = true
      setTimeout(() => {
        this.submitting = false
      }, 3000)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.banner {
  text-align: center;
  margin-bottom: 1rem;
}

h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
