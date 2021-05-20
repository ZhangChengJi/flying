<template>
  <a-spin :spinning="this.$store.state.loading" tip="Loading..." delay="100">
  <a-layout >
    <navber/>
    <router-view/>
    <a-layout-footer style="text-align: center; padding: 0 ;">
      flying ©{{getDate}} Created by zcj <a href="https://github.com/ZhangChengJi/flying">github</a>
    </a-layout-footer>
  </a-layout>
  </a-spin>
</template>

<script>
import Navber from "./Navber"
export default {
  name: "Layout",
  data() {
    return {}
  },
  components: {
    Navber,
  },
  computed: {
    getDate: function () {
      var date1 = new Date();
      return date1.getFullYear();
    }
  }
}
</script>

<style>
#components-layout-demo-top-side .logo {
  width: 120px;
  height: 31px;
  background: rgba(255, 255, 255, 0.2);
  margin: 16px 28px 16px 0;
  float: left;
}
#c-layout{
 margin-left: 20px;
  margin-top: 20px;

}
ant-layout-footer {
  line-height: 1.5;
  background: #3ba0e9;
  color: #fff;
  line-height: 120px;
}
.spin-content {
  border: 1px solid #91d5ff;
  background-color: #e6f7ff;
  padding: 30px;
}
</style>