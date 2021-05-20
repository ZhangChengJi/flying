<template>

  <a-layout-sider width="400px" style="background: #fff;height: 730px">

    <a-card :title="generateTitle('Node.nodeStatus')" :bordered="false" size="small"  :headStyle="{backgroundColor: '#42b983',color: '#fff'}" :bodyStyle="{}">
      <a-alert type="error" v-if="errMessage!==''" show-icon :message="errMessage+generateTitle('Node.Error.conError') " banner closable />
    <a-list style="width: 350px;margin: 0 auto" item-layout="horizontal" size="small" :data-source="nodeList" >

      <a-list-item  slot="renderItem" slot-scope="item">
        <a-space :size="30">

        <svg class="icon" aria-hidden="true">
          <use :xlink:href="item.status ?'#icon-chenggong' : '#icon-bupipei' "></use>
        </svg>

        <a-list-item-meta>

          <dev slot="title" :href="item.name">{{ item.name }}</dev>

        </a-list-item-meta>

        <a-list-item-meta>
          <a slot="description">{{item.url}}</a>
        </a-list-item-meta>
          <a-list-item-meta>
            <span slot="description">{{item.content}}</span>
          </a-list-item-meta>
        </a-space>
      </a-list-item>

    </a-list>

    </a-card>
  </a-layout-sider>

</template>

<script>
import {getNodeAll} from "@/api/node";
import {generateTitle} from "@/utils/i18n"
export default {
  name:"NodeList",
  data() {
    return {

      nodeList: [],
      ifAlert: false,
      errMessage: '',
    };
  },
  created() {
    this.getNodeAll()
  },
  methods: {
    generateTitle,
    getNodeAll() {
      const _this = this;
      getNodeAll().then(response => {
        console.log(response)
        let list = response.data.data.list;
        _this.errMessage=''
        list.forEach(function (item) {
          if (!item.status){
            _this.errMessage+= item.name +" "
          }
        })

        this.nodeList = response.data.data.list

      })
    },
  }
};
</script>

<style scoped>
.ant-space-item{
  margin-right: 10px !important;
}
</style>