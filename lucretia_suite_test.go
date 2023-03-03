package lucretia_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLucretia(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lucretia Suite")
}
