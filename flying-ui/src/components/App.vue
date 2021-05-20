<template>
  <a-layout style="padding: 24px ">
    <add-app ref="add" :parent="3"/>
    <update-app ref="update"/>
    <a-layout-content
        :style="{ background: '#fff', padding: '0 ', margin: 0, minHeight: '700px',}">
      <a-card :title="generateTitle('App.appManagement')" :bordered="false" size="small" :headStyle="{backgroundColor: '#42b983',color: '#fff' }" :bodyStyle="{padding: '24px',}">
        <a-button  class="editable-add-btn" type="primary" style="margin-bottom: 8px" @click="addShow()">
          {{ generateTitle('App.addApp') }}
        </a-button>
      <a-table :columns="columns" :data-source="appList"  :row-key="appList => appList.id"  :pagination="pagination"
               @change="handleTableChange">
    <a slot="name" slot-scope="text">{{ text }}</a>
    <span slot="customTitle"><a-icon type="smile-o" /> Name</span>
    <span slot="node" slot-scope="node">
      <a-tag
          v-for="tag in node"
          :key="tag"
          :color="tag === 'loser' ? 'volcano' : tag.length > 1 ? '#87d068' : '#108ee9'"
      >
        {{ tag.toUpperCase() }}
      </a-tag>
    </span>
        <span slot="createTime" slot-scope="text">{{  text |formatDate}}</span>
        <span slot="action" slot-scope="text, record,index" >
  <a-space>
      <a-config-provider :auto-insert-space-in-button="false">
      <a-button type="primary"  @click="updateShow(text.id)">
        {{ generateTitle('App.update')}}
      </a-button>
    </a-config-provider>
      <a-button type="danger" @click="showDeleteConfirm(text)">
      {{ generateTitle('App.delete')}}
    </a-button>
  </a-space>
    </span>
  </a-table>
      </a-card>
    </a-layout-content>
  </a-layout>
</template>
<script>
import {deleteNode} from "@/api/node";
import UpdateApp from "@/components/UpdateApp";
import {generateTitle} from "@/utils/i18n";
import  {getAppList,deleteApp} from "@/api/app";
import AddApp from "@/components/AddApp";
import {formatTimeToStr} from "@/utils/date";
import {alertMessage} from "@/utils/alert";

export default {
  data() {
    return {
      appList: [],
      pagination: {
        defaultPageSize:5,
        onShowSizeChange:(current, pageSize)=>this.pageSize = pageSize,
        total:0 //总条数
      },
    };
  },
  computed:{
    columns() {
      return [{
        title: this.generateTitle('App.List.appName'),
        dataIndex: 'name',
        key: 'name',
        slots: { title: 'customTitle' },
        scopedSlots: { customRender: 'name' },
      },
        {
          title: this.generateTitle('App.List.appId'),
          dataIndex: 'appId',
          key: 'appId',
        },
        {
          title: this.generateTitle('App.List.useEnvironment'),
          key: 'node',
          dataIndex: 'node',
          scopedSlots: { customRender: 'node' },
        },
        {
          title: this.generateTitle('App.List.createTime'),
          dataIndex: 'createTime',
          key: 'createTime',
          scopedSlots: { customRender: 'createTime' },
        },
        {
          title: this.generateTitle('App.List.operation'),
          key: 'action',
          scopedSlots: { customRender: 'action' },
        }]
    }
},
  components:{
    UpdateApp,
    AddApp
  },
  filters: {
    formatDate: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    }
  },
  mounted() {
    this.getAppList()
  },
  methods: {
    generateTitle,
    addShow(){
      this.$refs.add.showModal()

    },
    updateShow(id){
      this.$refs.update.showModal(id)
    },
    handleTableChange(pagination) {
      console.log(pagination);
      const pager = {...this.pagination};
      pager.current = pagination.current;
      this.pagination = pager;
      this.getAppList({
        results: pagination.pageSize,
        page: pagination.current,
      });
    },
    getAppList(params = {}){
      getAppList(this.params).then(resp=> {
        const pagination = {...this.pagination};
        console.log(resp)
        pagination.total = resp.data.data.total
        this.pagination = pagination;
        this.appList = resp.data.data.list
      })
    },
    showDeleteConfirm(text) {
      console.log(text.name)
      const _this = this;
      this.$confirm({
        title: h => <div style="">{this.$t('App.areDelete')}<span style="color: red; font-size: 20px">{text.name}</span></div>,
        content: this.$t('App.confirmDelete'),
        okText: this.$t('Button.determine'),
        okType: 'danger',
        cancelText: this.$t('Button.cancel'),
        onOk() {
          deleteApp(text).then(resp=> {
            if (alertMessage(_this,resp.data)){
              _this.getAppList()
            }

          });
        },
        onCancel() {
          console.log('Cancel');
        },
      });
    },
  }
};
</script>
