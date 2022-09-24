# CA key

openssl genrsa -out local/boot.key 2048
# CA 证书
openssl req -x509 -new -key local/boot.key -out local/boot.pem -days 3650

# 服务器KEY
openssl genrsa -out web/app.key 2048

# 服务器认证请求
证书认证请求并不是证书，需要CA的私钥进行签名之后方是证书
openssl req -new -key  web/app.key -out web/app.csr


# 创建证书附加用途文件

```
[root@server ~/certs]# vim web/app.ext
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth, clientAuth
subjectAltName=@SubjectAlternativeName

[ SubjectAlternativeName ]
IP.1=192.168.1.1
IP.2=192.168.1.2

```

```
[root@server ~/certs]# vim web/app.ext
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth, clientAuth
subjectAltName=@SubjectAlternativeName

[ SubjectAlternativeName ]
DNS.1=medusa.com
DNS.2=*.medusa.com
```


# 服务器证书 (证书时间不得长于13个月)
openssl x509 -req -in web/app.csr -CA local/boot.pem -CAkey local/boot.key -CAcreateserial -out web/app.crt -days 365 -sha256 -extfile web/app.ext