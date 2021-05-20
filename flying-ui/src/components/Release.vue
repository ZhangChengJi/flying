<template>

</template>

<script>
export default {
name: "Release",
  data(){
  return{

  }
  },
  methods:{
    showConfirm() {
      this.$confirm({
        title: '此spring 命名空间配置发生变化,是否进行发布?',
        content: h => <div style="color:red;">Some descriptions</div>,
        onOk() {
          console.log('OK');
        },
        onCancel() {
          console.log('Cancel');
        },
        class: 'test',
      });
    },
  }
}
</script>

<style scoped>

</style>