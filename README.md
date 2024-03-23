# FrontEndOJGolang
## 面向前端开发者的在线评测系统——后端接口

随着在线评测系统（OnlineJudge）相关技术的不断发展，在线评测系统越来越多地应用于数据结构和算法教学、程序设计竞赛以及企业在线面试领域。

然而大多数的在线评测系统仅能评测在服务器环境直接运行的程序语言（如Java、C、C++），尚缺少对前端开发者技能练习和竞赛需求的支持。

因此，为解决上述问题，需要设计一种面向前端开发者的在线评测系统，该系统能够在服务器环境渲染、执行前端代码，并提供前端界面相似度对比等功能。

## 体验[不保证体验服务器稳定]

https://oj.project256.com/login

账号:test

密码:123456

## 代码仓库

本系统主要分为三个子项目

[前端页面](https://github.com/SUTFutureCoder/FrontEndOJFrontEnd) 基于Vue NPM WebSocket

[后端接口](https://github.com/SUTFutureCoder/FrontEndOJGolang) 基于Golang MySQL WebSocket

[评测核心](https://github.com/SUTFutureCoder/FrontEndOJGolang) 基于Golang ChromeDP WebSocket

## 运行方法

### 安装部署

请参考DockerFile

```bash
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
```

### MySQL结构

导入 caroline.sql



### 配置

conf/app.ini



### 默认密码

#### 用户名

管理员

#### 密码

654321



## 效果预览

### 题目列表

![题目列表](https://github.com/SUTFutureCoder/FrontEndOJGolang/blob/master/sample/image40.png?raw=true)

### 题目浏览及提交

![题目浏览及提交](https://github.com/SUTFutureCoder/FrontEndOJGolang/blob/master/sample/image41.png?raw=true)

### 提交结果

使用WebSocket实现实时通信

![提交结果界面](https://github.com/SUTFutureCoder/FrontEndOJGolang/blob/master/sample/image42.png?raw=true)

### 界面相似度对比

基于ResNet对比界面相似度

![界面相似度对比](https://github.com/SUTFutureCoder/FrontEndOJGolang/blob/master/sample/image43.png?raw=true)

### 竞赛列表

![竞赛列表界面](https://github.com/SUTFutureCoder/FrontEndOJGolang/blob/master/sample/image44.png?raw=true)

### 竞赛榜单

![竞赛榜单界面](https://github.com/SUTFutureCoder/FrontEndOJGolang/blob/master/sample/image45.png?raw=true)

### 个人信息面板

![个人信息面板](https://github.com/SUTFutureCoder/FrontEndOJGolang/blob/master/sample/image46.png?raw=true)

### 题目配置界面

![题目配置界面](https://github.com/SUTFutureCoder/FrontEndOJGolang/blob/master/sample/image48.png?raw=true)

## 代码流程

### 代码用例图

![代码用例图](https://github.com/SUTFutureCoder/FrontEndOJGolang/blob/master/sample/image21.png?raw=true)

### 系统功能结构图

![系统功能结构图](https://github.com/SUTFutureCoder/FrontEndOJGolang/blob/master/sample/image22.png?raw=true)

### E-R图

![E-R图](https://github.com/SUTFutureCoder/FrontEndOJGolang/blob/master/sample/image25.png?raw=true)

### 活动图

![核心活动图](https://github.com/SUTFutureCoder/FrontEndOJGolang/blob/master/sample/image30.png?raw=true)

### 评测流程图

![核心流程图](https://github.com/SUTFutureCoder/FrontEndOJGolang/blob/master/sample/image31.png?raw=true)

### 延迟评测流程图

当需要评测动画或是动态效果，需要等待一定时间再执行时，系统支持延迟一定时间再评测。

![延迟评测核心流程图](https://github.com/SUTFutureCoder/FrontEndOJGolang/blob/master/sample/image34.png?raw=true)

### 界面相似度对比流程图

![界面相似度对比流程图](https://github.com/SUTFutureCoder/FrontEndOJGolang/blob/master/sample/image38.png?raw=true)

## 说明

### 关于使用效果

此系统为作者攻读硕士学位所设计，时间较为久远，暂不保证使用效果，可供代码参考。

### 关于名字

作者当时重玩Protal2上头，所以起了caroline的名字，测试也用lab指代。
