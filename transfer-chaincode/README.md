#  交易中心-智能合约 [上海数据交易中心](https://fe.chinadep.com/user/public/helpDocument.html) 

[参考](https://www.chaindesk.cn/witbook/11/117)

链是由一个通道 + 一个共享账本 + N个 peer 节点组成；

## 上海数据交易中心应用

营销应用
 
  * 数据单品主要包括：
  
    入口特征、学识程度、电商购买意向、应用使用偏好列表，以及包含汽车、母婴、金融等特有数据单品

征信应用

  * 主要数据单品为身份要素验证：
      
      三要素：（身份证、姓名、手机）
      
      四要素：（身份证、姓名、手机、银行卡）验证
      
      后续会有扩展
       
## 如何接入
 
  通过SDK的Peer操作智能合约相关逻辑

## 数据处理

  CouchDB数据库
    
# 如何用智能合约去实现
  
  应用程序通过向区块链网络发送交易来调用智能合约，从而操作账本中的状态。

  完整交易流程参见下图
  
  ![交易流程图](https://pic4.zhimg.com/80/v2-155a55394f1d508c6c10bdf73aa3084f_hd.jpg)
  
  ChainCode角色图
  
  ![ChainCode角色图](https://pic3.zhimg.com/80/v2-30fe8099f29c036254b59f3a0ac2147e_hd.jpg)
  
  
  * 步骤1、 启动网络

  * 步骤2、Fabric Chaincode 负责实现具体的只能合约业务逻辑方法
   
  * 步骤3、本地通过SDK实现具体的操作
       
      * 后台语言JAVA通过[SDK](https://github.com/hyperledger/fabric-sdk-java)插入区块，获取区块信息
       
      * 后台语言GO通过[SDK](https://github.com/hyperledger/fabric-sdk-go)插入区块，获取区块信息
       
      * 后台语言Python通过[SDK](https://github.com/hyperledger/fabric-sdk-py)插入区块，获取区块信息 
    

## 链码功能
   
   Peer节点主要功能是调用链码执行交易和记账，其中交易执行由背书节点的链码负责，记账功能由记账节点负责。

   ![fabric](/home/ray/go/data-transfer-chaincode/transfer-chaincode/pkg/chaincode.jpg)

   链码与Peer节点的交互过程如下：
   
   * A、链码通过gRPC与Peer节点交互，当Peer节点收到客户端的交易提案请求后，会发送一个链码消息对象（包含交易提案信息、调用者信息）给对应的链码。
   
   * B、链码调用Invoke方法，通过发送获取数据（GetState）和写入数据（PutState）消息，向Peer节点获取账本状态信息和发送预提交状态。
   
   * C、链码发送模拟执行结果给Peer节点，Peer节点对交易提案和模拟执行结果进行背书签名。  
  
## 部署与应用
  
docker

* 执行Dockerfile生成docker镜像
  
  ```bash
     docker build -t data-transfer-center
  ```
   
  
 
    


















