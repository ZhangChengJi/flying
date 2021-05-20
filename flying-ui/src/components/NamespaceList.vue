<template>

  <a-layout style="padding: 0 24px 24px 24px">

    <add-namespace ref="add"/>

    <a-layout-content
        :style="{ background: '#fff', padding: '0 ', margin: 0, minHeight: '700px',}">
      <a-card :title="$t('Namespace.namespace')" :bordered="false" size="small" style=""
              :headStyle="{backgroundColor: '#42b983',color: '#fff' }" :bodyStyle="{padding: '24px',}">
        <a-icon :title="$t('Button.create')" @click="addShow()" slot="extra" type="plus-circle"
                style="color:#ffffff;font-size: 23px;cursor:pointer"/>
          <router-view />

      </a-card>
    </a-layout-content>
  </a-layout>
</template>

<script>

import YamlEditor from "@/components/YamlEditor";
import PropertiesEditor from "@/components/PropertiesEditor";
import AddNamespace from "@/components/AddNamespace";
import {getAppNamespace, deleteNamespace} from "@/api/namespace";
import {createRelease} from "@/api/release";
import VueEvent from "@/model/VueEvent";
import {mapGetters} from "vuex";
import UpdateNamespace from "@/components/UpdateNamespace";
import SyncNamespace from "@/components/SyncNamespace";
import {formatTimeToStr} from "@/utils/date";

export default {
  name: "NamespaceList",
  components: {SyncNamespace, UpdateNamespace, YamlEditor, PropertiesEditor, AddNamespace},
  data() {
    return {

      ds: true,
      value2: 'yaml',
      mode: 'text/x-yaml',
      ModalText: 'Content of the modal',
      customStyle: '',
      other: '',
      namespaceList: [],
      isNotData: false,
      notTitle: '无法查看该namespace配置',
      status: this.nodeStatus,
      read: 'nocursor',
      activeKey: 0,
      nodeStatus: '',
      nodeList: [],
      appInfo: '',
      iconList: ['cloud-server', 'deployment-unit', 'coffee', 'meh', 'fire', 'alert', 'robot'],
      meunKey: '',


    };
  },
  watch: {

  },
  computed: {
    ...mapGetters(["asId"]),
    ...mapGetters(["appNodeId"]),
    ...mapGetters(["appNodeId"])
    //...mapGetters(["nodeStatus"])
  },
  mounted() {


  },
  filters: {
    formatDate: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return '最后发布时间：' + formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    }
  },
  methods: {
    addShow() {
      this.$refs.add.showModal()
    },
    update(id) {
      this.$refs.update.showModal(id)
    },
    sync(id) {
      this.$refs.sync.showModal(id)
    },
    change(data) {
      this.ds = data;
    },
    showDeleteConfirm(id, name) {
      const _this = this;
      let n = <span style="color: red">name</span>
      this.$confirm({
        title: h => <div style="">是否确定要删除namespace: <span style="color: red; font-size: 20px">{name}</span></div>,
        content: '删除后客户端会无法找到该namespace的配置',
        okText: '确定',
        okType: 'danger',
        cancelText: '取消',
        onOk() {
          deleteNamespace(id).then(resp => {
            let type = 'success';
            let msg = '删除成功'
            let de = '';
            if (resp.data.code !== 0) {
              type = 'error'
              msg = '删除失败'
              de = resp.data.msg
            } else {
              _this.getAppNamespace(_this.appNodeId[0]);
            }
            _this.$notification[type]({
              message: msg,
              description:
              de,
            });
          });
        },
        onCancel() {
          console.log('Cancel');
        },
      });
    },
    getAppNamespace(node) {

        this.namespaceList = []
        getAppNamespace({"aId": this.$route.query.asId, "nodeId": node}).then(response => {

          this.namespaceList = response.data.data.list
          this.status = response.data.data.node.status
          if (this.status) {
            response.data.data.list.length <= 0 ? this.isNotData = true : this.isNotData = false
          }
        });

    },
    getTime() {
      var date1 = new Date();
      var year = date1.getFullYear();
      var month = date1.getMonth() + 1;
      var day = date1.getDate();
      var hours = date1.getHours();
      var minutes = date1.getMinutes();
      var seconds = date1.getSeconds();
      return year + "年" + month + "月" + day + "日" + hours + ":" + minutes + ":" + seconds
    },
    showConfirm(id, name) {
      let reName = this.getTime() + '-RE';
      const _this = this;
      this.$confirm({
        okText: '发布',
        width: '25%',
        title: "此命名空间配置发生变化,是否进行发布?",
        content: h => <div style="">
          <span style="font-size: 16px">命名空间: &nbsp; &nbsp;   </span><span
            style="font-size: 16px;color:red">{name}</span>
          <br/>
          <span style="font-size: 16px">发布名称: &nbsp; &nbsp;   </span><span
            style="font-size: 16px;color:#F5B400">{reName}</span>
        </div>,
        onOk() {
          createRelease({nId: id, releaseName: reName}).then(resp => {
            let type = 'success';
            let msg = '发布成功'
            let de = '';
            if (resp.data.code !== 0) {
              type = 'error'
              msg = '发布失败'
              de = resp.data.msg
            } else {
              _this.getAppNamespace(_this.appNodeId[0]);
            }
            _this.$notification[type]({
              message: msg,
              description:
              de,
            });
          })
        },
        onCancel() {
          console.log('Cancel');
        },
        class: 'test',
      });
    },


    onSubmit() {
      this.$refs.ruleForm.validate(valid => {
        if (valid) {
          alert('submit!');
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    resetForm() {

    },
  }
}
</script>

<style scoped>
.editor-container {
  position: relative;
  height: 100%;
}


</style>
<!-- 标签样式
-->
<style>
.namespace-panel .namespace-attribute-public {
  margin-right: 5px;
}

.no-radius {
  border-radius: 0;
}

.label-info {
  background-color: #42b983;
}

.label {
  display: inline;
  padding: .2em .6em;
  font-size: 100%;
  font-weight: bold;
  line-height: 1;
  color: #fff;
  text-align: center;
  white-space: nowrap;
  vertical-align: baseline;

}

element.style {
  width: 30px;
}

p, td, span {
  word-wrap: break-word;
  word-break: break-all;
}

* {
  -webkit-box-sizing: border-box;
  -moz-box-sizing: border-box;
  box-sizing: border-box;
}

.test {

}
</style>
