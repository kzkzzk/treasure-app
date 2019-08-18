<template>
    <div class="signin">
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

        <h2>Sign in</h2>
        <input type="text" placeholder="email" v-model="email">
        <input type="password" placeholder="Password" v-model="password">
        <button @click="signIn">Signin</button>
        <p>You don't have an account?
            <router-link to="/signup">create account now!!</router-link>
        </p>
    </div>
</template>

<script>
import firebase from 'firebase'
export default {
  name: 'Signin',
  data: function () {
    return {
      email: '',
      password: ''
    }
  },
    methods: {
        signIn: function () {
            //サーバーでの認証に使うJWT(res.user.qa)をローカルストレージに保管
            firebase.auth().signInWithEmailAndPassword(this.email, this.password).then(res => {
                localStorage.setItem('jwt', res.user.ra)
                this.$router.push('/')
            }, err => {
                alert(err.message)
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
.signin {
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