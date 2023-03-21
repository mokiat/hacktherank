package testdata_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestData Suite")
}
