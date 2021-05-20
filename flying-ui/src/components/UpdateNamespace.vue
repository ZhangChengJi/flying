<template>

  <a-drawer
      :title="$t('Namespace.updateNamespace')"
      :width="800"
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
      <a-form-model-item ref="name" :label="$t('Namespace.List.namespaceName')" >
        <a-input disabled
            v-model="form.name"

        />
      </a-form-model-item>
      <a-form-model-item :label="$t('Node.node')" >

          <a-checkbox  :value="node.name" default-checked disabled>
            {{ node.name }}
          </a-checkbox>


      </a-form-model-item>
      <a-form-model-item :label="$t('Namespace.List.format')" >
        <a-radio ref="" v-model="form.format"  disabled >
          {{form.format}}
        </a-radio>

      </a-form-model-item>
      <a-form-model-item :label="$t('Namespace.List.configuration')" prop="value">

      </a-form-model-item>
      <div :class="[isFull?'full':'editor-container']" >
        <a-button class="fullbutton" @click="showfull" icon="fullscreen"></a-button>
        <properties-editor v-model="form.value"  v-if="!isYaml" :re="read"   ref="editor" />
        <yaml-editor v-model="form.value" v-if="isYaml" :re="read" ref="editor"/>
      </div>
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
          zIndex: 999,
        }"
    >
      <a-button :style="{ marginRight: '8px' }" @click="onClose">
        {{ $t('Button.cancel') }}
      </a-button>
      <a-button type="primary"  :loading="iconLoading" @click="submit">
        {{ $t('Button.update') }}
      </a-button>
    </div>

  </a-drawer>

</template>

<script>
import {getAppNode, getNodeAll} from "@/api/node";
import { mapGetters } from "vuex";
const options2 = [
  { label: 'yaml', value: 'yaml' },
  { label: 'properties', value: 'properties' },
];
import YamlEditor from '@/components/YamlEditor';
import PropertiesEditor from "@/components/PropertiesEditor";
import {createNamespace, findNamespace,updateNamespace} from "@/api/namespace";
import {alertMessage} from "@/utils/alert";

export default {
  name: "UpdateNamespace",
  components:{
    YamlEditor,
    PropertiesEditor,

  },
  data(){
    return{
      isFull:false,
      fu: false,
      value: '',
      format: 'yaml',
      isYaml:  true ,
      options2: options2,
      visible: false,
      loading: false,
      iconLoading: false,
      labelCol: { span: 4},
      wrapperCol: { span: 14 },
      node: [],
      read: false,
      form: {
        value: '',
        nodeId: 0
      },
    }
  },
  computed: {
    current: function () {
      return this.$route.query.nodeId
    },
    rules(){
      return{
        name: [
          { required: true, message: this.$t('Namespace.Verification.pleaseInputNamespaceName'), trigger: 'blur' },
        ],
        nodes: [
          {
            type: 'array',
            required: true,
            message:  this.$t('Namespace.Verification.selectNode'),
            trigger: 'change',
          },
        ],
        format: [
          { required: true, message: this.$t('Namespace.Verification.selectFormat'), trigger: 'change' },
        ],
        value: [{ required: true, message: this.$t('Namespace.Verification.pleaseInputConfig'), trigger: 'blur' }]
      }
    }
  },
  methods:{
    showfull(){
      this.isFull = !this.isFull
      if(this.isFull){
        this.$refs.editor.refresh();
      }
    },
    findNamespace(id,appId,nodeId){
      findNamespace({"id":id, "nodeId": nodeId}).then(resp =>{
        if(resp.data.data.namespace.format==='yaml'){
          this.isYaml=true
        }else {
          this.isYaml=false
        }
        this.form=resp.data.data.namespace;
       this.node=resp.data.data.node;
      });
    },
    showModal(id,appId,nodeId) {
      this.visible=true;
     this.findNamespace(id,appId,nodeId)
    },
    onChange2(e) {
      if(e.target.value==='yaml'){
        this.isYaml=true
      }else{
        this.isYaml=false
      }
    },
    submit() {
      const _this = this;
      this.$refs.ruleForm.validate(valid => {
        if (valid) {
          this.iconLoading = true;
          this.form.nodeId=Number(this.$route.query.nodeId)
          updateNamespace(this.form).then(rep =>{
            this.iconLoading = false;
            debugger
            if (alertMessage(this,rep.data)){
              this.visible = false;
              this.$parent.getAppNamespace(this.$route.query.appId,this.$route.query.nodeId,this.$route.query.status);
              this.resetForm();
            }
          });

        } else {
          console.log('error submit!!');
          return false;
        }
      });

    },
    handleCancel() {
      console.log('Clicked cancel button');
      this.visible = false;
      this.iconLoading = false;
      this.resetForm();
    },
    resetForm(){
      this.value=''
      this.format='yaml'
      this.$refs.ruleForm.resetFields();
    },
    onClose() {
      this.visible = false;
      this.resetForm();
    },

  }
}
</script>

<style scoped>
.full{
  position: fixed !important;
  top: 20px;
  left: 20px;
  z-index: 99999;
  width: calc(100vw - 40px);
  height: calc(100vh - 40px);
}

.fullbutton{
  position: absolute;
  right: 20px;
  top: 5px;
  z-index: 9;
}
.full >>> .properties-editor,.full >>> .CodeMirror {
  height: calc(100vh - 40px) !important;
}
.full >>> .yaml-editor,.full >>> .CodeMirror {
  height: calc(100vh - 40px) !important;
}
.editor-container{
  position: relative;

}

.ant-form-item-control {
  line-height: 20;
}
.properties-editor{
  width: 100%; height: 450px; overflow:auto;
}
.yaml-editor{
  width: 100%; height: 450px; overflow:auto;
}
.properties-editor >>> .CodeMirror {
  height: 450px;
  min-height: 273px;
}
.yaml-editor >>> .CodeMirror {
  height: 450px;
  min-height: 273px;

}
.yaml-editor >>> .CodeMirror-scroll{
  min-height: 273px;
}
.yaml-editor >>> .cm-s-rubyblue span.cm-string {
  color: #F08047;
}

</style>