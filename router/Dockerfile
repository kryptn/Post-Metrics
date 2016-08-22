from golang:1.7-alpine

env GOBIN /go/bin

copy . /

run apk --update add git \
 && cd src \
 && go get \
 && go build router.go \
 && mv router ../bin


cmd ["router"] 
