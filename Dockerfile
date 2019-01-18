# 指定#源镜像
FROM golang:1.11

#设置工作目录
WORKDIR $HOME/go/data-transfer-chaincode


#将服务器的go工程代码加入到docker容器中
ADD transfer-chaincode $HOME/go/data-transfer-chaincode
ADD transfer-libray $HOME/go/data-transfer-chaincode/data-transfer-share-libray
ADD sdk-api $HOME/go/data-transfer-chaincode/farbicsdk-api

RUN echo "构建环境！:"
RUN cd $HOME/go/data-transfer-chaincode && ./run.sh


#go构建可执行文件
#RUN ./build

#执行go构建的可执行文件
#CMD ["./data-transfer-chaincode"]

#暴露端口
EXPOSE 8000

#最终运行docker的命令
#ENTRYPOINT  ["./data-transfer-chaincode"]

