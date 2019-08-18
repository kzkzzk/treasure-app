import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import Signup from '@/components/Signup'
import Signin from '@/components/Signin'
import ShopCreate from '@/components/shops/Create'
import TagCreate from '@/components/tags/Create'
import firebase from 'firebase'
 
Vue.use(Router)
 
let router = new Router({
  mode: 'history',
//   base: process.env.BASE_URL,
  routes: [
    {
      path: '*',
      redirect: 'signin'
    },
    {
      path: '/',
      name: 'Index',
      component: Index,
      meta: { requiresAuth: true } //routeに認証が必要かを判断
    },
    {
      path: '/signup',
      name: 'Signup',
      component: Signup
    },
    {
      path: '/signin',
      name: 'Signin',
      component: Signin
    },
    {
      path: '/shops/new',
      name: 'Create',
      component: ShopCreate,
      meta: { requiresAuth: true } //routeに認証が必要かを判断
    },
    {
      path: '/tags/new',
      name: 'Create',
      component: TagCreate,
      meta: { requiresAuth: true, isAdmin: true } //routeに認証が必要かを判断
    },
  ]
})

// router.beforeEach()を追加
router.beforeEach((to, from, next) => {
    let currentUser = firebase.auth().currentUser
    let requiresAuth = to.matched.some(record => record.meta.requiresAuth)
    let isAdmin = to.matched.some(record => record.meta.isAdmin)
    if (requiresAuth && !currentUser) next('signin')
    else if (!requiresAuth && currentUser) next()
    // else next()
    else if (currentUser) {
      currentUser.getIdTokenResult()
      .then((idTokenResult) => {
         // Confirm the user is an Admin.
         if (!!idTokenResult.claims.admin) {
           // Show admin UI.
            console.log("admin");
            next()
         } else if (isAdmin) {
          console.log("not permitted");
          next('index');
         } else {
           // Show regular user UI.
           console.log("regular");
           next()
         }
      })
      .catch((error) => {
        console.log(error);
        next('index');
      });
    } else {
      next()
    }
  })
  
  export default router
  