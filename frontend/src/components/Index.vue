<template>
  <div>
    <header style="display: flex; display: -webkit-flex; -webkit-flex-direction: row;
   flex-direction: row;">
      <div style="flex-basis: 400px;">
        <h1>
        <span style="margin-right: 20px;">甘味処</span>
        <span>和</span>
        <span style="margin-left: 10px; font-size: 20px;">~なごみ~</span>
        </h1>
      </div>

      <div style="align-items: center;">
        <p>Hello {{ name }}!!</p>
      </div>
    </header>
    <div>
      <div style="width: 80%; margin: 0 auto;">
        <h2>ショップ一覧</h2>
        <div style="width: 80%; margin: 0 auto;">
          <div v-for='(shop, index) in shops' :key='index' style="display: flex; display: -webkit-flex;" class="shop">
            <div style="flex-basis: 50%;">
              <div style="width: 70%; margin: 0 auto;">
                <img :src='`http://localhost:8080/${shop.image}`' v-if='shop.image.length' style="width: 100%;" />
              </div>
            </div>
            <div style="flex-basis: 50%; padding: 20px;">
              <ul class="tag">
                <li v-for='(tag, index) in shop.tags' :key='index'>
                  {{ tag.name }}
                </li>
              </ul>
              <div style="text-align: left;">
              <h3 style="font-size: 30px; margin-top: 0;">{{ shop.name }}</h3>
              <p>{{ shop.address }}</p>
              <p>TEL: {{ shop.tel }}</p>
              </div>
            </div>
          </div>
        </div>
        <div>
          <h3>タグ一覧</h3>
          <ul class="tag">
            <li v-for='(tag, index) in tags' :key='index'>
              {{ tag.name }}
            </li>
          </ul>
        </div>
      </div>
    </div>
    <router-link to="/shops/new">和菓子店投稿</router-link>
    <button @click="signOut" v-if='name.length'>Sign out</button>
    <button @click="apiTags" v-if='name.length'>private</button>
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
      msg: '甘味処 和',
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

.shop:nth-child(even) {
  flex-direction: row-reverse;
  margin-left: auto;
  margin-right: 0;
}

.shop {
  margin-bottom: 50px;
  width: 80%;
  margin-right: auto;
}

.tag {
  color: #fff;
  font-family: 'Raleway', sans-serif;
  text-align: left;
  margin: 0 auto 10px;
}

.tag li {
  font-weight: bold;
  font-size: 18px;
  display: inline-block;
  height: 30px;
  line-height: 30px;
  margin: 0 10px 5px 5px;
  padding: 0 10px 0 25px;
  position: relative;
  background: linear-gradient(45deg, #e4a3cd 25%, transparent 25%, transparent 75%, #e4a3cd 75%),
              linear-gradient(45deg, #e4a3cd 25%, transparent 25%, transparent 75%, #e4a3cd 75%);
  background-color: #cc7eb1;
  background-size: 6px 6px;
  background-position: 0 0, 3px 3px;
}

.tag li:after {
  background: #fff;
  border-radius: 50%;
  content: '';
  display: block;
  height: 4px;
  line-height: 30px;
  position: absolute;
  top: 13px;
  left: 7px;
  width: 4px;
}
</style>