<template>
  <b-row>
    <b-col>
      <b-row>
        <b-col :md="4">
          <h5 class="text-info">添加受控服务器</h5>
          <p>添加服务器前，请确认<b>堡垒主密钥</b>已经添加到受控服务器<code>root</code>用户的<code>authorized_keys</code>中</p>
          <b-form @submit="createServer">
            <b-form-group label="名称" label-for="name-input">
              <b-form-input id="name-input" type="text" v-model="form.data.name" placeholder="输入名称"></b-form-input>
              <b-form-text text-variant="muted">最大长度24，最小长度3，创建后<b>不可修改</b>，限定<code>数字</code> <code>英文</code> <code>-</code> <code>_</code>，且只能用<code>英文</code>开头</b-form-text>
            </b-form-group>
            <b-form-group label="地址" label-for="address-input">
              <b-form-input id="address-input" type="text" v-model="form.data.address" required placeholder="输入IP地址"></b-form-input>
              <b-form-text text-variant="muted">创建后<b>不可修改</b>，建议使用内网<code>IPv4</code>地址</b-form-text>
            </b-form-group>
            <b-form-group label="端口号" label-for="port-input">
              <b-form-input id="port-input" type="text" v-model="form.data.port" required placeholder="22"></b-form-input>
              <b-form-text text-variant="muted">创建后<b>不可修改</b>，一般为<code>22</code></b-form-text>
            </b-form-group>
            <b-form-group label="标签" label-for="tag-input">
              <b-form-input id="tag-input" type="text" v-model="form.data.tag" required placeholder="输入标签"></b-form-input>
              <b-form-text text-variant="muted">最大长度100，使用<code>英文逗号</code>分隔，每个标签限定<code>数字</code> <code>英文</code> <code>-</code> <code>_</code>，且只能用<code>英文</code>开头，所有受管服务器默认具有<code>default</code>标签</b-form-text>
            </b-form-group>
            <b-form-group label="备注" label-for="desc-input">
              <b-form-textarea id="desc-input" :rows="3" v-model="form.data.desc" placeholder="输入备注"></b-form-textarea>
              <b-form-text text-variant="muted">最大长度100</b-form-text>
            </b-form-group>
            <b-form-group v-if="form.error">
              <b-form-text text-variant="danger">{{form.error}}</b-form-text>
            </b-form-group>
            <b-form-group v-if="form.success">
              <b-form-text text-variant="success">{{form.success}}</b-form-text>
            </b-form-group>
            <b-form-group>
              <b-button type="submit" :block="true" :disabled="loading" variant="info">添加</b-button>
            </b-form-group>
          </b-form>
        </b-col>
        <b-col>
          <h5 class="text-info">堡垒主密钥</h5>
          <p>堡垒使用以下密钥以<code>root</code>身份登录受控服务器，执行管理操作。</p>
          <p><b-form-textarea :rows="10" v-model="masterPublicKey" disabled></b-form-textarea></p>
          <h5 class="text-info">受控服务器</h5>
          <b-table striped hover :items="servers" :fields="fields">
            <template slot="address" scope="data">
              {{data.item.address}}:{{data.item.port}}
            </template>
          </b-table>
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
      masterPublicKey: '正在加载',
      form: {
        data: {
          port: 22
        },
        error: null
      },
      servers: [],
      fields: {
        id: {
          label: 'ID',
          sortable: true
        },
        name: {
          label: '名称',
          sortable: true
        },
        address: {
          label: '地址'
        },
        tags: {
          label: '标签',
          formatter: 'formatTags'
        }
      }
    }
  },
  created () {
    this.updateAdminInfo()
    this.updateServers()
  },
  methods: {
    formatTags (value) {
      return value.join(',')
    },
    updateAdminInfo () {
      this.loading = true
      this.$api.adminCheck().then(({body}) => {
        this.loading = false
        this.masterPublicKey = body.masterPublicKey
      }, ({body}) => {
        this.loading = false
      })
    },
    createServer () {
      this.form.data.port = parseInt(this.form.data.port) || 22
      this.loading = true
      this.form.error = null
      this.form.success = null
      this.$api.createServer(this.form.data).then(() => {
        this.form.data = { port: 22 }
        this.form.success = '创建成功'
        this.loading = false
        this.updateServers()
      }, ({body}) => {
        this.loading = false
        this.form.error = body.message
      })
    },
    updateServers () {
      this.loading = true
      this.$api.listServers().then(({body}) => {
        this.servers = body.servers
        this.loading = false
      })
    },
    destroyServer (id) {
    }
  }
}
</script>
