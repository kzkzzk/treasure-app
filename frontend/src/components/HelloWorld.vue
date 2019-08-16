<template>
  <div class='hello'>
    <h1>{{ msg }}</h1>
    <h2>Essential Links</h2>
    <button @click="apiUsers">public</button>
    <button @click="apiTags">private</button>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'HelloWorld',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  },
  methods: {
    apiUsers: async function () {
      let res = await axios.get('http://localhost:8080/users')
      this.msg = res.data
    },
    apiTags: async function () {
      let res = await axios.get('http://localhost:8080/tags')
      this.msg = res.data
    }
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
button {
  margin: 10px 0;
  padding: 10px;
}
</style>