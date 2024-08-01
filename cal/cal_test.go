package cal_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/blalor/flight-cal/cal"

    "github.com/arran4/golang-ical"
)

// https://icalendar.org/validator.html#results

var _ = Describe("Cal", func() {
    It("does … something?", func() {

        c, err := CreateFlightCal("✈️", "AA 1234", "", "RIC", "2018-08-01T05:01", "SFO", "2018-08-01T08:42")
        Expect(err).To(BeNil())

        Expect(c.Events()).To(HaveLen(1))

        evt := c.Events()[0]
        Expect(evt.GetProperty(ics.ComponentPropertyDtStart).Value).To(Equal("20180801T090100Z"))
        Expect(evt.GetProperty(ics.ComponentPropertyDtEnd).Value).To(Equal("20180801T154200Z"))
        Expect(evt.GetProperty(ics.ComponentPropertySummary).Value).To(Equal("✈️ AA 1234 RIC → SFO"))
        Expect(evt.GetProperty(ics.ComponentPropertyLocation).Value).To(Equal("Richmond International Airport")) // ¯\_(ツ)_/¯
    })
})
