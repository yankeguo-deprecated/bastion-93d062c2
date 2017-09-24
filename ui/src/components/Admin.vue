<template>
  <b-row>
    <b-col>
      <b-row class="admin-title">
        <b-col>
          <h3>系统管理</h3>
        </b-col>
      </b-row>
      <b-row class="admin-tabs">
        <b-col>
          <b-nav tabs>
            <b-nav-item :to="{ name: 'servers' }">服务器</b-nav-item>
            <b-nav-item :to="{ name: 'users' }">用户</b-nav-item>
            <b-nav-item :to="{ name: 'grants' }">授权</b-nav-item>
            <b-nav-item :to="{ name: 'audit_logs' }">审核日志</b-nav-item>
          </b-nav>
        </b-col>
      </b-row>
      <b-row class="admin-content">
        <b-col>
          <router-view></router-view>
        </b-col>
      </b-row>
    </b-col>
  </b-row>
</template>

<script>
export default {
  name: 'admin',
  data () {
    return {}
  }
}
</script>

<style scoped>
.admin-title {
  margin-top: 16px;
  margin-bottom: 12px;
}

.admin-content {
  padding-top: 16px;
}
</style>
