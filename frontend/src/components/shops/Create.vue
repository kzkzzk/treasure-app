<template>
  <div class="contents">
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
</template>

<script>
//import EIcon from '../components/EIcon.vue';

export default {
  name: 'Create',
  data() {
    return {
      uploadedImage: '',
      img_name: '',
    };
  },
  methods: {
    onFileChange(e) {
      const files = e.target.files || e.dataTransfer.files;
      this.createImage(files[0]);
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
  },
};
</script>