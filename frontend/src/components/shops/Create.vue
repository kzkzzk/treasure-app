<template>
  <div class="contents">
    <div>
      <input v-model="name" type="text" class="form-control" placeholder="店名">
    </div>
    <div>
      <input v-model="address" type="text" class="form-control" placeholder="住所">
    </div>
    <div>
      <input v-model="tel" type="tel" class="form-control" id="exampleInputEmail1" placeholder="電話番号">
    </div>
    <div>
      <div v-for='(tag, index) in tags' :key='index'>
        <label :for='`tag_${tag.id}`'>
        <input type="checkbox" :id='`tag_${tag.id}`' :value='`${tag.id}`' v-model="tag_ids">
        {{ tag.name }}</label>
      </div>
    </div>

    <div>
      <label v-show="!uploadedImage" class="input-item__label"
        >画像を選択
        <input type="file" @change="onFileChange" />
      </label>
      <div class="preview-item">
        <img
          v-show="uploadedImage"
          class="preview-item-file"
          :src="uploadedImage"
          alt=""
          style="max-width: 100%;"
        />
        <div v-show="uploadedImage" class="preview-item-btn" @click="remove">
          <span class="preview-item-name">{{ img_name }}</span>
          <span>×</span>
        </div>
      </div>
    </div>
    <button @click="apiShops">post</button>
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
      address: '',
      tel: '',
      file: null,
      uploadedImage: '',
      img_name: '',
      tags: [],
      tag_ids: [],
    };
  },
  created () {
    this.apiTags()
  },
  methods: {
    onFileChange(e) {
      const files = e.target.files || e.dataTransfer.files;
      this.createImage(files[0]);
      this.file = files[0];
      // 拡張子で分ける（※.が1つの想定です）
      const imgNameExe = files[0].name.split('.');

      // 拡張子から前
      let imgName = imgNameExe[0];

      // 拡張子から後ろ
      const imgExe = imgNameExe[1];

      // 表示したいMaxのByte数（全角10文字、半角20文字）
      const maxBytes = 20;
      const imgNameBytes = encodeURIComponent(imgName).replace(/%../g, 'x').length;

      // 画像ファイルとMax Byte数の比較
      if (imgNameBytes > maxBytes) {
        const zenkaku = imgNameBytes - imgName.length;
        if (zenkaku > 0) {
          imgName = imgName.slice(0, maxBytes / 2 - imgName.length) + '..';
        } else {
          imgName = imgName.slice(0, maxBytes - imgNameBytes) + '..';
        }
      }

      // 短くカットしたものと.と拡張子の文字列の連結
      imgName = imgName + '.' + imgExe;
      this.img_name = imgName;
    },
    // アップロードした画像を表示
    createImage(file) {
      const reader = new FileReader();
      reader.onload = e => {
        this.uploadedImage = e.target.result;
      };
      reader.readAsDataURL(file);
    },
    remove() {
      this.uploadedImage = false;
    },
    apiShops: async function () {
      let params = new FormData();
      let _this = this;
      if (this.uploadedImage) {
         params.append('image', this.file);
      }
      params.append('name', this.name);
      params.append('address', this.address);
      params.append('tel', this.tel);
      params.append('tag_ids', this.tag_ids);

      axios.post(`${API_ENDPOINT}/shops`, params, {
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`,
          'Content-Type': 'multipart/form-data'
          }
      })
      .then(function(res) {
          _this.$router.push('/');
      })
      .catch(function(error) {
          console.log("err")
          console.log(error)
      })
    },
    apiTags: async function () {
      let res = await axios.get(`${API_ENDPOINT}/tags`, {
        headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
      })
      this.tags = res.data
    },
  },
};
</script>