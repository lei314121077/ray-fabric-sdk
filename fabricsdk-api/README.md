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

作者：kamiSDY
链接：https://www.jianshu.com/p/8abae69d9ea9
來源：简书
简书著作权归作者所有，任何形式的转载都请联系作者获得授权并注明出处。

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







# 参考


[Fabric SDK配置](https://cloud.tencent.com/info/c8da23c2b40acf91744b3ae7d8eb503b.html)






# 待办事项

* TODO

    未调试、yaml文件


























