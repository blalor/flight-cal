package tz_lookup

import (
    "github.com/ugjka/go-tz/v2"
    "time"
)

func LookupTZ(lat, lon float64) (*time.Location, error) {
    zones, err := tz.GetZone(tz.Point{Lon: lon, Lat: lat})
    if err != nil {
        return nil, err
    }

    loc, err := time.LoadLocation(zones[0])
    if err != nil {
        return nil, err
    }

    return loc, nil
}
