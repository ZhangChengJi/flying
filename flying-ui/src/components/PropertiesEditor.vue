<template>
  <div   class="properties-editor" >
    <textarea ref="textarea" />

  </div>
</template>

<script>
import CodeMirror from 'codemirror'
import 'codemirror/addon/lint/lint.css'
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/lucario.css'
import 'codemirror/mode/properties/properties'
import 'codemirror/addon/lint/lint'


export default {
  name: 'PropertiesEditor',
  // eslint-disable-next-line vue/require-prop-types
  props: ['value','re'],
  data() {
    return {
      propertiesEditor: false,
      isValid: true,

    }
  },
  watch: {
    value(value) {
      const editorValue = this.propertiesEditor.getValue()
      if (value !== editorValue) {
        this.propertiesEditor.setValue(this.value)
      }
    }
  },

  mounted()  {
    this.relt2();

  },
  methods: {
    relt2(){
      this.propertiesEditor = CodeMirror.fromTextArea(this.$refs.textarea, {
        lineNumbers: true, // 显示行号
        mode: 'properties', // 语法model
        gutters: ['CodeMirror-linenumbers', 'CodeMirror-foldgutter', 'CodeMirror-lint-markers'],  // 语法检查器
        theme: 'lucario', // 编辑器主题
        readOnly: this.re,
        lint: true, // 开启语法检查
        //fixedGutter:true, //设置gutter跟随编辑器内容水平滚动（false）还是固定在左侧（true或默认）
        lineWrapping:true, //代码折叠
        foldGutter: true,
        fullScreen:true,//全屏
        matchBrackets:true,//括号匹配
        //scrollbarStyle:null //隐藏滚动条

      })

      this.propertiesEditor.setValue(this.value)
      this.propertiesEditor.on('change', (cm) => {
        this.$emit('changed', cm.getValue())
        this.$emit('input', cm.getValue())
      })

    },
    getValue() {
      return this.propertiesEditor.getValue()
    },
    refresh(){
      setTimeout(() => {
        this.propertiesEditor.refresh();
      },1);
    }
  }
}
</script>

<style scoped>
.properties-editor{
  height: 100%;
  font-size: inherit;

}
.properties-editor >>> .CodeMirror {
  height: auto;
  min-height: 273px;
  font-family: Monaco;

}
.properties-editor >>> .CodeMirror-scroll{
  min-height: 273px;

}
.properties-editor >>> .cm-s-rubyblue span.cm-string {
  color: #F08047;
}
</style>
