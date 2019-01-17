# fabric-sdk api

[FaricSDK](https://godoc.org/github.com/hyperledger/fabric-sdk-go/pkg/fabsdk)使用基本流程如下：

* A、使用配置实例化fabsdk实例

* B、使用fabsdk实例基于组织和用户创建上下文环境

* C、以上下文环境作为参数，使用New函数创建客户端实例。可以为需要的每个上下文环境创建一个客户端实例。

* D、使用每个客户端实例提供的功能函数编写业务逻辑，构建解决方案。

* E、调用fsbsdk.Close（）函数释放资源和缓存。

* pkg

    * pkg/fab/channel/transactor.go： 从channel中获取orderer信息，给orderer发送相关tx。
    
    * pkg/client/channel/chclient.go：创建channel的client，执行query或者invoke（execute）
    
    * pkg/fab/endpointconfig.go： 根据配置文件，初始化channel，peer和orderer。
    
    * pkg/client/channel/invoke/txnhandler.go： endorseHandler，创建和发送Tx。
    
    * pkg/fab/txn/proposal.go： proposal相关操作，创爱，签名，发送。
    
    * pkg/fab/peer/peer.go：peer的endpoint相关属性和方法实现
    
    * pkg/fab/orderer/orderer.go： orderer的相关相关属性和方法实现
    
    * pkg/fab/peer/peerendorser.go： endorse相关属性和方法实现
    
    * pkg/fab/txn/txn.go： 交易相关操作（orderer），调用grpc方法


[Channel](https://godoc.org/github.com/hyperledger/fabric-sdk-go/pkg/client/channel)使用流程如下：

* A、准备通道客户端上下文

* B、创建通道客户端

* C、执行链码

* D、查询链码

[Fabric账本](https://cloud.tencent.com/info/b6e95f56948825222f260fc3f273ab74.html)存储原理

Fabric区块链网络中，每个通道都有其账本，每个Peer节点都保存着其所加入通道的账本，Peer节点的账本包含如下数据：

* A、账本编号，用于快速查询存在哪些账本

* B、账本数据，用于区块数据存储

* C、区块索引，用于快速查询区块／交易

* D、状态数据，用于最新的世界状态数据

* E、历史数据：跟踪键的历史

![peer账本存储](pkg/fabric账本.jpg)



## 实际中的运行架构

* APP：

    代表一个客户端（CLI）或SDK，作用是创建交易并获取到足够的背书之后向Orderer排序服务节点提交交易请求（Peer与Orderer节点提供了gRPC远程访问接口，供客户端调用）。

* CA：
    
    负责对网络中所有的证书进行管理(对Fabric网络中的成员身份进行管理), 提供标准的PKI服务。

* MSP（Member Service Provider）：

    为客户端和Peer提供证书的系统抽象组件。

* Channel：将一个大的网络分割成为不同的私有"子网"。

    通道的作用：通道提供一种通讯机制，将peers和orderer连接在一起，形成一个具有保密性的通讯链路（虚拟）， 进行数据隔离。
    
    要加入通道的每个节点都必须拥有自己的通过成员服务提供商（MSP）获得的身份标识。

* Orderer：

    对客户端提交的交易请求进行排序，之后生成区块广播给通道内的所有peer节点。

* Org1：

    代表联盟中的某一个组织（一个联盟中可以多个不同的组织组成）。

* Peer：
    
    表示组织中的节点；Peer节点以区块的形式从Orderer排序服务节点接收有序状态更新，维护状态和账本。在Fabtic网络环境中 Peer 节点可以划分为如下角色：
    
    * Endorsing peer：根据指定的策略调用智能合约，对结果进行背书， 返回提案响应到客户端。
    
    * Committing peer：验证数据并保存至账本中。
    
    * Anchor peer：跨组织通信。
    
    * Leading peer：作为组织内所有节点的的代表连接到Orderer排序服务节点, 将从排序服务节点接收到的批量区块广播给组织内的其它节点。
    
    网络中只有部分节点为背书节点； 网络中所有Peer节点为账本节点。

* Chaincode：

    链式代码，简称链码；运行在容器中，提供相应的API与账本数据进行交互。

* Ledger：
    
    是由排序服务构建的一个全部有序的交易哈希链块，保存在所有的peer节点中。
    账本提供了在系统运行过程中发生的可验证历史，它包含所有成功的状态更改（有效交易）和不成功的状态更改（无效交易）。


# 接口

* 1.新增记录接口：

我会给你一条以工单为主的信息（包括工单号，工单状态，订单号，供方id, 需方id）,你这边保存到库，返回状态给我

* 2.记录更新接口：

我会给你四个参数（工单号，用户id, 是供方还是需方，工单状态），你这边根据工单号和供方id/需方id去更新工单状态，返回状态给我

* 3.记录查询接口：

我会给你一组工单号，用户id, 是供方还是需方这三个参数，你这边根据工单号和供方id/需方id去匹配 非这组工单号 的工单记录，并返回数据给我


# 参考


[Fabric SDK配置](https://cloud.tencent.com/info/c8da23c2b40acf91744b3ae7d8eb503b.html)






# 待办事项

* TODO

    未调试、yaml文件


























