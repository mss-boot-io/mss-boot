module oauth2

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-oauth2/oauth2/v4 v4.4.3
	github.com/go-session/session v3.1.2+incompatible
	github.com/google/uuid v1.2.0
	github.com/mss-boot-io/mss-boot v0.0.0-20211230092308-0cee71cc9e5d
	github.com/sanity-io/litter v1.5.2
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.5.1
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

replace github.com/mss-boot-io/mss-boot => ../../
