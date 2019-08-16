import Vue from 'vue'
import App from './App.vue'
import router from './router'
// import store from './store'
import firebase from 'firebase'

Vue.config.productionTip = false

let app

const config = {
  apiKey: process.env.VUE_APP_FIREBASE_APIKEY,
  authDomain: process.env.VUE_APP_FIREBASE_AUTHDOMAIN,
  databaseURL: process.env.VUE_APP_FIREBASE_DATABASEURL,
  projectId: process.env.VUE_APP_FIREBASE_PROJECTID,
  messagingSenderId: process.env.VUE_APP_FIREBASE_MESSAGINGSENDERID,
  appId: process.env.VUE_APP_FIREBASE_APPID
}
firebase.initializeApp(config)


firebase.auth().onAuthStateChanged(user => {
  /* eslint-disable */
  if (!app) {
    new Vue({
      router,
      // store,
      render: h => h(App),
    }).$mount('#app')
  }
})
