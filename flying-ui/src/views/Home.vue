
<template>
  <a-layout id="c-layout"  >
    <node-status/>
    <app-list/>
  </a-layout>
</template>

<script>
// @ is an alias to /src
import NodeStatus from "@/components/NodeStatus";
import AppList from "@/components/AppList";
export default {
  name: 'Home',
  components: {
    NodeStatus,
    AppList,

  }
}
</script>
