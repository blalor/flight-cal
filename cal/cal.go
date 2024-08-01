package cal

import (
    "fmt"
    "time"
    "strings"

    "github.com/arran4/golang-ical"
    "github.com/google/uuid"

    "github.com/blalor/flight-cal/ap_info"
    "github.com/blalor/flight-cal/tz_lookup"
)

const timeFmtLong = "2006-01-02T15:04"
const timeFmtShort = "Mon 15:04"
const timeFmtTime = "15:04"
const timeFmtDate = "2006-01-02"

func changeDay(t time.Time, dir time.Duration) time.Time {
    loc := t.Location()
    day := t.Format(timeFmtDate)
    tmp := t.Add(time.Hour * dir)
    for tmp.Format(timeFmtDate) == day {
        tmp = tmp.Add(time.Hour * dir)
    }

    timeStr := fmt.Sprintf("%sT%s", tmp.Format(timeFmtDate), t.Format(timeFmtTime))
    t2, _ := time.ParseInLocation(timeFmtLong, timeStr, loc)
    return t2
}

func prevDay(t time.Time) time.Time {
    return changeDay(t, -1)
}

func nextDay(t time.Time) time.Time {
    return changeDay(t, 1)
}

func CreateFlightCal(prefix string, flight, record string, departAirport string, departTimeStr string, arriveAirport string, arriveTimeStr string) (*ics.Calendar, error) {
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

    departTime, err := time.ParseInLocation(timeFmtLong, departTimeStr, departTz)
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

    arriveTime, err := time.ParseInLocation(timeFmtLong, arriveTimeStr, arriveTz)
    if err != nil {
        departDayStr := departTime.Format(timeFmtDate)
        maybeTimeStr := fmt.Sprintf("%sT%s", departDayStr, arriveTimeStr)
        arriveTime, err = time.ParseInLocation(timeFmtLong, maybeTimeStr, arriveTz)
        if err != nil {
            return nil, err
        }
    }

    duration := arriveTime.Sub(departTime).Hours()

    if duration < -22 {
        arriveTime = nextDay(nextDay(arriveTime))
        duration = arriveTime.Sub(departTime).Hours()
    } else if duration < 0 {
        arriveTime = nextDay(arriveTime)
        duration = arriveTime.Sub(departTime).Hours()
    } else if duration > 22 {
        arriveTime = prevDay(arriveTime)
        duration = arriveTime.Sub(departTime).Hours()
    }

    dur := arriveTime.Sub(departTime)
    durStr := fmt.Sprintf("%dh%02dm", int(dur.Hours()), int(dur.Minutes()) % 60)

    var sb strings.Builder
    sb.WriteString("🛫 ")
    sb.WriteString(departApt.City)
    sb.WriteString(" ")
    sb.WriteString(departApt.Iata)
    sb.WriteString(" ")
    sb.WriteString(departTime.Format(timeFmtShort))
    sb.WriteString(" (local time)\\n")
    sb.WriteString("🛬 ")
    sb.WriteString(arriveApt.City)
    sb.WriteString(" ")
    sb.WriteString(arriveApt.Iata)
    sb.WriteString(" ")
    sb.WriteString(arriveTime.Format(timeFmtShort))
    sb.WriteString(" (local time)\\n")
    sb.WriteString("⏱️ ")
    sb.WriteString(durStr)

    if record != "" {
        sb.WriteString("\\n")
        sb.WriteString("🎫 ")
        sb.WriteString(record)
    }

    c := ics.NewCalendar()
    c.SetMethod(ics.MethodPublish)

    evt := c.AddEvent(uuid.New().String())
    evt.SetDtStampTime(time.Now())

    evt.SetStartAt(departTime)
    evt.SetEndAt(arriveTime)
    evt.SetSummary(fmt.Sprintf("%s %s %s → %s", prefix, flight, departApt.Iata, arriveApt.Iata))
    evt.SetLocation(departApt.Name)
    desc := sb.String()
    evt.SetDescription(desc)
    fmt.Println(desc)

    return c, nil
}
