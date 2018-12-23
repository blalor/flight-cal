package tz_lookup_test

import (
    . "github.com/blalor/flight-cal/tz_lookup"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    // "github.com/stretchr/testify/mock"

    "time"
)

var _ = Describe("TZ Lookup", func() {
    It("does a lookup", func() {
        // KRIC
        ny, _ := time.LoadLocation("America/New_York")
        Expect(LookupTZ(37.50519943237305, -77.3197021484375)).To(Equal(ny))
    })
})
