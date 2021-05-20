<template>
<dev>
  {{this.$route.query.env }}
</dev>
</template>

<script>

export default {
  name: "Namespace",
  components:{

  },
  mounted() {
    if(!this.$route.query.env){
      alert(1)
    }
  }
}
</script>

<style scoped>

</style>