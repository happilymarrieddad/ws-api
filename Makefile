generate:
	go generate ./...

install.deps:
	go install github.com/onsi/ginkgo/v2/ginkgo
	go get github.com/onsi/gomega/...
	go install github.com/golang/mock/mockgen@v1.6.0
