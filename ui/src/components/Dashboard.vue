<template>
  <b-row>
    <b-col>
      <b-row class="dashboard-title">
        <b-col>
          <h3>工作台</h3>
        </b-col>
      </b-row>
      <b-row>
        <b-col :md="12">
          <h5 class="text-info">连接沙箱</h5>
          <p v-if="sandbox.isKeyMissing" class="text-danger">没有配置沙箱 SSH 公钥，
            <router-link :to="{name: 'ssh_keys'}">去配置 >></router-link>
          </p>
          <p v-if="!sandbox.isKeyMissing">使用
            <code>ssh</code>访问
            <code>{{sandbox.address}}</code>
          </p>
        </b-col>
        <b-col>
          <b-row>
            <b-col>
              <h5 class="text-info">受管服务器</h5>
              <p>在沙箱内，使用
                <code>ssh</code>访问受管服务器</p>
            </b-col>
          </b-row>
          <b-row>
            <b-col>
              <b-table hover striped :items="servers" :fields="fields">
                <template slot="address" scope="data">{{data.item.address}}:{{data.item.port}}</template>
                <template slot="canSudo" scope="data"><b-badge variant="success" pill v-if="data.item.canSudo">SUDO</b-badge></template>
                <template slot="tags" scope="data"><b-badge variant="info" pill v-for="tag in data.item.tags" :key="tag">{{tag}}</b-badge></template>
  </b-table>
            </b-col>
          </b-row>
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
  created () {
    this.reloadDashboard()
  },
  data () {
    return {
      sandbox: {
        address: null,
        isKeyMissing: false,
        ssh: null
      },
      servers: [],
      fields: {
        id: {
          label: 'ID'
        },
        name: {
          label: '名称'
        },
        account: {
          label: '用户'
        },
        address: {
          label: '地址'
        },
        canSudo: {
          label: ' '
        },
        tags: {
          label: '标签'
        }
      }
    }
  },
  methods: {
    reloadDashboard () {
      this.$api.fetchDashboard().then(({body}) => {
        let dashboard = body.dashboard
        // sandbox info
        let address = dashboard.sandbox.address.split(':')
        if (address.length === 2) {
          this.sandbox.ssh = `ssh -p ${address[1]} ${address[0]}`
        } else {
          this.sandbox.ssh = `ssh ${address[0]}`
        }
        this.sandbox.address = dashboard.sandbox.address
        this.sandbox.isKeyMissing = dashboard.sandbox.isKeyMissing
        // servers
        dashboard.servers.forEach((s) => {
          if (s.port === 22) {
            s.ssh = `ssh ${s.account}@${s.address}`
          } else {
            s.ssh = `ssh -p ${s.port} ${s.account}@${s.address}`
          }
        })
        this.servers = dashboard.servers
      })
    }
  }
}
</script>

<style scoped>
.dashboard-title {
  margin-top: 16px;
  margin-bottom: 12px;
}
.card-body {
  padding: 0.5rem;
}
</style>
