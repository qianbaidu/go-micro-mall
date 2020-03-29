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


### 运行服务支持组件：jaeger、consul、mysql
```
cd docker
docker-compose -f docker-compose.yml up -d
```
- jaeger web Ul: http://localhost:16686
- consul web Ul: http://localhost:8500/ui/dc1/services

### 日志告警系统 (elasticsearch + kibana + elasticAlerter)
```
cd docker
docker-compose -f docker-compose-monitor.yml up -d
```
- elasticsearch : http://localhost:9200
- kibana 日志查询: http://localhost:5601
- elasticAlerter配置修改路径： `docker/volumes/jhipster-alerter`

### 运行服务
- 1.启动配置中心
```
cd config-srv
make run
```
- 2.启动用户服务
```
cd user
make run
```
- 3.启动hello server
```
cd hello
make run
```
- 4.启动api服务
```
cd api
make run
```
- 5.启动gateway
```
cd gateway
make run
```

### 服务测试
-  1.测试api
```
curl -XGET http://localhost:9090/user/test
```
- 2.测试gateway
```
curl -XGET http://localhost:8080/user/test
```
-  3.测试register
```
curl  http://127.0.0.1:8080/user/register -X POST   -H 'Content-Type:application/json'  -d '{"phone":"12345","password":"12345"}'
```
```
{"error":"","success":true,"message":"register success","data":""}
```

- 4.测试login
```
curl  http://127.0.0.1:8080/user/login -X POST   -H 'Content-Type:application/json'  -d '{"phone":"12345","password":"12345"}'
```
```
{"error":"","success":true,"message":"login success","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiIxMjM0NSIsImV4cCI6MTU4NTU2MjI5MCwiaWF0IjoxNTg1NDc1ODkwLCJpc3MiOiJnby5taWNyby5zcnYuYXV0aCJ9.GbFmfSXOGCm-mt6htyEVfO__pSIfDZg-eQCjPRHi5sQ"}}

```




