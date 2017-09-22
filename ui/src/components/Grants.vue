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
          <b-form-input autocomplete="off" autocorrect="off" autocapitalize="off" spellcheck="false" type="text" v-model="form.data.userLogin" required placeholder="用户登录名"></b-form-input>
        </b-form-group>
        &nbsp;
        <b-form-group>
          <b-form-input id="expires-in-input" :disabled="form.data.expiresUnit === 'i'" type="number" v-model="form.data.expiresIn"></b-form-input>
        </b-form-group>
        &nbsp;
        <b-form-group>
        <b-form-select v-model="form.data.expiresUnit" :options="form.expiresUnitOptions"></b-form-select>
        </b-form-group>
        &nbsp;
        <b-form-group>
        <b-form-select v-model="form.data.canSudo" :options="form.canSudoOptions"></b-form-select>
        </b-form-group>
        &nbsp;
        <b-button type="submit" variant="info" :disabled="$state.isLoading">创建/更新</b-button>
      </b-form>
      <br/>
      <b-table striped hover :items="grants" :fields="fields">
        <template slot="expiresAt" scope="data">
          <span v-if="!data.item.expiresAt" class="text-success">无</span>
          <span v-if="data.item.expiresAt && !data.item.isExpired" class="text-success">{{data.item.expiresAt}}</span>
          <span v-if="data.item.expiresAt && data.item.isExpired" class="text-danger">{{data.item.expiresAt}}</span>
        </template>
        <template slot="operation" scope="data">
          <b-link href="" :disabled="$state.isLoading" @click="editGrant(data.index)" class="text-info">编辑</b-link>&nbsp;|&nbsp;
          <b-link href="" :disabled="$state.isLoading" @click="destroyGrant(data.item.id)" class="text-danger">删除</b-link>
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
          canSudo: false
        },
        expiresUnitOptions: [
          { value: 'd', text: '天' },
          { value: 'h', text: '小时' },
          { value: 'm', text: '分钟' },
          { value: 'i', text: '无期限' }
        ],
        canSudoOptions: [
          { value: false, text: '无SUDO' },
          { value: true, text: 'SUDO' }
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
        },
        operation: {
          label: '操作'
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
    destroyGrant (id) {
      if (!confirm('确认要删除么')) {
        return
      }
      this.$state.begin()
      this.$api.destroyGrant({id}).then(() => {
        this.$state.end()
        this.reloadGrants()
      }, () => {
        this.$state.end()
      })
    },
    editGrant (index) {
      let item = this.grants[index]
      this.form.data.userLogin = item.userLogin
      this.form.data.canSudo = item.canSudo
      if (!item.expiresAt) {
        this.form.data.expiresUnit = 'i'
      }
    },
    createGrant () {
      let data = {
        canSudo: this.form.data.canSudo,
        tag: this.activeTag,
        userLogin: this.form.data.userLogin
      }
      if (this.form.data.expiresUnit !== 'i') {
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
        this.$state.end()
      }, () => {
        this.$state.end()
      })
    }
  }
}
</script>

<style>
input#expires-in-input {
  width: 6rem;
}
</style>
