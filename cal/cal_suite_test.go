package cal_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cal Suite")
}
