<template>
  <b-row>
    <b-col>
      <b-row>
        <b-col :md="3">
          <b-row>
            <b-col>
              <h5 class="text-info">新建沙箱连接公钥</h5>
            </b-col>
          </b-row>
          <b-row>
            <b-col>
              <p>沙箱连接公钥用于从外部连接沙箱环境</p>
              <b-form @submit="createSSHKey">
                <b-form-group label="名称" label-for="name-input">
                  <b-form-input id="name-input" type="text" v-model="form.name" placeholder="输入名称，20字以内"></b-form-input>
                </b-form-group>
                <b-form-group label="公钥" label-for="public-key-input">
                  <b-form-textarea id="public-key-input" :rows="3" v-model="form.publicKey" required placeholder="ssh-rsa AAA..."></b-form-textarea>
                </b-form-group>
                <b-form-group v-if="form.error">
                  <b-form-text text-variant="danger">{{form.error}}</b-form-text>
                </b-form-group>
                <b-form-group>
                  <b-button type="submit" :block="true" :disabled="loading" variant="info">创建</b-button>
                </b-form-group>
              </b-form>
            </b-col>
          </b-row>
        </b-col>
        <b-col>
          <b-row>
            <b-col>
              <h5 class="text-info">目标公钥</h5>
              <p>目标公钥用于从沙箱连接目标服务器，位于沙箱内 <code>/root/.ssh/id_rsa.pub</code>，请勿修改</p>
              <p><code>{{ currentUser.fingerprint }}</code></p>
              <p><b-form-textarea :rows="8" v-model="currentUser.publicKey" disabled></b-form-textarea></p>
            </b-col>
          </b-row>
          <b-row>
            <b-col>
            </b-col>
          </b-row>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <h5 class="text-info">沙箱连接公钥</h5>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <b-table striped hover :items="sshKeys" :fields="fields">
            <template slot="fingerprint" scope="data">
              <code>{{data.item.fingerprint}}</code>
            </template>
            <template slot="operation" scope="data">
              <b-button :disabled="loading" size="sm" variant="danger" @click="destroySSHKey(data.item.id)">删除</b-button>
            </template>
          </b-table>
        </b-col>
      </b-row>
    </b-col>
  </b-row>
</template>

<script>
export default {
  name: 'ssh_keys',
  head: {
    title: {
      inner: 'SSH 密钥'
    }
  },
  computed: {
    currentUser () {
      return this.$store.state.user.currentUser
    }
  },
  created () {
    this.updateSSHKeys()
  },
  data () {
    return {
      loading: false,
      sshKeys: [],
      fields: {
        id: {
          label: 'ID',
          sortable: true
        },
        name: {
          label: '名称',
          sortable: true
        },
        createdAt: {
          label: '创建日期',
          sortable: true
        },
        usedAt: {
          label: '最后使用',
          sortable: true
        },
        operation: {
          label: '操作'
        }
      },
      form: {
        name: null,
        publicKey: null
      }
    }
  },
  methods: {
    updateSSHKeys () {
      this.loading = true
      this.$api.listSSHKeys({userId: 'current'}).then(({body}) => {
        this.sshKeys = body.sshKeys
        this.loading = false
      })
    },
    createSSHKey () {
      this.loading = true
      this.form.error = null
      this.$api.createSSHKey({
        userId: 'current',
        name: this.form.name,
        publicKey: this.form.publicKey
      }).then(({body}) => {
        this.loading = false
        this.form.name = null
        this.form.publicKey = null
        this.form.error = null
        this.updateSSHKeys()
      }, ({body}) => {
        this.loading = false
        this.form.error = body.message
      })
    },
    destroySSHKey (id) {
      this.loading = true
      this.$api.destroySSHKey({id}).then(() => {
        this.loading = false
        this.updateSSHKeys()
      }, () => {
        this.loading = false
        this.updateSSHKeys()
      })
    }
  }
}
</script>
