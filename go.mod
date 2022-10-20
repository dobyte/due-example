module github.com/dobyte/due-example

go 1.17

require google.golang.org/protobuf v1.28.1

require (
	github.com/BurntSushi/toml v1.2.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dobyte/due v0.0.4
	github.com/dobyte/due/locate/redis v0.0.4
	github.com/dobyte/due/network/ws v0.0.4
	github.com/dobyte/due/registry/etcd v0.0.4
	github.com/dobyte/due/transport/grpc v0.0.4
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.etcd.io/etcd/api/v3 v3.5.4 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.4 // indirect
	go.etcd.io/etcd/client/v3 v3.5.4 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/net v0.0.0-20220524220425-1d687d428aca // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.0.0-20220908164124-27713097b956 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220525015930-6ca3db687a9d // indirect
	google.golang.org/grpc v1.50.1 // indirect
)

replace (
	github.com/dobyte/due => ../due
	github.com/dobyte/due/locate/redis => ../due/locate/redis
	github.com/dobyte/due/network/ws => ../due/network/ws
	github.com/dobyte/due/registry/etcd => ../due/registry/etcd
	github.com/dobyte/due/transport/grpc => ../due/transport/grpc
)
