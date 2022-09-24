FROM golang:1-bullseye AS build


RUN  mkdir -p -m 0700 ~/.ssh && ssh-keyscan git.myscrm.cn >> ~/.ssh/known_hosts

RUN --mount=type=ssh git clone git@git.myscrm.cn:vip/incentive-core.git