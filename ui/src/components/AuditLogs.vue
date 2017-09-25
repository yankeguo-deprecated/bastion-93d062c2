<<template>
  <b-row>
    <b-col>
      <h5 class="text-info">审核日志</h5>
      <b-pagination :disabled="state.isLoading" align="center" @input="changePage" size="sm" :total-rows="auditLogs.total" v-model="auditLogs.activePage" :per-page="auditLogs.limit">
      </b-pagination>
      <b-table striped hover :items="auditLogs.data" :fields="auditLogs.fields">
      </b-table>
      <b-pagination :disabled="state.isLoading" align="center" @input="changePage" size="sm" :total-rows="auditLogs.total" v-model="auditLogs.activePage" :per-page="auditLogs.limit">
      </b-pagination>
    </b-col>
  </b-row>
</template>
<<script>
import VueState from '../lib/vue-state'

export default {
  name: 'audit_logs',
  head: {
    title: {
      inner: '审核日志'
    }
  },
  data () {
    return {
      state: new VueState(),
      auditLogs: {
        activePage: 1,
        total: 0,
        offset: 0,
        limit: 50,
        data: [],
        fields: {
          id: {
            label: 'ID'
          },
          source: {
            label: '来源'
          },
          action: {
            label: '动作'
          },
          target: {
            label: '目标'
          },
          createdAt: {
            label: '时间'
          }
        }
      }
    }
  },
  created () {
    this.reloadAuditLogs()
  },
  methods: {
    reloadAuditLogs () {
      this.state.begin()
      this.$api.listAuditLogs({offset: this.auditLogs.offset}).then(({body}) => {
        this.auditLogs.data = body.auditLogs
        this.auditLogs.offset = body.offset
        this.auditLogs.limit = body.limit
        this.auditLogs.total = body.total
        this.state.end()
      }, ({body}) => {
        this.state.end()
      })
    },
    changePage (page) {
      this.auditLogs.offset = (this.auditLogs.activePage - 1) * this.auditLogs.limit
      this.reloadAuditLogs()
    }
  }
}
</script>
