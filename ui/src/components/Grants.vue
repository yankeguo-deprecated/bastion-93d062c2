<template>
  <b-row>
    <b-col md="3">
      <b-list-group>
        <b-list-group-item @click="switchTag(tag)" href="" tag="a" v-for="tag in tags" :key="tag" :disabled="$state.isLoading" :active="tag === activeTag">{{tag}}</b-list-group-item>
      </b-list-group>
    </b-col>
    <b-col>
      <p class="text-info" v-if="activeTag === 'default'">
        <code>default</code>为所有服务器默认具有的标签，请谨慎使用
      </p>
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
    }
  }
}
</script>

<style>
</style>
