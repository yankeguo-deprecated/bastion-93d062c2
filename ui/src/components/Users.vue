<template>
  <b-row>
    <b-modal ref="editPasswordModal" :hide-footer="true">
      <template slot="modal-title">
        <span class="text-info">修改密码</span>
      </template>
      <b-form @submit.prevent="updatePassword">
        <b-form-group label="用户名">
          <b-form-input disabled v-model="formEditPassword.data.login"></b-form-input>
        </b-form-group>
        <b-form-group label="密码">
          <b-form-input type="password" v-model="formEditPassword.data.newPassword" placeholder="密码至少六位"></b-form-input>
        </b-form-group>
        <b-form-group label="确认密码">
          <b-form-input type="password" v-model="formEditPassword.data.newPasswordConfirmation" placeholder="重复输入密码"></b-form-input>
        </b-form-group>
  <b-form-group v-if="formEditPassword.error">
    <b-form-text text-variant="danger">{{formEditPassword.error}}</b-form-text>
  </b-form-group>
  <b-form-group v-if="formEditPassword.success">
    <b-form-text text-variant="success">{{formEditPassword.success}}</b-form-text>
  </b-form-group>
        <b-form-group>
          <b-button type="submit" :block="true" :disabled="state.isLoading" variant="info">修改</b-button>
        </b-form-group>
      </b-form>
    </b-modal>
    <b-col md="3">
      <h5 class="text-info">添加用户</h5>
      <b-form @submit.prevent="createUser">
        <b-form-group label="登录名" label-for="login-input">
          <b-form-input autocomplete="off" autocorrect="off" autocapitalize="off" spellcheck="false" id="login-input" type="text" v-model="formAdd.data.login" placeholder="输入登录名"></b-form-input>
          <b-form-text text-variant="muted">最大长度24，最小长度3，创建后
            <b>不可修改</b>，限定
            <code>数字</code>
            <code>英文</code>
            <code>-</code>
            <code>_</code>，且只能用
            <code>英文</code>开头</b-form-text>
        </b-form-group>
        <b-form-group label="密码" label-for="password-input">
          <b-form-input id="password-input" type="password" v-model="formAdd.data.password" placeholder="输入密码"></b-form-input>
        </b-form-group>
        <b-form-group v-if="formAdd.error">
          <b-form-text text-variant="danger">{{formAdd.error}}</b-form-text>
        </b-form-group>
        <b-form-group>
          <b-button type="submit" :block="true" :disabled="state.isLoading" variant="info">添加</b-button>
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
            <b-link v-if="data.item.isBlocked" href="" :disabled="state.isLoading" @click="updateAuthority({id: data.item.id, isAdmin: data.item.isAdmin, isBlocked: false})" class="text-success">启用</b-link>
            <b-link v-if="!data.item.isBlocked" href="" :disabled="state.isLoading" @click="updateAuthority({id: data.item.id, isAdmin: data.item.isAdmin, isBlocked: true})" class="text-danger">停用</b-link>
            &nbsp;|&nbsp;
            <b-link v-if="data.item.isAdmin" href="" :disabled="state.isLoading" @click="updateAuthority({id: data.item.id, isAdmin: false, isBlocked: data.item.isBlocked})" class="text-danger">降权</b-link>
            <b-link v-if="!data.item.isAdmin" href="" :disabled="state.isLoading" @click="updateAuthority({id: data.item.id, isAdmin: true, isBlocked: data.item.isBlocked})" class="text-success">提权</b-link>
            &nbsp;|&nbsp;
            <b-link @click="showEditPassword(data.item)">修改密码</b-link>
          </span>
        </template>
      </b-table>
    </b-col>
  </b-row>
</template>

<script>
import _ from 'lodash'
import VueState from '../lib/vue-state'

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
      state: new VueState(),
      formAdd: {
        data: {},
        error: null
      },
      formEditPassword: {
        data: {
          login: null,
          id: null,
          password: null,
          passwordRepeat: null
        },
        success: null,
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
      this.state.begin()
      this.$api.listUsers().then(({body}) => {
        this.state.end()
        this.users = _.sortBy(body.users, (u) => u.isBlocked)
        this.users.forEach((u) => {
          if (u.id === this.$store.state.user.currentUser.id) {
            u.isMe = true
          }
        })
      }, ({body}) => {
        this.state.end()
      })
    },
    createUser () {
      this.state.begin()
      this.formAdd.error = null
      this.$api.createUser(this.formAdd.data).then(({body}) => {
        this.state.end()
        this.formAdd.login = null
        this.formAdd.password = null
        this.reloadUsers()
      }, ({body}) => {
        this.state.end()
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
      this.state.begin()
      this.$api.updateAuthority({id, isAdmin, isBlocked}).then(() => {
        this.state.end()
        this.reloadUsers()
      }, () => {
        this.state.end()
      })
    },
    showEditPassword (item) {
      this.formEditPassword.data.login = item.login
      this.formEditPassword.data.id = item.id
      this.$refs.editPasswordModal.show()
    },
    updatePassword () {
      if (this.formEditPassword.data.newPassword !== this.formEditPassword.data.newPasswordConfirmation) {
        this.formEditPassword.error = '重复密码不正确'
        return
      }
      this.state.begin()
      this.formEditPassword.success = null
      this.formEditPassword.error = null
      this.$api.updateUserPassword({
        id: this.formEditPassword.data.id,
        newPassword: this.formEditPassword.data.newPassword,
        password: ''
      }).then(() => {
        this.state.end()
        this.formEditPassword.success = '修改成功'
        setTimeout(() => {
          this.$refs.editPasswordModal.hide()
        }, 1000)
      }, ({body}) => {
        this.state.end()
        this.formEditPassword.error = body.message
      })
    }
  }
}
</script>
