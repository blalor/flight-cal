package ap_info_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestApInfo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ApInfo Suite")
}
