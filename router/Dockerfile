from golang:1.7-alpine

env GOBIN /go/bin

run apk --update add git

copy . /

run cd src \
 && go get \
 && go build router.go \
 && mv router ../bin


cmd ["router"] 
