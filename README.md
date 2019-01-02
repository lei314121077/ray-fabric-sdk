# 数据交易中心-智能合约


## 目录

```bash

    data-transfer-chaincode
    ├── data-transfer-share-libray     # 公共包
    │   ├── bin
    │   ├── pkg
    │   └── src                        # 自定义包
    │       └── ray
    │       
    ├── fabricsdk-api                  # api
    │   ├── bin
    │   ├── key.sh
    │   ├── main.go
    │   ├── pkg
    │   ├── README.md
    │   ├── server.rsa.crt
    │   ├── server.rsa.key
    │   └── src
    │     
    ├── README.md                      # markdown 描述文档
    │     
    └── transfer-chaincode             # chain code包
        ├── bin
        ├── build.sh
        ├── data-transfer-chaincode
        ├── docker-compose.yml
        ├── Dockerfile
        ├── main.go
        ├── pkg
        ├── README.md
        └── src
       
```


## 版本&环境

* Hyperledger Fabric v1.0.0

```bash
    
    # 更新版本
    go get -u github.com/hyperledger/fabric

    # 克隆版本
    git clone -b v1.0.0 https://github.com/hyperledger/fabric.git

```

* Docker 
    
    ```bash
      docker --version
      Docker version 18.09.0
    ```
    
* docker-compose
    
    ```bash
      docker --version
      docker-compose version 1.17.0
    ```
    
* golang 
    
    ```bash
      go version
      go version go1.10 
    ```    

* 拉取镜像
    
    ```bash
    # 上DockerHub官方镜像网站： https://hub.docker.com/u/hyperledger/?page=1
    docker pull hyperledger/fabric-tools:x86_64-1.0.0
    docker pull hyperledger/fabric-couchdb:x86_64-1.0.0
    docker pull hyperledger/fabric-kafka:x86_64-1.0.0
    docker pull hyperledger/fabric-zookeeper:x86_64-1.0.0
    docker pull hyperledger/fabric-orderer:x86_64-1.0.0
    docker pull hyperledger/fabric-javaenv:x86_64-1.0.0
    docker pull hyperledger/fabric-ccenv:x86_64-1.0.0
    docker pull hyperledger/fabric-ca:x86_64-1.0.0
    docker pull hyperledger/fabric-baseos:x86_64-0.3.1
    docker pull hyperledger/fabric-baseimage:x86_64-0.3.1
    docker pull hyperledger/fabric-membersrvc:latest
    docker images
    ```


## 部署




## 业务流程



## 其它

































