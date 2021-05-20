<template>

  <dev id="appList" style="width: 100%;min-width: 900px">
    <add-app ref="add" :parent="1"/>
    <a-layout style="padding: 0 24px 24px">
      <a-layout-content width="100%"
                        :style="{ background: '#fff', margin: 0,height:'730px', minWidth: '384px',width: '100%' }"
      >

        <a-card :title="generateTitle('App.appProject')" :bordered="false" size="small" style="height: 100%;"
                :headStyle="{backgroundColor: '#42b983',color: '#fff'}"
                :bodyStyle="{padding: '24px',minWidth: '384px',width: '100%',height:'680px', overflow: 'auto' }">

          <a-icon title="添加应用" @click="addShow()" slot="extra" type="plus-circle"
                  style="color:#ffffff;font-size: 23px;cursor:pointer"/>

            <a-list :grid="{ gutter: 16,   column: 6 }" :data-source="appList" :loading="loading">

              <div
                  v-if="showLoadingMore && this.totalPage>this.query.page"
                  slot="loadMore"
                  :style="{ textAlign: 'center', marginTop: '12px', height: '32px', lineHeight: '32px' }"
              >
                <a-spin v-if="loadingMore"/>

                <div v-else style="margin: 0 auto;text-align: center">
                  <a-button @click="loadPage()" style="width: 100px;height: 10px;margin: 0 auto" type="link">
                    <svg t="1615960868472" class="icon" viewBox="0 0 1024 1024" version="1.1"
                         xmlns="http://www.w3.org/2000/svg" p-id="4722" width="100" height="100">
                      <path
                          d="M105.025641 0h236.307692c57.764103 0 105.025641 47.261538 105.025641 105.025641v236.307692c0 57.764103-47.261538 105.025641-105.025641 105.025641H105.025641c-57.764103 0-105.025641-47.261538-105.025641-105.025641V105.025641c0-57.764103 47.261538-105.025641 105.025641-105.025641z m0 577.641026h236.307692c57.764103 0 105.025641 47.261538 105.025641 105.025641v236.307692c0 57.764103-47.261538 105.025641-105.025641 105.025641H105.025641c-57.764103 0-105.025641-47.261538-105.025641-105.025641V682.666667c0-57.764103 47.261538-105.025641 105.025641-105.025641z m577.641026 0h236.307692c57.764103 0 105.025641 47.261538 105.025641 105.025641v236.307692c0 57.764103-47.261538 105.025641-105.025641 105.025641H682.666667c-57.764103 0-105.025641-47.261538-105.025641-105.025641V682.666667c0-57.764103 47.261538-105.025641 105.025641-105.025641z"
                          fill="#4AB984" p-id="4723"></path>
                      <path
                          d="M848.082051 24.94359L984.615385 161.476923c34.133333 34.133333 34.133333 87.958974 0 122.092308L848.082051 421.415385c-34.133333 34.133333-87.958974 34.133333-122.092307 0L589.45641 283.569231c-34.133333-34.133333-34.133333-87.958974 0-122.092308L727.302564 24.94359c32.820513-32.820513 87.958974-32.820513 120.779487 0z"
                          fill="#7ECDA7" p-id="4724" data-spm-anchor-id="a313x.7781069.0.i4" class="selected"></path>
                    </svg>
                    <span style="font-size: 10px;margin: 0 auto;color:#4AB984 ">加载更多</span>
                  </a-button>
                </div>
              </div>
              <a-list-item style="width: 142px" slot="renderItem" slot-scope="item, index">
                <a-card hoverable bordered size="small"
                        style="text-align: center;border-radius:3px; border:1px solid #DFF4E8"
                        @click="app(item.appId)">
                  <span style="font-size: 15px; color: #42b983  ">{{ item.name }}</span> <br/> <span
                    style="font-size: 10px;color: #7ECDA7 ">{{ item.appId }}</span>
                </a-card>

              </a-list-item>

            </a-list>

          <!--       v-if="this.query.page<this.totalPage -->



        </a-card>

      </a-layout-content>
    </a-layout>
  </dev>
</template>

<script>

import {getAppList} from "@/api/app";
import AddApp from "@/components/AddApp";
import {getAppNode} from "@/api/node";
import {generateTitle} from "@/utils/i18n";

export default {
  name: "AppList",
  components: {
    AddApp
  },
  data() {
    return {

      visible: false,
      appList: [],
      nodeList: [],
      total: 0,
      page: 1,
      totalPage: 0,
      query: {page: 1, pageSize: 10},
      isRouterAlive: true,
      isNotData: false,
      loading: true,
      loadingMore: false,
      showLoadingMore: true,
    }
  },
  created() {
    this.getAppList(this.query)

  },
  methods: {
    generateTitle,
    refreshApp(){
      this.query.page=1
      this.appList=[]
      this.getAppList(this.query)

    },
    getAppList(page) {
      this.loading = true;
      this.loadingMore = true;
      this.showLoadingMore = false;
      getAppList(page).then(response => {
        this.appList = this.appList.concat(response.data.data.list)
        this.total = response.data.data.total
        response.data.data.list.length <= 0 ? this.isNotData = true : this.isNotData = false
        this.totalPage = Math.ceil(this.total / 12) || 1;
        this.loading = false;
        this.loadingMore = false;
        this.showLoadingMore = true;
        console.log(this.totalPage)
        console.log(this.query.page)
      })
    },
    loadPage() {

      this.page++
      this.query.page++
      this.getAppList(this.query)
    },
    addShow() {
      this.$refs.add.showModal()
    },
    app(appId) {
      this.$store.commit('changeAsId', appId)
      this.$store.commit('clearAppNodeId')
      getAppNode(appId).then(response => {
        let data =response.data.data
        this.$router.push({path: '/app/namespace', query: {appId: appId,nodeId:data.node[0].id,status:data.node[0].status}});
      });

    },

  }

}
</script>

<style scoped>
.app {
  text-align: center;
  margin: 0 auto;

}

.gutter-example >>> .ant-row > div {
  background: transparent;
  border: 0;
}

.gutter-box {

  padding: 5px 0;
}

.icon {
  width: 20px;
  height: 15px;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
  color: rgb(138, 135, 135);

}

</style>
