<template>

  <div>
    <button v-on:click="getNewsList">点击获取</button>
    <div>{{ msg }}</div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'index',
  data() {
    return {
      msg: ""
    }
  },
  methods: {
    getNewsList() {
      axios.post('/api/v1/page', {
        "pageSize": 10,
        "pageNo": 1,
      }).then((res) => {
        this.msg = res.data
      })
    }
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}

</style>
