<template>
  <dev>
    <a-drawer
        :title="generateTitle('App.updateApp')"
        :width="500"
        :visible="visible"
        :body-style="{ paddingBottom: '80px' }"
        @close="onClose"
        :closable="false"
        :destroyOnClose="true"
        VNode="slot"
    >
      <a-form-model
          ref="ruleForm"
          :model="form"
          :rules="rules"
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
      >
        <a-row :gutter="20">
          <a-col :span="20">
            <a-form-model-item ref="appId" :label="generateTitle('App.List.appId')" prop="appId" >
              <a-input v-model="form.appId"
                       :placeholder="generateTitle('App.placeholder.appId')"
                       disabled="true"
              />
            </a-form-model-item>
          </a-col>

        </a-row>
        <a-row :gutter="20">
          <a-col :span="20">
            <a-form-model-item :label="generateTitle('App.List.appName')" ref="name" prop="name">
              <a-input v-model="form.name"
                       :placeholder="generateTitle('App.placeholder.appName')"
              />
            </a-form-model-item>
          </a-col>


        </a-row>
        <a-row :gutter="20">
        <a-alert
            :message="generateTitle('App.placeholder.updateCheckNode')"
            banner
            type="warning"
            closable
        />
        </a-row>
        <a-row :gutter="20">
          <a-col :span="20">
            <a-form-model-item :label="generateTitle('Node.node')"  prop="nodes">
              <a-checkbox-group v-model="form.nodes">

                <template v-for="item in nodeList" >

                  <a-tooltip>
                    <template slot="title" v-if="!item.status">
                      {{ generateTitle('App.placeholder.selectNode') }}
                    </template>

                    <a-checkbox :value="item.id" :name="item.name"  v-bind:disabled="!item.status" style="margin-right: 25px">
                      {{ item.name }}
                    </a-checkbox>

                  </a-tooltip>

                </template>



              </a-checkbox-group>
            </a-form-model-item>
          </a-col>

        </a-row>

      </a-form-model>
      <div
          :style="{
          position: 'absolute',
          right: 0,
          bottom: 0,
          width: '100%',
          borderTop: '1px solid #e9e9e9',
          padding: '10px 16px',
          background: '#fff',
          textAlign: 'right',
          zIndex: 1,
        }"
      >
        <a-button :style="{ marginRight: '8px' }" @click="onClose">
          {{ generateTitle('Button.cancel') }}
        </a-button>
        <a-button type="primary"  :loading="iconLoading" @click="submit">
          {{ generateTitle('Button.update') }}
        </a-button>
      </div>
    </a-drawer>
  </dev>
</template>

<script>

import {getNodeAll, updateNode} from "@/api/node";
import {findApp,updateApp} from "@/api/app";
import {generateTitle} from "@/utils/i18n";
import {alertMessage} from "@/utils/alert"

export default {
  name: "AddApp",
  data() {
    return {
      form: {
        appId: '',
        name: '',
        nodes: []
      },
      visible: false,
      nodeList: [],
      iconLoading: false,
    }
  },
  computed: {
    rules() {
      return {
        appId: [
          {required: true, message: this.generateTitle('App.Verification.pleaseInputAppId'), trigger: 'blur'}
        ],
        name: [
          {required: true, message: this.generateTitle('App.Verification.pleaseInputAppName'), trigger: 'blur'}
        ],
        nodes: [{
          type: 'array',
          required: true,
          message: this.generateTitle('App.Verification.selectNode'),
          trigger: 'change',
        }]
      }
    }
  },
  created() {
    this.getNodeAll()
  },
  methods: {
    generateTitle,
    showModal(id) {
      this.visible = true;
      this.findApp(id);
    },
    findApp(id){
      findApp(id).then(resp =>{
        console.log(resp)
        this.form=resp.data.data.app

      })
    },
    getNodeAll() {
      getNodeAll().then(response => {
        this.nodeList = response.data.data.list

      })
    },
    submit() {
      this.$refs.ruleForm.validate(valid => {
        if (valid) {
          this.iconLoading = true;
          updateApp(this.form).then(rep => {
            this.iconLoading = false;
            if (alertMessage(this, rep.data)) {
              this.visible = false;
              this.resetForm();
              this.$parent.$parent.$parent.getAppList()
            }

          })

        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },

    onClose() {
      this.visible = false;
      this.iconLoading = false;
      this.resetForm();
    },
    resetForm() {
      this.$refs.ruleForm.resetFields();
    }
  }

}
</script>

<style scoped>
.ant-drawer-header {
  background-color: #42b983;
}
</style>