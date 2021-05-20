<template>
  <a-layout style="padding: 0 24px 24px">

    <a-layout-content
        :style="{ background: '#fff', padding: '24px', margin: 0, minHeight: '280px',height:'700px' }"
    >
      <router-view/>
    </a-layout-content>
  </a-layout>
</template>

<script>
export default {
name: "AppMain"
}
</script>

<style scoped>

</style>