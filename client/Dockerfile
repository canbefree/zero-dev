# protoc 
FROM ubuntu as stage_download_protoc
RUN sed -i s@deb.debian.org@mirrors.tuna.tsinghua.edu.cn@g /etc/apt/sources.list
RUN apt-get update && apt-get install -y zip wget
COPY .build/install_proto.sh .
RUN /bin/bash install_proto.sh

# proto-gen-go proto-gen-grpc 
FROM golang:1-bullseye AS stage_gen_go_gateway
# ENV GOPROXY=https://goproxy.cn
COPY .build/install_grpc_gateway.sh .
RUN /bin/bash install_grpc_gateway.sh

# google proto
FROM golang:1-bullseye as stage_download_google_apis
RUN git config --global url."https://ghproxy.com/https://github.com".insteadOf "https://github.com"
RUN git clone  https://github.com/canbefree/google-apis.git google
RUN mv google /usr/local/include/

FROM golang:1-bullseye AS build
ARG GOPROXY
# ENV GOPROXY=https://goproxy.cn
# ARG GIT_USERNAME=example 
# ARG GIT_EMAIL=example@mail.com


WORKDIR /workspace

RUN go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/ramya-rao-a/go-outline@latest \
    && go install -v github.com/fatih/gomodifytags@latest  \
    && go install -v github.com/cweill/gotests/gotests@latest \
    && go install -v github.com/josharian/impl@latest  \
    && go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/go-delve/delve/cmd/dlv@latest \
    && go install -v github.com/davidrjenni/reftools/cmd/fillstruct@latest

COPY go.mod .
COPY go.sum .
RUN go mod download



ENV CGO_ENABLED=0
COPY --from=stage_download_protoc /usr/bin/protoc /usr/bin
COPY --from=stage_gen_go_gateway /go/bin/proto* /go/bin
# docker COPY 和 cp 在拷贝目录的时候有差异
COPY --from=stage_download_google_apis /usr/local/include/google /usr/local/include/google

COPY . .
RUN go build -o server .


FROM build AS development

# RUN git config --global user.name "${GIT_USERNAME}"
# RUN git config --global user.email "${GIT_EMAIL}"

# ENV HTTP_PROXY="http://proxy:3010"
# ENV HTTPS_PROXY="https://proxy:3012"

CMD ["go", "run", "main.go"]

FROM alpine:3.12
EXPOSE 8000
COPY --from=build /workspace/server /server
CMD ["/server222"]
