module github.com/kulycloud/test-service

go 1.15

require (
	github.com/kulycloud/common v1.0.0
	github.com/kulycloud/protocol v1.0.0
)

replace github.com/kulycloud/common v1.0.0 => ../common

replace github.com/kulycloud/protocol v1.0.0 => ../protocol
