<template>
  <b-row>
    <b-col md="3">
      <h5 class="text-info">添加用户</h5>
      <b-form @submit="createUser">
        <b-form-group label="登录名" label-for="login-input">
          <b-form-input autocomplete="off" autocorrect="off" autocapitalize="off" spellcheck="false" id="login-input" type="text" v-model="formAdd.data.login" placeholder="输入登录名"></b-form-input>
          <b-form-text text-variant="muted">最大长度24，最小长度3，创建后<b>不可修改</b>，限定<code>数字</code> <code>英文</code> <code>-</code> <code>_</code>，且只能用<code>英文</code>开头</b-form-text>
        </b-form-group>
        <b-form-group label="密码" label-for="password-input">
          <b-form-input id="password-input" type="password" v-model="formAdd.data.password" placeholder="输入密码"></b-form-input>
        </b-form-group>
        <b-form-group v-if="formAdd.error">
          <b-form-text text-variant="danger">{{formAdd.error}}</b-form-text>
        </b-form-group>
        <b-form-group>
          <b-button type="submit" :block="true" :disabled="loading" variant="info">添加</b-button>
        </b-form-group>
      </b-form>
    </b-col>
    <b-col>
      <h5 class="text-info">用户列表</h5>
      <b-table striped hover :items="users" :fields="fields">
        <template slot="id" scope="data">
          <span v-if="data.item.isBlocked" class="text-muted">{{data.item.id}}</span>
          <span v-if="!data.item.isBlocked">{{data.item.id}}</span>
        </template>
        <template slot="login" scope="data">
          <span v-if="data.item.isBlocked" class="text-muted">{{data.item.login}}</span>
          <span v-if="!data.item.isBlocked">{{data.item.login}}</span>
        </template>
        <template slot="usedAt" scope="data">
          <span v-if="data.item.isBlocked" class="text-muted">{{data.item.usedAt}}</span>
          <span v-if="!data.item.isBlocked">{{data.item.usedAt}}</span>
        </template>
        <template slot="authority" scope="data">
          <b-badge pill v-if="data.item.isMe" variant="info">我</b-badge>
          <b-badge pill v-if="data.item.isBlocked" variant="danger">已停用</b-badge>
          <b-badge pill v-if="data.item.isAdmin" variant="success">管理员</b-badge>
        </template>
        <template slot="operation" scope="data">
          <span v-if="data.item.id !== $store.state.user.currentUser.id">
            <b-link v-if="data.item.isBlocked" href="" :disabled="loading" @click="updateAuthority({id: data.item.id, isAdmin: data.item.isAdmin, isBlocked: false})" class="text-success">启用</b-link>
            <b-link v-if="!data.item.isBlocked" href="" :disabled="loading" @click="updateAuthority({id: data.item.id, isAdmin: data.item.isAdmin, isBlocked: true})" class="text-danger">停用</b-link>
            &nbsp;|&nbsp;
            <b-link v-if="data.item.isAdmin" href="" :disabled="loading" @click="updateAuthority({id: data.item.id, isAdmin: false, isBlocked: data.item.isBlocked})" class="text-danger">降权</b-link>
            <b-link v-if="!data.item.isAdmin" href="" :disabled="loading" @click="updateAuthority({id: data.item.id, isAdmin: true, isBlocked: data.item.isBlocked})" class="text-success">提权</b-link>
          </span>
        </template>
      </b-table>
    </b-col>
  </b-row>
</template>

<script>
import _ from 'lodash'

export default {
  name: 'users',
  head: {
    title: {
      inner: '用户管理'
    }
  },
  computed: {
    currentUser () {
      return this.$store.state.user.currentUser
    }
  },
  data () {
    return {
      loading: false,
      formAdd: {
        data: {},
        error: null
      },
      users: [],
      fields: {
        id: {
          label: 'ID',
          sortable: true
        },
        login: {
          label: '登录名',
          sortable: true
        },
        usedAt: {
          label: '最后使用'
        },
        authority: {
          label: ''
        },
        operation: {
          label: '操作'
        }
      }
    }
  },
  created () {
    this.reloadUsers()
  },
  methods: {
    reloadUsers () {
      this.loading = true
      this.$api.listUsers().then(({body}) => {
        this.loading = false
        this.users = _.sortBy(body.users, (u) => u.isBlocked)
        this.users.forEach((u) => {
          if (u.id === this.$store.state.user.currentUser.id) {
            u.isMe = true
          }
        })
      }, ({body}) => {
        this.loading = false
      })
    },
    createUser () {
      this.loading = true
      this.formAdd.error = null
      this.$api.createUser(this.formAdd.data).then(({body}) => {
        this.loading = false
        this.formAdd.login = null
        this.formAdd.password = null
        this.reloadUsers()
      }, ({body}) => {
        this.loading = false
        this.formAdd.error = body.message
      })
    },
    updateAuthority ({id, isAdmin, isBlocked}) {
      if (id === this.$store.state.user.currentUser.id) {
        alert('无法修改当前用户的权限')
        return
      }
      if (!confirm('确定要修改用户权限么')) {
        return
      }
      this.loading = true
      this.$api.updateAuthority({id, isAdmin, isBlocked}).then(() => {
        this.reloadUsers()
        this.loading = false
      }, () => {
        this.loading = false
      })
    }
  }
}
</script>
