module goedd

go 1.26

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.30.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stackus/dotenv v0.0.0-20221206033122-02295762494b
	github.com/stackus/errors v0.1.8
	golang.org/x/sync v0.22.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.70.0
	google.golang.org/grpc v1.83.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.6.2
	google.golang.org/protobuf v1.36.11
)

require (
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/net v0.57.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.40.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260803160001-6ac0973c030d // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260803160001-6ac0973c030d // indirect
)
