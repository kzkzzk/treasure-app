import Vue from 'vue'
import App from './App.vue'
import router from './router'
// import store from './store'
import firebase from 'firebase'

Vue.config.productionTip = false

let app

const config = {
  apiKey: 'AIzaSyADyZwHaQeIp-qPneOyLVhXjKKYj9hzN94',
  authDomain: 'treasure-cf395.firebaseapp.com',
  databaseURL: 'https://treasure-cf395.firebaseio.com',
  projectId: 'treasure-cf395',
  storageBucket: '',
  messagingSenderId: '754038262219',
  appId: "1:754038262219:web:36ecef7cc1b7f85f"
}
firebase.initializeApp(config)


firebase.auth().onAuthStateChanged(user => {
  /* eslint-disable no-new */
  if (!app) {
    new Vue({
      router,
      // store,
      render: h => h(App),
    }).$mount('#app')
  }
})

// new Vue({
//   router,
//   // store,
//   render: h => h(App),
// }).$mount('#app')
