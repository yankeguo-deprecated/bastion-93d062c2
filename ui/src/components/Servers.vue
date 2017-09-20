<template>
  <b-row>
    <b-modal ref="modalAdd" :hide-footer="true">
      <template slot="modal-title">
        <span class="text-info">添加受管服务器</span>
      </template>
      <b-form @submit="createServer">
        <b-form-group label="名称" label-for="name-input">
          <b-form-input id="name-input" type="text" v-model="formAdd.data.name" placeholder="输入名称"></b-form-input>
          <b-form-text text-variant="muted">最大长度24，最小长度3，创建后<b>不可修改</b>，限定<code>数字</code> <code>英文</code> <code>-</code> <code>_</code>，且只能用<code>英文</code>开头</b-form-text>
        </b-form-group>
        <b-form-group label="地址" label-for="address-input">
          <b-form-input id="address-input" type="text" v-model="formAdd.data.address" required placeholder="输入IP地址"></b-form-input>
          <b-form-text text-variant="muted">创建后<b>不可修改</b>，建议使用内网<code>IPv4</code>地址</b-form-text>
        </b-form-group>
        <b-form-group label="端口号" label-for="port-input">
          <b-form-input id="port-input" type="text" v-model="formAdd.data.port" required placeholder="22"></b-form-input>
          <b-form-text text-variant="muted">创建后<b>不可修改</b>，一般为<code>22</code></b-form-text>
        </b-form-group>
        <b-form-group label="标签" label-for="tag-input">
          <b-form-input id="tag-input" type="text" v-model="formAdd.data.tag" required placeholder="输入标签"></b-form-input>
          <b-form-text text-variant="muted">最大长度100，使用<code>英文逗号</code>分隔，每个标签限定<code>数字</code> <code>英文</code> <code>-</code> <code>_</code>，且只能用<code>英文</code>开头，所有受管服务器默认具有<code>default</code>标签</b-form-text>
        </b-form-group>
        <b-form-group label="备注" label-for="desc-input">
          <b-form-textarea id="desc-input" :rows="3" v-model="formAdd.data.desc" placeholder="输入备注"></b-form-textarea>
          <b-form-text text-variant="muted">最大长度100</b-form-text>
        </b-form-group>
        <b-form-group v-if="formAdd.error">
          <b-form-text text-variant="danger">{{formAdd.error}}</b-form-text>
        </b-form-group>
        <b-form-group>
          <b-button type="submit" :block="true" :disabled="loading" variant="info">添加</b-button>
        </b-form-group>
      </b-form>
    </b-modal>
    <b-modal ref="modalEdit" :hide-footer="true">
      <template slot="modal-title">
        <span class="text-info">编辑受管服务器</span>
      </template>
      <b-form @submit="updateServer">
        <b-form-group label="名称" label-for="name-input">
          <b-form-input id="name-input" type="text" disabled v-model="formEdit.data.name"></b-form-input>
        </b-form-group>
        <b-form-group label="地址" label-for="address-input">
          <b-form-input id="address-input" type="text" disabled v-model="formEdit.data.address"></b-form-input>
        </b-form-group>
        <b-form-group label="端口号" label-for="port-input">
          <b-form-input id="port-input" type="text" disabled v-model="formEdit.data.port"></b-form-input>
        </b-form-group>
        <b-form-group label="访问密钥" label-for="token-input">
          <b-form-input id="token-input" type="text" disabled v-model="formEdit.data.token"></b-form-input>
          <b-form-text text-variant="muted">访问密钥用于从受管服务器访问堡垒API</b-form-text>
        </b-form-group>
        <b-form-group label="标签" label-for="tag-input">
          <b-form-input id="tag-input" type="text" v-model="formEdit.data.tag" required placeholder="输入标签"></b-form-input>
          <b-form-text text-variant="muted">最大长度100，使用<code>英文逗号</code>分隔，每个标签限定<code>数字</code> <code>英文</code> <code>-</code> <code>_</code>，且只能用<code>英文</code>开头，所有受管服务器默认具有<code>default</code>标签</b-form-text>
        </b-form-group>
        <b-form-group label="备注" label-for="desc-input">
          <b-form-textarea id="desc-input" :rows="3" v-model="formEdit.data.desc" placeholder="输入备注"></b-form-textarea>
          <b-form-text text-variant="muted">最大长度100</b-form-text>
        </b-form-group>
        <b-form-group v-if="formEdit.error">
          <b-form-text text-variant="danger">{{formEdit.error}}</b-form-text>
        </b-form-group>
        <b-form-group>
          <b-button type="submit" :block="true" :disabled="loading" variant="info">保存</b-button>
        </b-form-group>
      </b-form>
    </b-modal>
    <b-col>
      <b-row>
        <b-col :md="3">
          <h5 class="text-info">添加受管服务器</h5>
          <p>添加服务器后，在右边<code>详情</code>按钮查看<code>访问密钥</code>，并参照文档在受管服务器上建立守护进程</p>
          <b-button :block="true" :disabled="loading" variant="info" @click="showModalAdd">添加</b-button>
        </b-col>
        <b-col>
          <h5 class="text-info">受管服务器</h5>
          <p>标签默认包含<code>default</code></p>
          <b-table striped hover :items="servers" :fields="fields">
            <template slot="tags" scope="data">
              <b-badge pill v-for="t in data.item.tags" variant="info">{{t}}</b-badge>
            </template>
            <template slot="operation" scope="data">
              <b-button :disabled="loading" size="sm" variant="info" @click="showModalEdit(data.index)">详情</b-button>
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
      formAdd: {
        data: {
          port: 22
        },
        error: null
      },
      formEdit: {
        data: {
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
          label: '标签'
        },
        operation: {
          label: '操作'
        }
      }
    }
  },
  created () {
    this.reloadAdminInfo()
    this.reloadServers()
  },
  methods: {
    formatTags (value) {
      return value.filter((k) => k !== 'default').join(',')
    },
    reloadAdminInfo () {
      this.loading = true
      this.$api.adminCheck().then(({body}) => {
        this.loading = false
        this.masterPublicKey = body.masterPublicKey
      }, ({body}) => {
        this.loading = false
      })
    },
    createServer () {
      this.formAdd.data.port = parseInt(this.formAdd.data.port) || 22
      this.loading = true
      this.formAdd.error = null
      this.$api.createServer(this.formAdd.data).then(() => {
        this.formAdd.data = { port: 22 }
        this.loading = false
        this.$refs.modalAdd.hide()
        this.reloadServers()
      }, ({body}) => {
        this.loading = false
        this.formAdd.error = body.message
      })
    },
    reloadServers () {
      this.loading = true
      this.$api.listServers().then(({body}) => {
        body.servers.forEach(function (s) {
          s.tags = s.tags.filter((t) => t !== 'default')
        })
        this.servers = body.servers
        this.loading = false
      })
    },
    updateServer () {
      this.loading = true
      this.formEdit.error = null
      this.$api.updateServer(this.formEdit.data).then(({body}) => {
        this.loading = false
        this.$refs.modalEdit.hide()
        const server = body.server
        for (let s of this.servers) {
          if (s.id === server.id) {
            s.tags = server.tags
            s.desc = server.desc
          }
        }
      }, ({body}) => {
        this.loading = false
        this.formEdit.error = body.message
      })
    },
    destroyServer (id) {
    },
    showModalAdd () {
      this.$refs.modalAdd.show()
    },
    showModalEdit (index) {
      this.formEdit.data = this.servers[index]
      this.formEdit.data.tag = this.formEdit.data.tags.filter((e) => e !== 'default').join(',')
      this.$refs.modalEdit.show()
    }
  }
}
</script>

<style scoped>
span.badge {
  margin-left: 0.1rem;
  margin-right: 0.1rem;
}
</style>
