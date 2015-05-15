export GOPATH=$GOPATH:`pwd`
go clean
go build -v -o google_auth_proxy
