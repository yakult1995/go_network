FROM centos:7

# 環境変数
ENV PATH /home/centos/go
ENV PATH /usr/local/go/bin:$GOPATH/bin:$PATH

# アップデート、インストール
RUN yum update -y
RUN yum install -y git wget

# Golangのインストール
RUN wget https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz
RUN tar vzfx go1.9.linux-amd64.tar.gz
RUN rm go1.9.linux-amd64.tar.gz
RUN mv go /usr/local/

# ライブラリのインストール
RUN go get github.com/golang/dep
RUN go get github.com/BurntSushi/toml
RUN go get golang.org/x/sys/unix
