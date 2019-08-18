<template>
  <div class="contents">
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