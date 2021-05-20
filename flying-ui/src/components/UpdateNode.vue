<template>
  <a-drawer
      :title="generateTitle('Node.updateNode')"
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
          <a-form-model-item ref="name" :label="generateTitle('Node.List.nodeName')" prop="name">
            <a-input v-model="form.name"
                     :placeholder="generateTitle('Node.placeholder.nodeName')"
            />
          </a-form-model-item>
        </a-col>

      </a-row>
      <a-row :gutter="20">
        <a-col :span="20">
          <a-form-model-item :label="generateTitle('Node.List.nodeAddressNotUpdate')" ref="url" prop="url">
            <a-input v-model="form.url" disabled="true"
                     :placeholder="generateTitle('Node.placeholder.nodeAddress')"
            />
          </a-form-model-item>
        </a-col>
      </a-row>
      <a-row :gutter="20">
        <a-col :span="20">
          <a-form-model-item :label="generateTitle('Node.List.nodeDescription')" ref="content" prop="content">
            <a-input v-model="form.content" type="textarea" />
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
</template>

<script>
import {findNode} from "@/api/node";
import {updateNode} from "@/api/node";
import {generateTitle} from "@/utils/i18n";
import {alertMessage} from "@/utils/alert";
export default {
name: "AddNode",
  data(){
  return{
    visible: false,
    form: [],
    iconLoading: false,
  }
  },
  computed:{
    rules() {
      return{
        name: [
          {required: true, message: this.generateTitle('Node.Verification.pleaseInputNodeName'), trigger: 'blur'}
        ],
        url: [
          {required: true, message: this.generateTitle('Node.Verification.pleaseInputNodeAddress'), trigger: 'blur'}
        ],
      }

    }
  },
  methods: {
    generateTitle,
    findNode(id){
      findNode(id).then(resp =>{
            console.log(resp)
        this.form=resp.data.data.node

      })
    },
    showModal(id) {
      this.visible = true;
      this.findNode(id);
    },

    submit() {
      this.$refs.ruleForm.validate(valid => {
        if (valid) {
          this.iconLoading = true;
          updateNode(this.form).then(rep =>{
            this.iconLoading = false;
             if(alertMessage(this,rep.data)){
               this.visible = false;
               this.resetForm();
               this.$parent.$parent.$parent.getNodeList()
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

</style>