<template>
  <b-row>
    <b-col>
      <b-row class="dashboard-title">
        <b-col>
          <h3>工作台</h3>
        </b-col>
      </b-row>
    </b-col>
  </b-row>
</template>

<script>
export default {
  name: 'dashboard',
  head: {
    title: {
      inner: '工作台'
    }
  },
  data () {
    return {
      msg: '欢迎使用 Bastion',
      form: {
      }
    }
  }
}
</script>

<style scoped>
.dashboard-title {
  margin-top: 16px;
}
</style>
