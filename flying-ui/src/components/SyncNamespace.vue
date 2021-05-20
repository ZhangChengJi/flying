<template>
  <a-modal
      :title="$t('Namespace.syncConfig')"
      :visible="visible"
      width="50%"
      :dialog-style="{ top: '20px' }"
      :maskClosable="false"
      :confirm-loading="confirmLoading"
      @ok="handleOk"
      :okText="this.$t('Button.sync')"
      @cancel="handleCancel"
  >


    <a-alert
        :message="this.$t('Release.releaseSyncDescription')"
        banner
        type="warning"
        closable
    />
    <br/>
    <a-alert
        :message="this.$t('Release.releaseSyncDescription2')"
        banner
        type="info"
        closable
    />
    <br/>
    <br/>
    <a-form-model
        ref="ruleForm"
        :model="form"
        :rules="rules"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
    >
      <a-form-model-item :label="this.$t('Release.selectSyncNode')" prop="nodes">
        <a-checkbox-group v-model="form.nodes" prop="nodes">
        <template v-for="item in nodeList" >
          <a-tooltip>
            <template slot="title" v-if="!item.status">
              {{ generateTitle('App.placeholder.selectNode') }}
            </template>
            <a-checkbox :value="item.id" :name="item.name"  v-bind:disabled="!item.status"  style="margin-right: 25px">
              {{ item.name }}
            </a-checkbox>
          </a-tooltip>
        </template>
        </a-checkbox-group>

      </a-form-model-item>
    </a-form-model>
  </a-modal>

</template>

<script>
import {getSyncNode} from "@/api/node";
import {createNamespace,syncNamespace,findRemoteNamespace} from "@/api/namespace";
import {mapGetters} from "vuex";
import {alertMessage} from "@/utils/alert";
import {generateTitle} from "@/utils/i18n";

export default {
  name: "SyncNamespace",
  data(){
    return {
      visible: false,
      loading: false,
      confirmLoading: false,
      nodeList: [],
      nodes: [],
      labelCol: { span: 8},
      wrapperCol: { span: 14 },
      form: {
        id: 0,
        nodes: [],
        nId: 0,
      }
    }
  },
  computed: {
    ...mapGetters(["appNodeId"]),
    rules() {
      return {
        nodes: [
          {
            type: 'array',
            required: true,
            message: this.$t('Namespace.Verification.selectNode'),
            trigger: 'change',
          },
        ],
      }
    }
  },
  methods:{
    generateTitle,
    showModal(id,appId,nodeId) {
      //this.form.nId=id

      this.visible=true;
      this.getAppNode(id,appId,nodeId);
    },
    getAppNode(id,appId,nodeId){
      this.form.id=id;
      this.form.nId=Number(nodeId);
      let syncId = {appId: appId, nodeId: nodeId};
      getSyncNode(syncId).then(response => {
        this.nodeList=response.data.data.node
      })
    },
    syncNamespace(){
      console.log(this.form)
      syncNamespace(this.form).then(rep =>{
        this.confirmLoading = false;
        if (alertMessage(this,rep.data)){
          this.visible = false;
          this.$parent.getAppNamespace(this.$route.query.appId,this.$route.query.nodeId,this.$route.query.status);
          this.resetForm();
        }
      })
    },
    handleOk(e) {
      const _this = this;
      this.$refs.ruleForm.validate(valid => {
        if (valid) {
          this.confirmLoading = true;
          findRemoteNamespace(this.form).then(resp=>{
            console.log(resp)
            if (resp.data.code===10042){
              var util = require('util');
              this.$confirm({
                title: util.format(this.$t('Release.releaseNamespaceNodeSyncHistory'),resp.data.msg),
                content: '',
                okText: this.$t('Button.determine'),
                okType: 'danger',
                cancelText: this.$t('Button.cancel'),
                onOk() {
                  _this.syncNamespace()
                },
                onCancel() {
                  _this.confirmLoading = false;
                  console.log('Cancel');
                },
              });
            }else {
              _this.syncNamespace()
            }

          })




        } else {
          console.log('error submit!!');
          return false;
        }
      });

    },
    handleCancel() {
      this.visible = false;
      this.confirmLoading = false;
      this.resetForm();
    },
    resetForm(){
      this.$refs.ruleForm.resetFields();
    }
  }
}
</script>

<style scoped>

</style>