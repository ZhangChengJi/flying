<template>
  <a-drawer
      :title="$t('Namespace.addNamespace')"
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
      <a-form-model-item ref="name" :label="$t('Namespace.List.namespaceName')" prop="name">
        <a-input
            v-model="form.name"
            @blur="
          () => {
            $refs.name.onFieldBlur();
          }
        "
        />
      </a-form-model-item>
      <a-form-model-item :label="$t('Node.node')" prop="nodes">
        <a-checkbox-group v-model="form.nodes">
          <template v-for="item in nodeList" >

            <a-tooltip>
              <template slot="title" v-if="!item.status">
                {{ $t('App.placeholder.selectNode') }}
              </template>
          <a-checkbox :value="item.id" :name="item.name"  v-bind:disabled="!item.status"  style="margin-right: 25px">
            {{ item.name }}
          </a-checkbox>
            </a-tooltip>
          </template>

        </a-checkbox-group>
      </a-form-model-item>
      <a-form-model-item :label="$t('Namespace.List.format')" prop="format">
        <a-radio-group ref="" v-model="form.format" :options="options2" @change="onChange2" >
        </a-radio-group>

      </a-form-model-item>
      <a-form-model-item :label="$t('Namespace.List.configuration')" prop="value" >

      </a-form-model-item>
      <div :class="[isFull?'full':'editor-container']" >
        <a-button class="fullbutton" @click="showfull" icon="fullscreen"></a-button>
        <properties-editor  v-model="form.value"  v-if="!isYaml" :re="read" ref="editor" />
        <yaml-editor  v-model="form.value" v-if="isYaml" :re="read" ref="editor" />
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
        {{ $t('Button.create') }}
      </a-button>
    </div>
  </a-drawer>
</template>

<script>
import {getAppNode, getNodeAll} from "@/api/node";
import { mapGetters } from "vuex";
import VueEvent from "@/model/VueEvent";
const options2 = [
  { label: 'yaml', value: 'yaml' },
  { label: 'properties', value: 'properties' },
];
import YamlEditor from '@/components/YamlEditor';
import PropertiesEditor from "@/components/PropertiesEditor";
import {createNamespace} from "@/api/namespace";
import {alertMessage} from "@/utils/alert";

export default {
  name: "AddNamespace",
  components:{
    YamlEditor,
    PropertiesEditor
  },
  data(){
    return{
      isFull:false,
      value: '',
      format: 'yaml',
      isYaml: true,
      options2: options2,
      visible: false,
      loading: false,
      iconLoading: false,
      labelCol: { span: 5 },
      wrapperCol: { span: 14 },
      nodeList: [],
      appId: this.$route.query.asId,
      form: {
        name: '',
        nodes: [],
        format: 'yaml',
        value: '',
      },
    }
  },
  computed: {
    ...mapGetters(["asId"]),
    ...mapGetters(["appNodeId"]),
    rules() {
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
    getAppNode(){
      getAppNode(this.$route.query.appId).then(response => {
        this.nodeList=response.data.data.node

      })
    },
    showModal() {
      this.getAppNode();
      this.visible=true;
    },
    onChange2(e) {
      if(e.target.value==='yaml'){
        this.isYaml=true
      }else{
        this.isYaml=false
      }
    },
    submit() {
      this.$refs.ruleForm.validate(valid => {
        if (valid) {
          this.iconLoading = true;
          createNamespace(this.$route.query.appId,this.form.name,this.form.nodes,this.form.format,this.form.value).then(rep =>{
            console.log(rep.data)
            this.iconLoading = false;
            if (alertMessage(this,rep.data)){
              this.visible = false;
              VueEvent.$emit('refresh-namespace')
              this.resetForm();
            }
          })

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
      this.iconLoading = false;
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