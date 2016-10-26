Gateway
-------
Gateway is a API gateway based on http. It works at 7 layer.

# Features
* Traffic Control
* Circuit Breaker
* Loadbalance
* Routing based on URL
* API aggregation(support url rewrite)
* Backend Server heath check
* Use [fasthttp](https://github.com/valyala/fasthttp)
* Admin WEBUI

# Install
Gateway dependency [etcd](https://github.com/coreos/etcd)

## Compile from source
```
git clone https://github.com/fagongzi.git
cd $GOPATH/src/github.com/fagongzi/gateway
go build cmd/proxy/proxy.go
go build cmd/admin/admin.go
```

## Download binary file
[linux-64bit](http://7xtbpp.com1.z0.glb.clouddn.com/gateway-linux64.tar.gz)

# Online Demo

* admin
 
  http://demo-admin.fagongzi.win admin/admin

* proxy
  
  http://demo-proxy.fagongzi.win 

# Architecture
![](./images/arch.png)

## Components
Gateway has three component: proxy, admin, etcd.

### Proxy
The proxy provide http server. Proxy is stateless, you can scale proxy node to deal with large traffic.

### Admin 
The admin is a backend manager system. Admin also is a stateless node, you can use a Nginx node for HA. One Admin node can manager a set of proxy which has a same etcd prefix configuration.

### Etcd
The Etcd store gateway's mete data.

## Concept of gateway

* Server

  Server is a backend server which provide restfule json service.The server is the basic unit at gateway.

* Cluster

  Cluster is a set of servers which provide the same service. The Loadbalancer select a usable server to use.

* Aggregation

  Aggregation is a set of URLs that correspond to some clusters. A http request arrive proxy, the proxy dispatcher the request to specify clusters, then wait responses and merge to response client.
Notes, if your set a rewrite rule, it must container full request url, because proxy need set path value to query string or set query string to path value, to meet the demand that backend server url design. 

* Routing

  Routing is a approach to control http traffic to clusters. You can use cookie, query string, request header infomation in a expression for control.

# What gateway can help you
## Redefine your API URL
Your backend server provide some restful API, You can redefine the API URL that provide to API caller.Use this funcation you can provide beautiful APIs.  

## Dynamic URL & Aggregation
You can define a URL and configuration a URL set which you want to aggregation.

## Protect backend server
Gateway can use **Traffic Control** and **Circuit Breaker** functions to avoid backend crash by hight triffic.

## AB Test
Gateway's **Routing** fucntion can help your AB Test.

## Admin UI
![](./images/dashboard.png)

## Proxy UI
![](./images/proxy.png)

## Cluster UI
![](./images/cluster.png)

## Server UI
![](./images/server.png)

## Aggregation UI
![](./images/aggregations.png)

## Routing UI
![](./images/routing.png)

## Monitor UI
![](./images/metries.png)
