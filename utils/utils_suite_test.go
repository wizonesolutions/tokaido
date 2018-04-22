package utils_test

import (
	"./"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Utils Suite")
}

var _ = BeforeSuite(func() {
	utils.FatalError = func(err error) string {
		return "utils.FatalError called"
	}
})
