<template>

  <a-layout style="padding: 24px ">
    <add-node ref="add"/>
    <update-node ref="update"/>

    <a-layout-content
        :style="{ background: '#fff', padding: '0 ', margin: 0, minHeight: '700px',}">
      <a-card :title="generateTitle('Node.nodeManagement')" :bordered="false" size="small" :headStyle="{backgroundColor: '#42b983',color: '#fff' }" :bodyStyle="{padding: '24px',}">

        <a-alert
            message=""
            :description="errMessage+generateTitle('Node.Error.conError')"
            type="error"
            show-icon
            closable
            banner
            style="margin-bottom: 10px"
            v-if="errMessage!==''"
        />
        <a-button  class="editable-add-btn" type="primary" style="margin-bottom: 8px" @click="addShow()">
          {{generateTitle('Node.addNode') }}
        </a-button>
          <!--  ##############环境列表###################      -->
    <a-table :columns="columns" :data-source="nodeList" :row-key="nodeList => nodeList.id"  :pagination="pagination"
             @change="handleTableChange">

    <a slot="name" slot-scope="text">{{ text }}</a>

    <span slot="status"  slot-scope="text" >
      <svg class="icon" aria-hidden="true">
          <use :xlink:href="text ? '#icon-chenggong' : '#icon-bupipei' "></use>
        </svg>
    </span>


    <span slot="action" slot-scope="text, record,index" >
  <a-space>
    <a-tooltip :title="generateTitle('Node.testCon')">
   <a-button shape="circle"  icon="reload" :loading="text.loading"  @click="enterIconLoading(text,index)">
    </a-button>
    </a-tooltip>
      <a-config-provider :auto-insert-space-in-button="false">
      <a-button type="primary" @click="updateShow(text.id)">
        {{ generateTitle('Node.update')}}
      </a-button>
    </a-config-provider>
      <a-button type="danger"  @click="showDeleteConfirm(text)">
      {{ generateTitle('Node.delete')}}
    </a-button>
  </a-space>
    </span>
  </a-table>
    </a-card>
  </a-layout-content>
  </a-layout>
</template>

<script>
import AddNode from "@/components/AddNode";
import UpdateNode from "@/components/UpdateNode";
import {generateTitle} from "@/utils/i18n"
import {alertMessage} from "@/utils/alert";

import {getNodeList,deleteNode,ping} from "@/api/node";

export default {
  name: "Node",
  data() {
    return {


      query: {page: 1, pageSize: 3},
      nodeList: [],
      pagination: {
        defaultPageSize:5,
        onShowSizeChange:(current, pageSize)=>this.pageSize = pageSize,
        total:0 //总条数
      },
      errMessage: '',
      iconLoading: false,
    };
  },
  components:{
    AddNode,
    UpdateNode
  },
  computed:{
    columns(){
      return [
        {
          title: this.generateTitle('Node.List.conStatus'),
          dataIndex: 'status',
          key: 'status',
          slots: { title: 'status' },
          scopedSlots: { customRender: 'status'},
        },
        {
          title: this.generateTitle('Node.List.nodeName'),
          dataIndex: 'name',
          key: 'name',
        },
        {
          title: this.generateTitle('Node.List.nodeAddress'),
          dataIndex: 'url',
          key: 'url',
        },
        {
          title: this.generateTitle('Node.List.nodeDescription'),
          key: 'content',
          dataIndex: 'content',
          scopedSlots: { customRender: 'content' },
        },
        {
          title: this.generateTitle('Node.List.operation'),
          key: 'action',
          scopedSlots: { customRender: 'action' },
        },
      ]
    }
  },
  mounted() {
    this.getNodeList()
  },
  methods:{
    generateTitle,
    handleTableChange(pagination) {
      console.log(pagination);
      const pager = { ...this.pagination };
      pager.current = pagination.current;
      this.pagination = pager;
      this.getNodeList({
        results: pagination.pageSize,
        page: pagination.current,
      });
    },
    getNodeList(params = {}){
      const _this = this;
      getNodeList(this.params).then(resp=>{
        const pagination = { ...this.pagination };
        pagination.total=resp.data.data.total
        this.pagination = pagination;
        this.nodeList=resp.data.data.list
        this.errMessage=''
        this.nodeList.forEach(function (item,index){
          item.loading=false
          if (!item.status){
            _this.errMessage+= item.name +" "
          }
        })
      })
    },
    addShow(){
      this.$refs.add.showModal()
    },
    updateShow(id){
      this.$refs.update.showModal(id)
    },
    showDeleteConfirm(text) {
      const _this = this;
      this.$confirm({
        title: h => <div style="">{this.$t('Node.areDelete')}<span style="color: red; font-size: 20px">{text.name}</span></div>,
        content: this.$t('Node.confirmDelete'),
        okText: this.$t('Button.determine'),
        okType: 'danger',
        cancelText:  this.$t('Button.cancel'),
        onOk() {
          deleteNode(text).then(resp=> {


            if (alertMessage(_this,resp.data)){
              _this.getNodeList()
            }
          });
        },
        onCancel() {
          console.log('Cancel');
        },
      });
    },
    deleteNode(obj){

    },
    enterIconLoading(text,index) {
      text.loading = true
      this.nodeList.splice(index,1,text);

      ping(text.id).then(resp=> {
          text.loading = false
          this.nodeList.splice(index,1,text);
        if (alertMessage(this,resp.data)){
          this.getNodeList()
        }
      });


    },
  }
}
</script>

<style scoped>
.shuanxin{
  background-color: #6a3fcc;
}
</style>