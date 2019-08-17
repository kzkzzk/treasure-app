<template>
  <div class='hello'>
    <p>Hello {{ name }}!!</p>
    <h1>{{ msg }}</h1>
    <div>
      <h3>タグ一覧</h3>
      <ul id="example-2">
        <li v-for='(tag, index) in tags' :key='index'>
          {{ tag.name }}
        </li>
      </ul>
    </div>
    <div>
      <h2>ショップ一覧</h2>
      <div v-for='(shop, index) in shops' :key='index'>
        <h3>{{ shop.name }}</h3>
        <div>
          <p>{{ shop.tel }}</p>
          <p>{{ shop.address }}</p>
        </div>
      </div>
    </div>
    <button @click="signOut">Sign out</button>
    <button @click="apiUsers">public</button>
    <button @click="apiTags">private</button>
  </div>
</template>

<script>
import axios from 'axios'
import firebase from 'firebase'

const API_ENDPOINT = process.env.VUE_APP_BACKEND_API_BASE;

export default {
  name: 'Index',
  data () {
    return {
      msg: 'ショップ一覧',
      name: firebase.auth().currentUser.email,
      tags: [],
      shops: [],
    }
  },
  created () {
    this.apiTags()
    this.apiShops()
  },
  methods: {
    signOut: function () {
      firebase.auth().signOut().then(() => {
        localStorage.removeItem('jwt')
        this.$router.push('/signin')
      })
    },
    apiUsers: async function () {
      let res = await axios.get(`${API_ENDPOINT}/users`, {
        headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
      })
      this.msg = res.data
    },
    apiTags: async function () {
      let res = await axios.get(`${API_ENDPOINT}/tags`, {
        headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
      })
      this.tags = res.data
    },
    apiShops: async function () {
      let res = await axios.get(`${API_ENDPOINT}/shops`, {
        headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
      })
      this.shops = res.data
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