module github.com/dobyte/due-example

go 1.16

require (
	github.com/dobyte/due v0.0.9
	github.com/dobyte/due/locate/redis v0.0.9
	github.com/dobyte/due/network/ws v0.0.9
	github.com/dobyte/due/registry/consul v0.0.9
	github.com/dobyte/due/transport/rpcx v0.0.9
	github.com/gin-gonic/gin v1.8.2
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.8 // indirect
)

replace (
	github.com/dobyte/due => ../due
	github.com/dobyte/due/locate/redis => ../due/locate/redis
	github.com/dobyte/due/network/ws => ../due/network/ws
	github.com/dobyte/due/registry/consul => ../due/registry/consul
	github.com/dobyte/due/transport/rpcx => ../due/transport/rpcx
)
