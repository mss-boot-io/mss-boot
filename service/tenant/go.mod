module tenant

go 1.16

require (
	github.com/google/uuid v1.2.0
	github.com/lwnmengjing/core-go v0.0.0-00010101000000-000000000000
	github.com/lwnmengjing/mss-boot v0.0.0-00010101000000-000000000000
	github.com/sanity-io/litter v1.5.1
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
	gorm.io/gorm v1.21.11
)

replace (
	github.com/lwnmengjing/core-go => ../../../core-go
	github.com/lwnmengjing/mss-boot => ../../
)
