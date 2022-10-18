FROM golang:1-bullseye AS build
# ENV GOPROXY=https://goproxy.cn,direct

# 加速构建
WORKDIR /workspace
COPY go.mod .
COPY go.sum .
RUN go mod download

ENV CGO_ENABLED=0

COPY . .
RUN go build -o server .


FROM build AS development

RUN go install github.com/josharian/impl@latest

RUN go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/ramya-rao-a/go-outline@latest \
    && go install -v github.com/fatih/gomodifytags@latest  \
    && go install -v github.com/cweill/gotests/gotests@latest \
    && go install -v github.com/josharian/impl@latest  \
    && go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/go-delve/delve/cmd/dlv@latest \
    && go install -v github.com/davidrjenni/reftools/cmd/fillstruct@latest

RUN git config --global user.name "${GIT_USERNAME}"
RUN git config --global user.email "${GIT_EMAIL}"


CMD ["go", "run", "main.go"]

FROM alpine:3.12
EXPOSE 8000
COPY --from=build /workspace/server /server
CMD ["/server"]
