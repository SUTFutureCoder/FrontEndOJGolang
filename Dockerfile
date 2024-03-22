FROM golang:latest
MAINTAINER SutFutureCoder "FutureCoder@aliyun.com"
# init env
ENV PATH="/root/anaconda3/bin:${PATH}"
ENV GOPROXY https://goproxy.io
ENV GO111MODULE on
WORKDIR /root
RUN apt-get install -y wget
# install anaconda3 pytorch
RUN wget http://mirrors.tuna.tsinghua.edu.cn/anaconda/archive/Anaconda3-2021.05-Linux-x86_64.sh
RUN /bin/bash Anaconda3-2021.05-Linux-x86_64.sh -b
RUN conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free/
RUN conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/main/
RUN conda install pytorch -c pytorch
RUN conda install torchvision -c pytorch
RUN conda install ipython
RUN conda install matplotlib pandas seaborn scipy numpy
# install google-chrome
RUN apt-get update
RUN apt-get -f install -y
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN apt install -y ./google-chrome-stable_current_amd64.deb
# run golang
COPY . /root/go/FrontEndOJGolang
WORKDIR /root/go/FrontEndOJGolang
RUN go run main.go

