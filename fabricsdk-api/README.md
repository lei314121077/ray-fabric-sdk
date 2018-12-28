# fabric-sdk api

Fabsdk使用基本流程如下：

* A、使用配置实例化fabsdk实例

* B、使用fabsdk实例基于组织和用户创建上下文环境

* C、以上下文环境作为参数，使用New函数创建客户端实例。可以为需要的每个上下文环境创建一个客户端实例。

* D、使用每个客户端实例提供的功能函数编写业务逻辑，构建解决方案。

* E、调用fsbsdk.Close（）函数释放资源和缓存。




# 参考

[官方API](https://godoc.org/github.com/hyperledger/fabric-sdk-go/pkg/fabsdk)

[Fabric SDK配置](https://cloud.tencent.com/info/c8da23c2b40acf91744b3ae7d8eb503b.html)


























