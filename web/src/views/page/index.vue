<template>

  <el-container>
    <el-header>Header</el-header>
    <el-container>
      <el-main>
        {{ page.data }}
        <el-pagination background layout="prev, pager, next"
                       :total="page.totalCount"
                       :page-size="page.pageSize"
                       :current-page="page.pageNo"/>
      </el-main>
      <el-aside width="200px">Aside</el-aside>
    </el-container>
    <el-footer>Footer</el-footer>

  </el-container>
</template>

<script>
import axios from 'axios'

export default {
  name: 'index',
  data() {
    return {
      page: {
        totalCount: 0,
        pageNo: 1,
        pageSize: 1,
        list: [
          {
            content: '',
            title: '',
            excerpt: '',
            url: '',
            createTime: '',
            modifyTime: '',
            commentStatus: '',
          }
        ]
      },
    }
  },
  methods: {
    getNewsList() {
      axios.post('/api/v1/page', {
        "pageSize": this.page.pageSize,
        "pageNo": this.page.pageNo,
      }).then((res) => {
        this.page = res.data
      })
    }
  }, created() {
    this.getNewsList()
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.el-header, .el-footer {
  background-color: #B3C0D1;
  color: #333;
  text-align: center;
  line-height: 60px;
}

.el-aside {
  background-color: #D3DCE6;
  color: #333;
  text-align: center;
  line-height: 200px;
}

.el-main {
  background-color: #E9EEF3;
  color: #333;
  text-align: center;
  line-height: 160px;
}

body > .el-container {
  margin-bottom: 40px;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}
</style>
