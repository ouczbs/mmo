module mmo

go 1.14

replace mmo => ./mmo

replace github.com/ouczbs/Zmin => ../Zmin

require (
	github.com/golang/protobuf v1.4.2
	github.com/ouczbs/Zmin v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.3.5
	google.golang.org/protobuf v1.25.0
)
