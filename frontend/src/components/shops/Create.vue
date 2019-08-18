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
      <div style="width: 80%; margin: 0 auto;">
        <h2>ショップ作成</h2>
        <div style="width: 80%; margin: 0 auto;">
          <div class="inputWithIcon">
            <input v-model="name" type="text" placeholder="店名">
            <font-awesome-icon icon="store" />
          </div>
          <div class="inputWithIcon">
            <input v-model="address" type="text" placeholder="住所">
            <font-awesome-icon icon="map-marker" />
          </div>
          <div class="inputWithIcon">
            <input v-model="tel" type="text" placeholder="電話番号">
            <font-awesome-icon icon="phone" />
          </div>
          <div>
            <div v-for='(tag, index) in tags' :key='index'>
              <label :for='`tag_${tag.id}`'>
              <input type="checkbox" :id='`tag_${tag.id}`' :value='`${tag.id}`' v-model="tag_ids">
              {{ tag.name }}</label>
            </div>
          </div>

          <div style="width: 200px; margin: 20px auto;">
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
      </div>
    </div>
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

<style scoped>
.link {
  text-decoration: none;
  color: inherit;
}

input[type="text"] {
  width: 100%;
  border: 2px solid #aaa;
  border-radius: 4px;
  margin: 8px 0;
  outline: none;
  padding: 8px;
  box-sizing: border-box;
  transition: 0.3s;
}

input[type="text"]:focus {
  border-color: dodgerBlue;
  box-shadow: 0 0 8px 0 dodgerBlue;
}

.inputWithIcon input[type="text"] {
  padding-left: 40px;
}

.inputWithIcon {
  position: relative;
  width: 200px;
  margin: 0 auto;
}

.inputWithIcon svg {
  position: absolute;
  left: 0;
  top: 8px;
  padding: 9px 8px;
  color: #aaa;
  transition: 0.3s;
}

.inputWithIcon input[type="text"]:focus + svg {
  color: dodgerBlue;
}
</style>