# readMe
go-micro-mall 是一款基于go-micro搭建的微服项目
主要技术栈：go-micro、gin、hystrix、jaeger、consul、elasticsearch、grpc、jwt

## 功能介绍
- 核心功能组件：
    - api gateway
    - 注册中心 （consul）
    - 配置中心
    - 熔断
    - 链路追踪
    - jwt
    - metrics
    - 集成日志推送ES
    - 集成Swagger

### 相关组件运行
- jaeger
```
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.17
```
- jaeger web Ul: http://localhost:16686

- 服务注册中心 consul
```
consul agent -dev
```
- consul web Ul: http://localhost:8500/ui/dc1/services


- 日志告警系统 (elasticsearch + kibana + elasticAlerter)
```
cd micro-log-monitor
docker-compose -f docker-compose.yml up -d
```

- 生成proto
```
protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. *.proto
```