<template>
  <b-row>
    <b-col>
      <b-row>
        <b-col>
          <h5 class="text-info">当前令牌</h5>
          <p>令牌用于 API 访问授权，Web 界面使用 HTML LocalStorage 存储令牌</p>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <label><b>ID</b></label>
          <p>{{currentToken.id}}</p>
        </b-col>
        <b-col>
          <label><b>创建日期</b></label>
          <p>{{currentToken.createdAt}}</p>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <label><b>备注</b></label>
          <p>{{currentToken.desc}}</p>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <h5 class="text-info">其他令牌</h5>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <b-table striped hover :items="restTokens" :fields="fields">
            <template slot="operation" scope="data">
              <b-button :disabled="loading" size="sm" variant="danger" @click="destroyToken(data.item.id)">删除</b-button>
            </template>
          </b-table>
        </b-col>
      </b-row>
    </b-col>
  </b-row>
</template>

<script>
export default {
  name: 'tokens',
  head: {
    title: {
      inner: '令牌'
    }
  },
  computed: {
    currentToken () {
      for (let token of this.$store.state.token.tokens) {
        if (token.id === this.$store.getters.currentTokenId) {
          return token
        }
      }
      return {}
    },
    restTokens () {
      return this.$store.state.token.tokens.filter((t) => t.id !== this.$store.getters.currentTokenId)
    }
  },
  created () {
    this.updateTokens()
  },
  data () {
    return {
      loading: false,
      fields: {
        id: {
          label: 'ID',
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
        desc: {
          label: '备注'
        },
        operation: {
          label: '操作'
        }
      }
    }
  },
  methods: {
    updateTokens () {
      this.loading = true
      this.$api.listTokens({userId: 'current'}).then(({body}) => {
        this.loading = false
        this.$store.commit('setTokens', body.tokens)
      }, () => {
        this.loading = false
      })
    },
    destroyToken (id) {
      this.loading = true
      this.$api.destroyToken({id}).then(() => {
        this.loading = false
        this.updateTokens()
      }, () => {
        this.loading = false
        this.updateTokens()
      })
    }
  }
}
</script>
