#!/usr/bin/env bash
# 私钥文件 tls.key  数字证书 tls.crt
openssl req -x509 -nodes -newkey rsa:2048 -keyout tls.key -out tls.crt -days 3650

# 服务端私钥与证书
#openssl生成的私钥中包含了公钥的信息，我们可以根据私钥生成公钥：

openssl rsa -in tls.key -out tls.key.public



## create server private key
#openssl genrsa -out server.key 2048
#
## create server request file ps myname is your server name
#openssl req -new -key server.key -out server.csr -subj "/C=CN/ST=BJ/L=beijing/O=myorganization/OU=mygroup/CN=myname"
#
## create client key
#openssl genrsa -out client.key 2048
#
## create client request file
#openssl req -new -key client.key -out client.csr -subj "/C=CN/ST=BJ/L=beijing/O=myorganization/OU=mygroup/CN=myname"
#







