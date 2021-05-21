<template>
  <a-drawer
      :title="generateTitle('Node.addNode')"
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
          <a-form-model-item  ref="url" :label="generateTitle('Node.List.nodeAddress')" prop="url">
            <a-input v-model="form.url"
                     :placeholder="generateTitle('Node.placeholder.nodeAddress')"
            />
          </a-form-model-item>
        </a-col>
      </a-row>
      <a-row :gutter="20">
        <a-col :span="20">
          <a-form-model-item ref="content" :label="generateTitle('Node.List.nodeDescription')" prop="content">
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
        {{ generateTitle('Button.create') }}
      </a-button>
    </div>
  </a-drawer>
</template>

<script>
import {createNode} from "@/api/node";
import {generateTitle} from "@/utils/i18n"
import {alertMessage} from "@/utils/alert";

export default {
name: "AddNode",
  data(){
  return{
    visible: false,
    form: {
      name: '',
      url: '',
      content: '',
    },
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
    showModal() {
      this.visible = true;
    },

    submit() {
      this.$refs.ruleForm.validate(valid => {
        if (valid) {
          this.iconLoading = true;
          createNode(this.form).then(rep =>{
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
      this.iconLoading = false;
      this.visible = false;
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