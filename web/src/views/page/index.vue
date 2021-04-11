<template>

    <el-container>
      <Navbar/>


      <el-container>
        <el-main>
          <mavon-editor class="md" v-model="page.list[0].content" :subfield="false" :defaultOpen="'preview'" :toolbarsFlag="false" :editable="false" :scrollStyle="true" :ishljs="true"/>
          <!--        <mavon-editor v-model="page.list[0].content" :ishljs="true" :boxShadow="false" @save="save" defaultOpen="preview"/>-->
          <el-pagination background layout="prev,pager,next,jumper" :total="page.totalCount" :page-size="page.pageSize" :current-page="page.pageNo" @current-change="handleCurrentChange"/>
        </el-main>
        <el-aside width="200px">Aside</el-aside>
      </el-container>

      <el-footer>Footer</el-footer>

    </el-container>


</template>

<script>
import axios from 'axios'
import Navbar from '@/components/navbar'

export default {
  name: 'index',
  components: {Navbar},
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
      }).then((response) => {
        const res = response.data;
        if (res.code === 0) {
          this.page = res.data
        } else {
          this.$notify.error({
            title: '错误',
            message: res.msg
          });
        }
      }).catch((error) => {
        this.$notify.error({
          title: '错误',
          message: error.message
        });
      })
    },
    handleCurrentChange(val) {
      this.page.pageNo = val
      this.getNewsList()
    },
    save(val) {
      console.log(val)
    }
  }, created() {
    this.getNewsList()
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
