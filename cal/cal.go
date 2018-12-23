package cal

import (
    "fmt"
    "time"

    "github.com/arran4/golang-ical"
    "github.com/google/uuid"

    "github.com/blalor/flight-cal/ap_info"
    "github.com/blalor/flight-cal/tz_lookup"
)

func CreateFlightCal(flight, departAirport string, departTimeStr string, arriveAirport string, arriveTimeStr string) (*ics.Calendar, error) {
    const timeFmt = "2006-01-02T15:04"

    departApt, err := ap_info.LookupByIata(departAirport)
    if err != nil {
        return nil, err
    }

    if departApt == nil {
        return nil, fmt.Errorf("no such airport %s", departAirport)
    }

    departTz, err := tz_lookup.LookupTZ(departApt.Latitude, departApt.Longitude)
    if err != nil {
        return nil, err
    }

    departTime, err := time.ParseInLocation(timeFmt, departTimeStr, departTz)
    if err != nil {
        return nil, err
    }

    arriveApt, err := ap_info.LookupByIata(arriveAirport)
    if err != nil {
        return nil, err
    }

    if arriveApt == nil {
        return nil, fmt.Errorf("no such airport %s", arriveAirport)
    }

    arriveTz, err := tz_lookup.LookupTZ(arriveApt.Latitude, arriveApt.Longitude)
    if err != nil {
        return nil, err
    }

    arriveTime, err := time.ParseInLocation(timeFmt, arriveTimeStr, arriveTz)
    if err != nil {
        return nil, err
    }

    c := ics.NewCalendar()
    c.SetMethod(ics.MethodPublish)

    evt := c.AddEvent(uuid.New().String())
    evt.SetDtStampTime(time.Now())

    evt.SetStartAt(departTime)
    evt.SetEndAt(arriveTime)
    evt.SetSummary(fmt.Sprintf("✈️%s %s → %s", flight, departApt.Iata, arriveApt.Iata))
    evt.SetLocation(departApt.Name)

    return c, nil
}
