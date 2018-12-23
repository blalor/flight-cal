package ap_info_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/blalor/flight-cal/ap_info"
)

var _ = Describe("ApInfo", func() {
    It("looks up an airport", func() {
        Expect(Lookup("KRIC")).NotTo(BeNil())
    })

    It("return nil on missing airport", func() {
        Expect(Lookup("TOTALLY_NOT_AN_AIRPORT")).To(BeNil())
    })
})
