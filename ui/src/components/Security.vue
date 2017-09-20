<template>
  <b-row>
    <b-col md="3">
      <h5 class="text-info">修改密码</h5>
      <b-form @submit="updateUserPassword">
        <b-form-group label="旧密码" label-for="password-input">
          <b-form-input type="password" id="password-input" v-model="form.password" required placeholder="输入旧密码"></b-form-input>
        </b-form-group>
        <b-form-group label="新密码" label-for="new-password-input">
          <b-form-input type="password" id="new-password-input" v-model="form.newPassword" required placeholder="密码至少6位"></b-form-input>
        </b-form-group>
        <b-form-group label="确认新密码" label-for="new-password-c-input">
          <b-form-input type="password" id="new-password-c-input" v-model="form.newPasswordConfirmation" required placeholder="重复输入新密码"></b-form-input>
        </b-form-group>
        <b-form-group v-if="form.success">
          <b-form-text text-variant="success">{{form.success}}</b-form-text>
        </b-form-group>
        <b-form-group v-if="form.error">
          <b-form-text text-variant="danger">{{form.error}}</b-form-text>
        </b-form-group>
        <b-form-group>
          <b-button type="submit" :block="true" :disabled="loading" variant="info">修改密码</b-button>
        </b-form-group>
      </b-form>
    </b-col>
    <b-col>
      <h5 class="text-info">操作记录<small>(最近30条)</small></h5>
      <b-table striped hover :items="auditLogs.data" :fields="auditLogs.fields">
      </b-table>
    </b-col>
  </b-row>
</template>

<script>
export default {
  name: 'security',
  head: {
    title: {
      inner: '密码和安全'
    }
  },
  created () {
    this.updateAuditLogs()
  },
  data () {
    return {
      loading: false,
      form: {
        password: null,
        newPassword: null,
        newPasswordConfirmation: null,
        success: null,
        error: null
      },
      auditLogs: {
        data: [],
        fields: {
          id: {
            label: 'ID',
            sortable: true
          },
          createdAt: {
            label: '日期'
          },
          source: {
            label: '来源'
          },
          action: {
            label: '动作'
          },
          target: {
            label: '目标'
          }
        }
      }
    }
  },
  methods: {
    updateUserPassword () {
      this.loading = true
      this.form.success = null
      this.form.error = null
      this.$api.updateUserPassword({
        id: 'current',
        password: this.form.password,
        newPassword: this.form.newPassword
      }).then(() => {
        this.form.password = null
        this.form.newPassword = null
        this.form.newPasswordConfirmation = null
        this.form.success = '修改成功，建议前往“令牌”页面删除不使用的令牌'
        this.loading = false
      }, ({body}) => {
        this.form.password = null
        this.form.error = body.message
        this.loading = false
      })
    },
    updateAuditLogs () {
      this.$api.listAuditLogsByUser({userId: 'current'}).then(({body}) => {
        this.auditLogs.data = body.auditLogs
      }, ({body}) => {
      })
    }
  }
}
</script>
