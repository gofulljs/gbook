# gbook

support gitbook 3.2.3, 其他版本自测， 解决 gitbook 太慢的问题

## 开启此项目初衷

gitbook 3.2.3 为历史项目，插件下载需要整体拉取，非常缓慢，也不支持单个拉取，这样`gbook`就应运而生了，

配置好后将有以下特性：

- 拉取插件特别快
- 不用频繁的切换 node 版本

## 依赖

- `node` <= 10, 建议 Node 版本 10.24.1
- `gitbook-cli` 必须
- 请自行设置好 npm 加速镜像，如淘宝镜像，建议使用 `nrm` 设置

## 稳定运行 gbook

将 BOOK_NODE_HOME 变量设置为需要的 node 版本路径，加入到系统变量中, 这样`gbook`会使用对应 node 版本执行命令了，不再受全局其他版本 node 影响。

## 主要功能

`gbook install`将是核心功能，你只需关注该命令即可

## 详细说明

参考 https://gofulljs.github.io/gbook/
