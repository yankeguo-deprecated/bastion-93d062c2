<template>
  <b-row>
    <b-col>
      <b-row class="settings-title">
        <b-col>
          <h3>设置</h3>
        </b-col>
      </b-row>
      <b-row class="settings-tabs">
        <b-col>
          <b-nav tabs>
            <b-nav-item :to="{ name: 'profile' }">个人信息</b-nav-item>
            <b-nav-item :to="{ name: 'change-password' }">修改密码</b-nav-item>
            <b-nav-item :to="{ name: 'ssh-keys' }">SSH 公钥</b-nav-item>
          </b-nav>
        </b-col>
      </b-row>
      <b-row class="settings-content">
        <b-col>
          <router-view></router-view>
        </b-col>
      </b-row>
    </b-col>
  </b-row>
</template>

<script>
export default {
  name: 'settings',
  data () {
    return {}
  }
}
</script>

<style scoped>
.settings-title {
  margin-top: 16px;
}

.settings-tabs {
}

.settings-content {
  padding-top: 16px;
}
</style>
