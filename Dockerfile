# 指定#源镜像
FROM golang:1.10

#设置工作目录
WORKDIR $HOME/go/data-transfer-chaincode

#将服务器的go工程代码加入到docker容器中
ADD . $HOME/go/data-transfer-chaincode

#go构建可执行文件
RUN ./build

#执行go构建的可执行文件
CMD ["./data-transfer-chaincode"]

#暴露端口
EXPOSE 7272

#最终运行docker的命令
ENTRYPOINT  ["./data-transfer-chaincode"]

