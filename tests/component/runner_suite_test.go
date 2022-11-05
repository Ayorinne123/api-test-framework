package component

import (
	"fmt"
	"goapitesting/utilities"
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
	utilities.Load_envconfig()

})

var _ = AfterSuite(func() {
	fmt.Println("After Suite")
})
