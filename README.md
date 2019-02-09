
#### basic-network
basic-network和官方fabric-sample/basic-network基本一样
```shell
cd basic-network
./generate.sh
然后修改docker-compose.yml
FABRIC_CA_SERVER_CA_CERTFILE
FABRIC_CA_SERVER_CA_KEYFILE
```

#### nodejs-api
```shell
cd nodejs-api
./startFabric.sh
会调用basic-network/start.sh，启动docker容器，安装channel，把peer节点并加入channel
然后会安装链码，初始化链码
./enrollAdmin.sh 会创建一个叫admin的管理员用户，并把加密证书保存在 ./hfc-key-store
./registerUser.sh 会创建一个叫user1的普通用户，并把加密证书保存在 ./hfc-key-store
```

#### chaincode
```shell
链码目录
```
