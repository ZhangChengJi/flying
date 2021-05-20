<template>

  <a-layout-header style="background-color: #42b983">
    <a-row  type="flex" justify="space-around"   align="middle">
      <a-col :flex="15">

          <a href="/">   <img style="width:160px" src="../../image/flying3.png" />  </a>

      </a-col>
      <a-col :flex="1" style="display:flex;justify-content:flex-end">
        <a-dropdown placement="bottomLeft">
          <a style="color: #F7FDF7; font-size: 14px" class="handle ant-dropdown-link" @click="e => e.preventDefault()"><a-icon type="setting" theme="filled" />
            {{ generateTitle('Header.operation') }} <a-icon type="down" />
          </a>
          <a-menu slot="overlay" style="">
            <a-menu-item>

                <router-link to="/list/node"> <span style="font-size: 15px;">{{ generateTitle('Header.envManagement') }} </span></router-link>

            </a-menu-item>
            <a-menu-item>
              <router-link to="/list/app"> <span style="font-size: 15px;">{{ generateTitle('Header.appManagement') }}</span></router-link>

            </a-menu-item>

          </a-menu>
        </a-dropdown>
        <span class="language-switch language-switch-primary" @click="switchLang">
              {{ generateTitle('Header.languageSwitchButton') }}
            </span>
        <a-dropdown placement="bottomLeft">
          <a style="color: #F7FDF7; font-size: 14px" class="ant-dropdown-link" >
            <a-avatar style="backgroundColor:#87d068" icon="user"  :src="this.headerImg" size="large"  @click="e => e.preventDefault()"/>
           {{username}}<a-icon type="down" />
          </a>
          <br/>
        <a-menu slot="overlay" style="">
          <a-menu-item  @click="showPassword()">

            <span style="font-size: 15px;"> {{ generateTitle('Header.changePassword') }}</span>

          </a-menu-item>
          <a-menu-item @click="logout">
          <span style="font-size: 15px;"> {{ generateTitle('Header.logout') }}</span>
          </a-menu-item>
        </a-menu>
        </a-dropdown>
      </a-col>
    </a-row>
    <div>
    </div>
   <dev>

   </dev>
    <update-password ref="password"/>
  </a-layout-header>
</template>

<script>
import {mapGetters} from "vuex";
import store from "@/store";
import UpdatePassword from "@/components/UpdatePassword";
import {generateTitle} from "@/utils/i18n"

export default {
  name: "Navber",
  data() {
    return {

    }
  },
  components:{
    UpdatePassword,
  },
  computed: {
    ...mapGetters(["username"]),
    ...mapGetters(["headerImg"]),

  },
  methods:{
    generateTitle,
    showPassword(){
     this.$refs.password.showModal()
    },
    switchLang(){
      const currentLanguage = this.$store.getters.docsite_language === 'en' ? 'zh' : 'en';
      this.$i18n.locale = currentLanguage;
      this.$store.dispatch("SetLanguage", currentLanguage)

    },
    logout(){
      let _this=this;
      this.$confirm({
        title: this.generateTitle('Login.logoutTitle'),
        // content: 'Some descriptions',
        okText: this.$t('Button.determine'),
        okType: 'danger',
        cancelText: this.$t('Button.cancel'),
        onOk() {
          store.dispatch('FedLogOut').then(() => {
            _this.$notification['success']({
              message: _this.$t('Login.logoutSuccessfully')
            });
            _this.$router.push({path: '/login'})

          })
        },
        onCancel() {
          console.log('Cancel');
        },
      });
    }
  }
}
</script>

<style scoped>
.handle{
  margin-right: 20px;
}
.ant-layout-header{
  padding: 0 24px;
}
.language-switch {
  float: right;
  display: inline-block;
  box-sizing: border-box;
  width: 24px;
  height: 24px;
  line-height: 20px;
  margin-top: 21px;
  margin-right: 40px;
  text-align: center;
  border-radius: 2px;
  cursor: pointer;
  font-family: PingFangSC-Medium;
  font-size: 14px;
  opacity: 0.6;
}

.language-switch:hover {
  opacity: 1;
}
.language-switch-primary {
  border: 1px solid #fff;
  color: #fff;
}
.language-switch-normal {
  border: 1px solid #333;
  color: #333;
}


.language-switch {
    margin-right: 20px;
  }
</style>