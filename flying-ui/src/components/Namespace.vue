<template>

<dev>
  <update-namespace ref="update"/>
  <sync-namespace ref="sync"/>
          <!--   环境无法连接此错误展示     -->
          <a-result v-show="!nodeStatus"
                    status="error"
                    :title="generateTitle('Namespace.namespaceNodeError')"
                    :sub-title="generateTitle('Namespace.namespaceNodeDescription')"
          >


            <div class="desc">
              <p style="font-size: 16px;">
                <strong>{{generateTitle('Namespace.namespaceTestCaption')}}</strong>
              </p>
              <p>
                <a-icon :style="{ color: 'red' }" type="close-circle"/>
                {{$t('Namespace.namespaceTestConn')}}
                <a>tcp://{{$t('Namespace.nodeAddress')}} </a>{{ $t('Namespace.isWrong') }}
              </p>
              <p>
                <a-icon :style="{ color: 'red' }" type="close-circle"/>
                {{$t('Namespace.refreshDescription')}}
                <a href="/list/node">{{$t('Namespace.refresh')}}</a>{{$t('Namespace.testConn')}}
              </p>
              <p>
                <a-icon :style="{ color: 'red' }" type="close-circle"/>
                {{$t('Namespace.checkRunning')}}

              </p>
            </div>
          </a-result>

          <a-empty v-if="isNotData && nodeStatus"/>

          <!--  ############### namespace.js list  ######################   -->

  <div
      v-if="showLoadingMore"
      slot="loadMore"
      :style="{ textAlign: 'center', marginTop: '12px', height: '32px', lineHeight: '32px' }">

    <a-spin v-if="loadingMore"/>
  </div>
          <template v-for="item in namespaceList" v-else>
        <span class="text-center namespace-attribute-public label label-info no-radius">
        <span ng-bind="namespace.format" style="width:30px;" class="ng-binding">{{ item.format }}</span>
            </span>
            <a-page-header
                style="border: 1px solid #42b983;margin-bottom: 20px; "
                :title="item.name"
                :sub-title="item.releaseTime| formatDate">
              <template slot="tags">
                <a-tag color="cyan" :visible="item.status">
                  {{$t('Namespace.releaseTag') }}
                </a-tag>
                <a-tag color="orange" :visible="!item.status">
                  {{$t('Namespace.noReleaseTag') }}
                </a-tag>
              </template>
              <template slot="extra">
                <a-button key="4" @click="sync(item.id)">
                  {{$t('Button.sync')}}
                </a-button>
                <a-config-provider :auto-insert-space-in-button="false">
                  <a-button key="3" type="primary" v-bind:disabled="item.status"
                            @click="showConfirm(item.id,item.name)">
                    {{$t('Button.release')}}
                  </a-button>
                </a-config-provider>
                <a-button key="2" @click="update(item.id)" type="">
                  {{$t('Button.update')}}
                </a-button>
                <a-button key="1" type="danger" icon="delete" @click="showDeleteConfirm(item.id,item.name)">
                  {{$t('Button.delete')}}
                </a-button>
              </template>

              <a-collapse :bordered="false">
                <template #expandIcon="props">
                  <a-icon type="right" :rotate="props.isActive ? 90 : 0"/>
                </template>
                <a-collapse-panel :key="item.id" :header="$t('Namespace.foldConfig')"
                                  style="background: #fff;border-radius: 4px;margin-bottom: -15px;border: 0;overflow: hidden;margin-left: -16px;margin-top: -15px;">
                  <div class="editor-container">
                    <yaml-editor v-model="item.value" v-if="item.format=='yaml'" :re="read"/>
                    <properties-editor v-model="item.value" v-if="item.format=='properties'" :re="read"/>
                  </div>
                </a-collapse-panel>
              </a-collapse>
            </a-page-header>
          </template>
</dev>

</dev>
</template>

<script>

import YamlEditor from "@/components/YamlEditor";
import PropertiesEditor from "@/components/PropertiesEditor";
import AddNamespace from "@/components/AddNamespace";
import {getAppNode} from "@/api/node";
import {getAppNamespace, deleteNamespace} from "@/api/namespace";
import {createRelease} from "@/api/release";
import VueEvent from "@/model/VueEvent";
import {mapGetters} from "vuex";
import UpdateNamespace from "@/components/UpdateNamespace";
import SyncNamespace from "@/components/SyncNamespace";
import {formatTimeToStr} from "@/utils/date";
import {generateTitle} from "@/utils/i18n";
import {alertMessage} from "@/utils/alert";

export default {
  name: "Namespace",
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
      read: 'nocursor',
      activeKey: 0,
      nodeStatus: true,
      nodeList: [],
      appInfo: '',
      iconList: ['cloud-server', 'deployment-unit', 'coffee', 'meh', 'fire', 'alert', 'robot'],
      meunKey: '',
      loading: true,
      loadingMore: true,
      showLoadingMore: true

    };
  },
  watch: {
    "$route.query.nodeId"(){
      this.getAppNamespace(this.$route.query.appId,this.$route.query.nodeId,this.$route.query.status)
    }
  },
  computed: {
    ...mapGetters(["asId"]),
    ...mapGetters(["appNodeId"]),
    ...mapGetters(["appNodeId"]),
    //...mapGetters(["nodeStatus"])
    releaseTime(){
      return this.$t('Namespace.lastReleaseTime') }
  },
  created() {
    VueEvent.$on('refresh-namespace',()=>{
      this.getAppNamespace(this.$route.query.appId,this.$route.query.nodeId,this.$route.query.status)
    })
    this.getAppNamespace(this.$route.query.appId,this.$route.query.nodeId,this.$route.query.status)

  },

  filters: {
    formatDate: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return  window.i18n.t('Namespace.lastReleaseTime')+formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    }
  },
  methods: {
    generateTitle,
    addShow() {
      this.$refs.add.showModal()
    },
    update(id) {
      this.$refs.update.showModal(id,this.$route.query.appId,this.$route.query.nodeId)
    },
    sync(id) {
      this.$refs.sync.showModal(id,this.$route.query.appId,this.$route.query.nodeId)
    },
    change(data) {
      this.ds = data;
    },
    showDeleteConfirm(id, name) {
      const _this = this;
      this.$confirm({
        title:  <div style="">{this.$t('Namespace.areDelete')}<span style="color: red; font-size: 20px">{name}</span></div>,
        content: this.$t('Namespace.confirmDelete'),
        okText: this.$t('Button.determine'),
        okType: 'danger',
        cancelText: this.$t('Button.cancel'),
        onOk() {
          deleteNamespace(id,Number(_this.$route.query.nodeId)).then(resp => {
            if(alertMessage(_this,resp.data)){
              _this.getAppNamespace(_this.$route.query.appId,_this.$route.query.nodeId,_this.$route.query.status);
            }
          });
        },
        onCancel() {
          console.log('Cancel');
        },
      });
    },
    getAppNamespace(appId,nodeId,st) {
      this.namespaceList=[];
      if(st==='true'){
        this.nodeStatus=true
        this.showLoadingMore=true;

        getAppNamespace({"appId":appId, "nodeId": nodeId}).then(response => {
          this.namespaceList = response.data.data.list
          this.status = response.data.data.node.status
          if (this.status) {
            response.data.data.list == null ? this.isNotData = true : this.isNotData = false
          }

        });

      }else{
        this.nodeStatus=false
      }
      this.showLoadingMore=false;
    },
    getTime() {
      var date1 = new Date();
      var year = date1.getFullYear();
      var month = date1.getMonth() + 1;
      var day = date1.getDate();
      var hours = date1.getHours();
      var minutes = date1.getMinutes();
      var seconds = date1.getSeconds();
      return year + month + day + hours  + minutes + seconds
    },
    showConfirm(id, name) {
      var date1 = new Date();

      let reName = formatTimeToStr(date1)+ '-RE';
      const _this = this;
      this.$confirm({
        okText: this.$t('Button.release'),
        width: '35%',
        title: this.$t('Namespace.releaseTitle'),
        content: h => <div style="">
          <span style="font-size: 16px">{this.$t('Namespace.List.namespaceName')}: &nbsp; &nbsp;   </span><span
            style="font-size: 16px;color:red">{name}</span>
          <br/>
          <span style="font-size: 16px">{this.$t('Namespace.releaseName')}: &nbsp; &nbsp;   </span><span
            style="font-size: 16px;color:#F5B400">{reName}</span>
        </div>,
        onOk() {
          createRelease({id: Number(id),nodeId:Number(_this.$route.query.nodeId), releaseName: reName}).then(resp => {
           if(alertMessage(_this,resp.data)){
             _this.getAppNamespace(_this.$route.query.appId,_this.$route.query.nodeId,_this.$route.query.status);
           }
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
