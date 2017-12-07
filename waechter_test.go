package waechter_test

import (
	waechter "github.com/ElectricCookie/go-waechter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoWaechter(t *testing.T) {
	// Set the hash expense to 128  to prevent the test from taking long. DONT USE SUCH A LOW VALUE IN PRODUCTION!
	waechter.HashExpense = 128
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoWaechter Suite")
}
