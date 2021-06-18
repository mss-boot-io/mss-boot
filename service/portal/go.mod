module portal

go 1.16

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/gin-gonic/gin v1.7.2
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.11
	github.com/lwnmengjing/core-go v0.0.0-00010101000000-000000000000
	github.com/lwnmengjing/mss-boot v0.0.0-00010101000000-000000000000
)

replace (
	github.com/lwnmengjing/core-go => ../../../core-go
	github.com/lwnmengjing/mss-boot => ../../
)
