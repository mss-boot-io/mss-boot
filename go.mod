module github.com/lwnmengjing/mss-boot

go 1.16

require (
	github.com/favadi/protoc-go-inject-tag v1.1.0 // indirect
	github.com/golang/protobuf v1.5.0
	github.com/lwnmengjing/core-go v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
	gorm.io/driver/mysql v1.1.0
	gorm.io/driver/postgres v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.11
)

replace github.com/lwnmengjing/core-go => ../core-go
