#!/usr/bin/env bash

#######前端打包############
echo "前端打包"
cd ../flying-ui && yarn install
yarn build && mv dist/ ../flying-admin/dist

########这里可以改成自己的镜像仓库##########
cd ../flying-admin/
echo 'flying-admin 镜像打包'
docker build -t registry.cn-qingdao.aliyuncs.com/zcj-oss/flying-admin:latest .
docker push registry.cn-qingdao.aliyuncs.com/zcj-oss/flying-admin:latest
docker rmi registry.cn-qingdao.aliyuncs.com/zcj-oss/flying-admin:latest
echo '镜像已清理'

cd ../flying-config/
echo 'flying-config 镜像打包'
docker build -t registry.cn-qingdao.aliyuncs.com/zcj-oss/flying-config:latest .
docker push registry.cn-qingdao.aliyuncs.com/zcj-oss/flying-config:latest
docker rmi registry.cn-qingdao.aliyuncs.com/zcj-oss/flying-config:latest
echo '镜像已清理'
