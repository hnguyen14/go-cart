package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoCart Suite")
}
