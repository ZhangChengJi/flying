<template>
  <div class="yaml-editor" >
    <textarea ref="textarea"  />

  </div>
</template>

<script>
import CodeMirror from 'codemirror'
import 'codemirror/addon/lint/lint.css'
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/lucario.css'
import 'codemirror/mode/yaml/yaml'
import 'codemirror/mode/properties/properties'
import 'codemirror/addon/lint/lint'
import 'codemirror/addon/lint/yaml-lint'
window.jsyaml = require('js-yaml') // 引入js-yaml为codemirror提高语法检查核心支持

export default {
  name: 'YamlEditor',
  // eslint-disable-next-line vue/require-prop-types
  props: ['value','re'],
  data() {
    return {
      yamlEditor: false,
      isValid: true,
      jsonData: '123232'

    }
  },
  watch: {
    value(value) {
      const editorValue = this.yamlEditor.getValue()
      if (value !== editorValue) {
        this.yamlEditor.setValue(this.value)
      }
    }
  },

  mounted()  {
    this.relt2();

  },

  methods: {
    relt2(){
      this.yamlEditor = CodeMirror.fromTextArea(this.$refs.textarea, {
        lineNumbers: true, // 显示行号
        mode: 'yaml', // 语法model
        gutters: ['CodeMirror-linenumbers', 'CodeMirror-foldgutter', 'CodeMirror-lint-markers'],  // 语法检查器
        theme: 'lucario', // 编辑器主题
        readOnly: this.re,
        lint: true, // 开启语法检查
        //fixedGutter:true, //设置gutter跟随编辑器内容水平滚动（false）还是固定在左侧（true或默认）
        lineWrapping:true, //代码折叠
        foldGutter: true,
        fullScreen:true,//全屏
        matchBrackets:true,//括号匹配
      })

      this.yamlEditor.setValue(this.value)
      this.yamlEditor.on('change', (cm) => {
        this.$emit('changed', cm.getValue())
        this.$emit('input', cm.getValue())
      })


    },
    getValue() {
      return this.yamlEditor.getValue()
    },
    refresh(){
      setTimeout(() => {
        this.yamlEditor.refresh();
      },1);
    }
  }
}
</script>

<style scoped>
.yaml-editor{
  height: 100%;
  font-size: inherit;

}
.yaml-editor >>> .CodeMirror {
  height: auto;
  min-height: 273px;
  font-family: Monaco;


}
.yaml-editor >>> .CodeMirror-scroll{
  min-height: 273px;
}
.yaml-editor >>> .cm-s-rubyblue span.cm-string {
  color: #F08047;
}

</style>
