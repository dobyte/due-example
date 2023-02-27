module github.com/dobyte/due-example

go 1.16

require (
	github.com/dobyte/due v0.0.11
	github.com/dobyte/due/eventbus/nats v0.0.0-20230224062538-c72aa90db9f2
	github.com/dobyte/due/eventbus/redis v0.0.0-20230224062538-c72aa90db9f2 // indirect
	github.com/dobyte/due/locate/redis v0.0.0-20230221014826-f2747a572641
	github.com/dobyte/due/network/ws v0.0.0-20230221014826-f2747a572641
	github.com/dobyte/due/registry/consul v0.0.0-20230224062538-c72aa90db9f2
	github.com/dobyte/due/transport/rpcx v0.0.0-20230221014826-f2747a572641
	github.com/gin-gonic/gin v1.8.2
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.8 // indirect
)
