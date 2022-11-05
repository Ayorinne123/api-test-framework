package component

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestComponentTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Component Testing Suite")
}

var _ = BeforeSuite(func() {
	fmt.Println("Before Suite")
})

var _ = AfterSuite(func() {
	fmt.Println("After Suite")
})
