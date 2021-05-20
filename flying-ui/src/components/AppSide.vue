<template>

  <a-layout-sider  style="background: #f0f2f5;min-height: 500px; min-width: 100px;padding: 0 ">
    <a-card :title="$t('Node.nodeList')" :bordered="false" size="small" style="" :headStyle="{backgroundColor: '#42b983',color: '#fff'}" :bodyStyle="{padding: '1px',}">
      <dev>
        <a-menu
            style="width: 100%;"
            :default-selected-keys="current"
            mode="inline">
          <template v-for="item in nodeList">
            <a-menu-item :key="item.id" se  v-bind:data-item="item.status"  :style="!item.status ? 'color: red' :''">
              <a-icon :type="!item.status? 'frown' : 'smile'" />
              <span >{{ item.name }}</span>

                <router-link :to="{path: '/app/namespace',query:{appId: getApp,nodeId:item.id,status:item.status}}"/>
            </a-menu-item>
          </template>
        </a-menu>
      </dev>
    </a-card>
    <br/>
    <br/>

    <a-card :title="$t('App.appInfo')" :bordered="false" size="small" style="" :headStyle="{backgroundColor: '#42b983',color: '#fff'}" :bodyStyle="{padding: ' 15px'}">
     <dev >
       <span style="">

         <i  style="font-family: Monaco;  font-style: normal; font-size: 14px">{{$t('App.List.appId')}}: <i style="font-size: 12px; color: #42b983">{{appInfo.appId}}</i></i>
       </span>
         <br/>
       <span style="{text-align: center}">
         <i style="font-family: Monaco;  font-style: normal;font-size: 14px"> {{$t('App.List.appName')}}: </i><i style="font-size: 12px;color: #42b983">{{appInfo.name}}</i>
       </span>
     </dev>

    </a-card>
  </a-layout-sider>



</template>

<script>
import {getAppNode, getNodeAll} from "@/api/node";
import VueEvent from "@/model/VueEvent";
import {mapGetters} from "vuex";
export default {
name: "AppSide",
  data() {
    return {
     // current: ['mail'],
      openKeys: ['sub1'],
      defaultSelectKey: [],
      key: '',
      nodeList: [],
      appInfo: '',
      iconList: ['cloud-server','deployment-unit','coffee','meh','fire','alert','robot'],
      appId:'',
      nodeId: ['1']
    };
  },
  watch: {
    openKeys(val) {
      console.log('openKeys', val);
    },
  },

  computed: {
    ...mapGetters(["appNodeId"]),
    current: function () {
      return this.$route.query.nodeId
    },
    getApp: function (){
      return this.$route.query.appId
    }
  },
  mounted() {
  //this.appId=this.$route.query.appId;
    this.getAppNode()

  },
  methods: {
    randomIcon(){
      var n = Math.floor(Math.random() * this.iconList.length + 1)-1;
      return this.iconList[n];

    },
    getAppNode(){
      getAppNode(this.getApp).then(response => {
           console.log(response)
          if (response.data!=null){
             let data =response.data.data
            this.appInfo=data.app
            this.nodeList=data.node
             }
      })
    },


    // handleClick(e) {
    //      console.log( e.key)
    //   this.$store.commit('changeAppNodeId', e.key.split(':')[0])
    //   this.$store.commit('changeNodeStatus',e.key.split(':')[1])
    //   if (e.key.split(':')[1]){
    //     this.getAppNamespace(e.key.split(':')[0],e.key.split(':')[1])
    //   }
    //
    // },

    titleClick(e) {
      console.log('titleClick', e);
    },
  }
}
</script>

<style scoped>
.ant-card-small > .ant-card-body {
  padding: 1px;
}
</style>