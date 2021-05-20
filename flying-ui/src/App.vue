
<template>
  <div id="app">

      <router-view/>
    <vue-ins-progress-bar></vue-ins-progress-bar>
  </div>
</template>
<script>

import store from "@/store";



export default {
  name: 'Home',
  components: {

  },mounted () {
    this.$insProgress.finish()
  },
  created () {
    this.$insProgress.start()

    this.$router.beforeEach((to, from, next) => {
      this.$insProgress.start()

      if (store.getters.x_token) {

        if (to.path === '/login') {
          return  next({
            name: "首页"
          })
        }else{
          return   next()
        }
      } else {
        if (to.path === '/login') {

          return  next()
        }else{

          return  next({path:'/login'})

        }
      }
    })

    this.$router.afterEach((to, from) => {
      this.$insProgress.finish()
    })
  },
  beforeCreate () {
    document.querySelector('body').setAttribute('style', 'background-color:#f0f2f5')
  },
}
</script>
<style>
#app {
  /*font-family: Avenir, Helvetica, Arial, sans-serif;*/
  /*-webkit-font-smoothing: antialiased;*/
  /*-moz-osx-font-smoothing: grayscale;*/
  /*text-align: center;*/
  /*color: #2c3e50;*/

}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

</style>

