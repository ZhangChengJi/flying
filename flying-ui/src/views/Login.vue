<template>
  <dev id="login" >

    <div class="bg">
      <a-form
      id="components-form-demo-normal-login"
      :form="form"
      class="login-form"
      @submit="handleLogin"
  >
        <div class="panel panel-login">
          <div class="panel-body">
   <dev class="login-img">  <img style="height:90px;margin: 0 auto" src="../image/flying3.png" /></dev>
      <a-form-item>
      <a-input
          v-decorator="[
          'userName',
          { rules: [{ required: true, message: $t('Login.pleaseInputUsername') }] },
        ]"
          :placeholder="this.$t('Login.username')"
          size="large"
      >
        <a-icon slot="prefix" type="user" style="color: rgba(0,0,0,.25)" />
      </a-input>
    </a-form-item>
    <a-form-item>
      <a-input
          v-decorator="[
          'password',
          { rules: [{ required: true, message: $t('Login.pleaseInputPassword') }] },
        ]"
          type="password"
          :placeholder="$t('Login.password')"
          size="large"
      >
        <a-icon slot="prefix" type="lock" style="color: rgba(0,0,0,.25)" />
      </a-input>
    </a-form-item>
    <a-form-item>

      <a-button type="primary" html-type="submit" class="login-form-button" size="large">
        {{$t('Login.login')}}
      </a-button>
    </a-form-item>
          </div>
        </div>
  </a-form>
    </div>

  </dev>
</template>

<script>




export default {
  data() {
    return {

    }
    },

  beforeCreate() {
    this.form = this.$form.createForm(this );
  },
  methods: {
    handleLogin(e) {
      e.preventDefault();
      this.form.validateFields((err, values) => {
        if (!err) {
          this.$store.dispatch("LoginIn", values).then(() => {
            this.$router.push({ path: this.redirect || '/' })
          }).catch(()=>{

          })
        }
      });
    },
  },
};
</script>
<style>
#login{

}
.login-img{

  word-break:keep-all;
  white-space:nowrap;
  margin: 20px auto;

  text-align: center;
  display: flex;
}
.login-form {
  max-width: 500px;
  position: fixed;

  height: 70%;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  margin: auto;
  text-align: center;

}
#components-form-demo-normal-login .login-form-forgot {
  float: right;
}
#components-form-demo-normal-login .login-form-button {
  width: 100%;
}
.bg{
  width: 100%;
  height: 100%;
  position: absolute;
  background-color: #42b983 !important;
}


.panel-login {
  border: none;
  -webkit-box-shadow: 0px 0px 49px 14px rgb(188 190 194 / 39%);
  -moz-box-shadow: 0px 0px 49px 14px rgba(188, 190, 194, 0.39);
  box-shadow: 0px 0px 49px 14px rgb(188 190 194 / 39%);
}

.panel {
  border-radius: 5px;
}
.panel {
  margin-bottom: 20px;
  background-color: #fff;
  border: 1px solid transparent;
  border-radius: 4px;
  -webkit-box-shadow: 0 1px 1px rgb(0 0 0 / 5%);
  box-shadow: 0 1px 1px rgb(0 0 0 / 5%);
}


.panel-body {
  padding: 15px;
}



div {
  display: block;
}
</style>