<template>
  <b-row>
    <b-col>
      <b-row>
        <b-col :md="4">
          <h5 class="text-info">添加服务器</h5>
          <p>添加服务器前，请确认<b>堡垒主密钥</b>已经添加到目标服务器<code>root</code>用户的<code>authorized_keys</code>中</p>
        </b-col>
        <b-col>
          <h5 class="text-info">堡垒主密钥</h5>
          <p>堡垒使用以下密钥以<code>root</code>身份登录目标服务器，以执行权限管理命令。</p>
          <p><b-form-textarea :rows="8" v-model="masterPublicKey" disabled></b-form-textarea></p>
        </b-col>
      </b-row>
    </b-col>
  </b-row>
</template>

<script>
export default {
  name: 'servers',
  head: {
    title: {
      inner: '服务器管理'
    }
  },
  data () {
    return {
      loading: false,
      masterPublicKey: '正在加载'
    }
  },
  created () {
    this.updateAdminInfo()
  },
  methods: {
    updateAdminInfo () {
      this.loading = true
      this.$api.adminCheck().then(({body}) => {
        this.loading = false
        this.masterPublicKey = body.masterPublicKey
      }, ({body}) => {
        this.loading = false
      })
    }
  }
}
</script>
