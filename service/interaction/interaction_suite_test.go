package interactionimport_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInteraction(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interaction Suite")
}
