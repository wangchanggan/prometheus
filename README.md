# Prometheus源码分析

Source Code From
https://github.com/prometheus/prometheus/archive/refs/tags/v2.24.0.zip

## 目录

-   [Prometheus源码分析](#prometheus源码分析)
    -   [目录](#目录)
    -   [源码目录结构说明](#源码目录结构说明)
    
## 源码目录结构说明
| 源码目录 | 说明 | 备注 |
| :----: | :---- | :---- |
| cmd/ | prometheus目录中的main.go是整个程序的入口，promtool目录是规则校验工具promtool的源码目录 | |
| config/ | 用于管理YAML配置文件的加载、解析及常用配置结构的定义 | |
| console_libraries/ | 目录中的menu.lib和prom.lib是prometheus内置的基本界面组件，在自定义模板中可直接使用 | |
| consoles/和template/ | 用于管理prometheus控制台，还可以通过自定义模板增加外部的访问服务 | |
| discovery/ | prometheus的服务发现模块，用于发现scrape服务和告警服务 | |
| notifier/ | prometheus的通知管理模块，规则运算产生告警后将告警信息通过notifier发送给发现的告警服务 | |
| pkg/| prometheus的基础元素（如维度、时间戳、字节池等） | |
| prompb/ | 定义了三种协议，分别为远程存储协议、rpc通信协议和types协议 | |
| promql/ | 目录为规则计算的具体实现，根据载入的规则进行规则计算，并生成告警指标 | |
| rules/ |  prometheus的规则管理模块，用于实现规则加载、计算调度和告警信息的回调 | |
| scrape/ | 负责监控对象的指标拉取 | |
| scripts/ | 跟踪生成protobuf代码所需工具的依赖关系 | |
| storage/ | prometheus的指标存储模块，有remote（远程存储）、tsdb（本地存储）两种类型 | |
| tsdb/ | 本地存储模块 | |
| util/ | 工具类 | |
| vendor/| 第三方依赖包 | |
| web/| prometheus Web服务模块 | |