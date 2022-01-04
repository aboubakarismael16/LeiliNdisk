go run service/apigw/main.go --registry=consul

go run service/account/main.go --registry=consul

#go run service/download/main.go --registry=consul
#
#go run service/transfert/main.go --registry=consul
#
go run service/upload/main.go --registry=consul
#
#go run service/dbproxy/main.go --registry=consul