<template>
    <a-layout-sider width="400px" style="background: #fff;">

        <a-list style="width: 350px;margin: 0 auto" item-layout="horizontal" size="small" :data-source="data" header="环境信息">
          <a-list-item  slot="renderItem" slot-scope="item">

            <a-list-item-meta  style="width: 10px">

              <dev  slot="avatar"><img  title="aa" width="40px" :src="item.status ? zhen:jia" alt="zzzzzz">

              </dev>

            </a-list-item-meta>

            <a-list-item-meta
            >

              <dev slot="title" :href="item.name">{{ item.name }}</dev>

            </a-list-item-meta>

            <a-list-item-meta>
              <a slot="description">{{item.url}}</a>
            </a-list-item-meta>

          </a-list-item>
        </a-list>


    </a-layout-sider>
</template>

<script>
const data = [
  {
    name: 'dev',
    url: 'http://192.169.1.241:8881/',
    status: false
  },
  {
    name: 'pro',
    url: 'http://192.169.1.241:8881/',
    status: true
  },
  {
    name: 'dev',
    url: 'http://192.169.1.241:8881/',
    status: false
  },
  {
    name: 'dev',
    url: 'http://192.169.1.241:8881/',
    status: true
  },
];
export default {
  data() {
    return {
      data,
      zhen: require('../../image/tubiao-lianjiezhuangtai-zhengchang.png'),
      jia: require('../../image/tubiao-lianjiezhuangtai-zhengchang-.png'),
    };
  },
};
</script>

<style scoped>
.ant-list-item-meta{

}
</style>