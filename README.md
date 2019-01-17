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
        │   ├── build.sh
        │   ├── config-bak.yaml
        │   ├── config.yaml
        │   ├── crypto-config.yaml
        │   ├── main.go
        │   ├── pkg
        │   │   └── fabric账本.jpg
        │   ├── README.md
        │   └── src
        │        ├── demo
        │        ├── httpsdk
        │        └── order
        │  
        ├── transfer-chaincode             # chain code包
        │   ├── bin
        │   ├── build.sh
        │   ├── data-transfer-chaincode
        │   ├── docker-compose.yml
        │   ├── Dockerfile
        │   ├── main.go
        │   ├── pkg
        │   ├── README.md
        │   └── src 
        │     
        └── README.md                      # markdown 描述文档
             
   ```


## 版本&环境

* Hyperledger Fabric v1.4.0

    ```bash
        
        # 更新版本
        go get -u github.com/hyperledger/fabric
    
        # 克隆版本
        git clone -b v1.4.0 https://github.com/hyperledger/fabric.git
      
        # 或者
        git checkout -b v1.4.0
    
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
      go version go1.11 
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
    
    # 查看镜像
    docker images
    
    # 如果下载有失败的镜像, 可再次执行下面的命令重新下载。 
    
    cd github.com/hyperledger/fabric/blob/master/scripts
    sudo ./bootstrap.sh 1.2.0
  
    ```


## 部署




## 业务流程



## 其它

































