<template>
  <b-row>
    <b-col md="3">
      <b-list-group>
        <h5 class="text-info">标签</h5>
        <b-list-group-item @click="switchTag(tag)" href="" tag="a" v-for="tag in tags" :key="tag" :disabled="$state.isLoading" :active="tag === activeTag">{{tag}}</b-list-group-item>
      </b-list-group>
    </b-col>
    <b-col>
      <p class="text-info" v-if="activeTag === 'default'">
        <code>default</code>为所有服务器默认具有的标签，请谨慎使用
      </p>
      <b-form :inline="true" @submit="createGrant">
        <b-form-group>
          <b-form-input type="text" v-model="form.data.userLogin" required placeholder="用户登录名"></b-form-input>
        </b-form-group>
        &nbsp;
        <b-form-group>
          <b-form-checkbox v-model="form.data.isInfinity">无过期时间</b-form-checkbox>
        </b-form-group>
        &nbsp;
        <b-form-group>
          <b-form-input v-if="!form.data.isInfinity" type="number" v-model="form.data.expiresIn"></b-form-input>
        </b-form-group>
        <b-form-group>
        <b-form-select v-if="!form.data.isInfinity" v-model="form.data.expiresUnit" :options="form.expiresUnitOptions"></b-form-select>
        </b-form-group>
        &nbsp;
        <b-form-group>
          <b-form-checkbox v-model="form.data.canSudo">SUDO</b-form-checkbox>
        </b-form-group>
        &nbsp;
        <b-button type="submit" variant="primary">更新</b-button>
      </b-form>
      <br/>
      <b-table striped hover :items="grants" :fields="fields">
        <template slot="expiresAt" scope="data">
          <span v-if="!data.item.expiresAt" class="text-success">无</span>
          <span v-if="data.item.expiresAt && !data.item.isExpired" class="text-success">{{data.item.expiresAt}}</span>
          <span v-if="data.item.expiresAt && data.item.isExpired" class="text-danger">{{data.item.expiresAt}}</span>
        </template>
      </b-table>
    </b-col>
  </b-row>
</template>

<script>
import _ from 'lodash'

export default {
  name: 'grants',
  head: {
    title: {
      inner: '授权管理'
    }
  },
  created () {
    this.reloadTags()
    this.reloadGrants()
  },
  data () {
    return {
      form: {
        data: {
          userLogin: null,
          expiresIn: 30,
          expiresUnit: 'd',
          isInfinity: false,
          canSudo: false
        },
        expiresUnitOptions: [
          { value: 'd', text: '天' },
          { value: 'h', text: '小时' },
          { value: 'm', text: '分钟' }
        ],
        error: null
      },
      grants: [],
      activeTag: 'default',
      tags: [],
      fields: {
        id: {
          label: 'ID',
          sortable: true
        },
        userLogin: {
          label: '用户',
          sortable: true
        },
        canSudo: {
          label: 'SUDO',
          sortable: true
        },
        expiresAt: {
          label: '过期时间'
        },
        updatedAt: {
          label: '更新时间'
        }
      }
    }
  },
  methods: {
    reloadGrants () {
      this.$state.begin()
      this.$api.listGrants({tag: this.activeTag}).then(({body}) => {
        this.grants = body.grants
        this.$state.end()
      }, () => {
        this.$state.end()
      })
    },
    reloadTags () {
      this.$state.begin()
      this.$api.listTags().then(({body}) => {
        this.tags = _.sortBy(body.tags, (t) => t !== 'default')
        this.$state.end()
      }, () => {
        this.$state.end()
      })
    },
    switchTag (tag) {
      if (this.$state.isLoading) {
        return
      }
      this.activeTag = tag
      this.reloadGrants()
    },
    createGrant () {
      let data = {
        canSudo: this.form.data.canSudo,
        tag: this.activeTag,
        userLogin: this.form.data.userLogin
      }
      if (!this.form.data.isInfinity) {
        let scale = 0
        switch (this.form.data.expiresUnit) {
          case 'd': {
            scale = 24 * 3600
            break
          }
          case 'h': {
            scale = 3600
            break
          }
          case 'm': {
            scale = 60
            break
          }
        }
        data.expiresIn = this.form.data.expiresIn * scale
      }
      this.$state.begin()
      this.$api.createGrant(data).then(() => {
        this.reloadGrants()
        this.$stae.end()
      }, () => {
        this.$stae.end()
      })
    }
  }
}
</script>

<style>
</style>
