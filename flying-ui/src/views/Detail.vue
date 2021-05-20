<template>
  <a-layout id="c-layout"   >
   <app-side/>
    <namespace-list />
  </a-layout>
</template>

<script>
import AppSide from "@/components/AppSide";
import Namespace from "@/components/Namespace";
import NamespaceList from "@/components/NamespaceList";
export default {
name: "Detail",
  data(){
  return{
    nodeId: '',
  }
  },
  components:{
    NamespaceList,
    AppSide,


  },
  methods:{

  }

}
</script>

<style scoped>

</style>