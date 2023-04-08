module github.com/dobyte/due-example

go 1.16

require (
	github.com/dobyte/due v0.0.17
	github.com/dobyte/due/locate/redis v0.0.0-20230401151017-df3a6121afbe
	github.com/dobyte/due/network/ws v0.0.0-20230401151017-df3a6121afbe
	github.com/dobyte/due/registry/consul v0.0.0-20230401151017-df3a6121afbe
	github.com/dobyte/due/transport/grpc v0.0.0-20230401151017-df3a6121afbe
	github.com/gin-gonic/gin v1.8.2
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/smallnest/rpcx v1.8.6
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	google.golang.org/grpc v1.50.1
)

replace (
	github.com/dobyte/due => ../due
	github.com/dobyte/due/eventbus/kafka => ../due/eventbus/kafka
	github.com/dobyte/due/locate/redis => ../due/locate/redis
	github.com/dobyte/due/network/ws => ../due/network/ws
	github.com/dobyte/due/registry/consul => ../due/registry/consul
	github.com/dobyte/due/transport/grpc => ../due/transport/grpc
	github.com/dobyte/due/transport/rpcx => ../due/transport/rpcx
)
