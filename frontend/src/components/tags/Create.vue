<template>
  <div class="contents">
    <header style="display: flex; display: -webkit-flex; -webkit-flex-direction: row;
   flex-direction: row;">
      <div style="flex-basis: 400px;">
        <router-link to="/" class="link">
        <h1>
        <span style="margin-right: 20px;">甘味処</span>
        <span>和</span>
        <span style="margin-left: 10px; font-size: 20px;">~なごみ~</span>
        </h1>
        </router-link>
      </div>

      <div style="align-items: center;">
        <p>Hello {{ name }}!!</p>
      </div>
    </header>
    <div>
      <h3>タグ作成</h3>
      <ul id="example-2">
        <li v-for='(tag, index) in tags' :key='index'>
          {{ tag.name }}
        </li>
      </ul>
    </div>
    <div>
      <input v-model="name" type="text" class="form-control" placeholder="タグ名">
    </div>
    <button @click="apiCreateTag">post</button>
  </div>
</template>

<script>
import axios from 'axios'
// import firebase from 'firebase'
//import EIcon from '../components/EIcon.vue';
const API_ENDPOINT = process.env.VUE_APP_BACKEND_API_BASE;

export default {
  name: 'Create',
  data() {
    return {
      name: '',
      tags: [],
    };
  },
  created () {
    this.apiTags()
  },
  methods: {
    apiTags: async function () {
      let res = await axios.get(`${API_ENDPOINT}/tags`, {
        headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
      })
      this.tags = res.data
    },

    apiCreateTag: async function () {
      const params = new URLSearchParams();
      const _this = this;
      params.append('name', this.name);

      axios.post(`${API_ENDPOINT}/tags`, params, {
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`,
          //'Content-Type': 'application/json'
          }
      })
      .then(function(res) {
          _this.apiTags();
      })
      .catch(function(error) {
          console.log(error)
      })
    },
  },
};
</script>

<style scoped>
.link {
  text-decoration: none;
  color: inherit;
}
</style>