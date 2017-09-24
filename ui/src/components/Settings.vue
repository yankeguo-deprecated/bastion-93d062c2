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
            <b-nav-item :to="{ name: 'security' }">密码和安全</b-nav-item>
            <b-nav-item :to="{ name: 'ssh_keys' }">SSH 公钥</b-nav-item>
            <b-nav-item :to="{ name: 'tokens' }">令牌</b-nav-item>
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
  margin-bottom: 12px;
}

.settings-content {
  padding-top: 16px;
}
</style>
