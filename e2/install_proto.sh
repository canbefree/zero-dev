
cd /tmp
if arch 2>&1 | grep 'aarch64' > /dev/null
then
    wget https://github.com/protocolbuffers/protobuf/releases/download/v21.6/protoc-21.6-linux-aarch_64.zip
    unzip protoc-21.6-linux-aarch_64.zip
    chmod +x bin/protoc
    mv bin/protoc /usr/bin
else
    wget https://github.com/protocolbuffers/protobuf/releases/download/v21.6/protoc-21.6-linux-aarch_64.zip
    unzip protoc-21.6-linux-aarch_64.zip
    chmod +x bin/protoc
    mv bin/protoc /usr/bin
fi
