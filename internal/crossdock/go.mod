module go.uber.org/yarpc/internal/crossdock

go 1.13

require (
	github.com/apache/thrift v0.13.0
	github.com/crossdock/crossdock-go v0.0.0-20160816171116-049aabb0122b
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/stretchr/testify v1.4.0
	github.com/uber/jaeger-client-go v2.22.1+incompatible
	github.com/uber/tchannel-go v1.22.2
	go.uber.org/fx v1.10.0
	go.uber.org/multierr v1.4.0
	go.uber.org/thriftrw v1.29.2
	go.uber.org/yarpc v1.42.1
	go.uber.org/zap v1.13.0
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	google.golang.org/grpc v1.28.0
)

replace go.uber.org/yarpc => ../..

// Pin to v0.10.0; 0.11 added context arguments which breaks TChannel Go.
//
// We're pinning to hash because before 0.12, Apache Thrift did not include a
// 'v' prefix for their SemVer releases, which is incompatible with Go
// modules.
replace github.com/apache/thrift => github.com/apache/thrift v0.0.0-20161221203622-b2a4d4ae21c7
