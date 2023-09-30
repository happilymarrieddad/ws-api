package wsclient_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWsclient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wsclient Suite")
}
