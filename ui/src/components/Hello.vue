<template>
  <div>
    <el-row class="banner">
      <img src="../assets/logo.png">
      <h1>{{ msg }}</h1>
    </el-row>
    <el-row>
      <el-col :md="{span:6, offset:9}">
        <el-form ref="form" :model="form" label-width="80px">
          <el-form-item label="登录名">
            <el-input placeholder="输入用户名" v-model="form.login"></el-input>
          </el-form-item>
          <el-form-item label="密码">
            <el-input placeholder="输入密码" type="password" v-model="form.password"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" size="large">登录</el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'index',
  data () {
    return {
      msg: '欢迎使用 Bastion',
      form: {
      }
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
