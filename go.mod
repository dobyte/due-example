module github.com/dobyte/due-example

go 1.16

require google.golang.org/protobuf v1.28.1

require (
	github.com/coreos/go-systemd/v22 v22.4.0 // indirect
	github.com/dobyte/due v0.0.7
	github.com/dobyte/due/locate/redis v0.0.0-20221201102446-6eff2a223daa
	github.com/dobyte/due/network/ws v0.0.0-20221201102446-6eff2a223daa
	github.com/dobyte/due/registry/etcd v0.0.0-20221201102446-6eff2a223daa
	github.com/dobyte/due/transport/rpcx v0.0.0-20221201102446-6eff2a223daa
	go.etcd.io/etcd/client/v3 v3.5.5 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/net v0.1.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	google.golang.org/genproto v0.0.0-20221018160656-63c7b68cfc55 // indirect
)
