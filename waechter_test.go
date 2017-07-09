package waechter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoWaechter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoWaechter Suite")
}
