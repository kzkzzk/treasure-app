<template>
  <div class="signup">
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
    <h2>Sign up</h2>
    <input type="text" placeholder="email" v-model="email">
    <input type="password" placeholder="Password" v-model="password">
    <button @click="signUp">Register</button>
    <p>Do you have an account?
      <router-link to="/signin">sign in now!!</router-link>
    </p>
  </div>
</template>

<script>
import firebase from 'firebase'
export default {
  name: 'Signup',
  data () {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    signUp: function () {
      firebase.auth().createUserWithEmailAndPassword(this.email, this.password).then(res => {
        console.log('Create account: ', res.user.email)
      }).catch(error => {
        console.log(error.message)
      })
    }
  }
}
</script>

<style scoped>
.link {
  text-decoration: none;
  color: inherit;
}

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
.signup {
  margin-top: 20px;
  display: flex;
  flex-flow: column nowrap;
  justify-content: center;
  align-items: center
}
input {
  margin: 10px 0;
  padding: 10px;
}
button {
  margin: 10px 0;
  padding: 10px;
}
</style>