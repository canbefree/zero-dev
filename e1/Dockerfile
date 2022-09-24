# protoc 
FROM ubuntu as stage_1
RUN sed -i s@deb.debian.org@mirrors.tuna.tsinghua.edu.cn@g /etc/apt/sources.list
RUN apt-get update && apt-get install -y zip wget
COPY install_proto.sh .
RUN /bin/bash install_proto.sh


FROM centos as stage_2
RUN curl https://github.com/protocolbuffers/protobuf/releases/download/v21.6/protoc-21.6-linux-x86_64.zip -o liunx.zip
