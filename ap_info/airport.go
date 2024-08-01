package ap_info

//go:generate go-bindata -pkg $GOPACKAGE -o assets.go assets/

import (
    "encoding/csv"
    "bytes"
    "io"
    "strconv"
)

type Airport struct {
    Ident     string
    Name      string
    Country   string
    Region    string
    City      string
    Iata      string
    GpsCode   string
    Latitude  float64
    Longitude float64
}

func (apt Airport) Flag() (string) {
    iso_country_code := []rune(apt.Country)
    ri1 := (iso_country_code[0] - 0x41 + 0x1f1e6)
    ri2 := (iso_country_code[1] - 0x41 + 0x1f1e6)
    return string([]rune{ri1, ri2})
}

func lookup(f func(map[string]string) bool) (*Airport, error) {
    dataReader := bytes.NewReader(MustAsset("assets/airports.csv"))

    c := csv.NewReader(dataReader)

    headers, err := c.Read()
    if err != nil {
        return nil, err
    }

    var apt *Airport

    for {
        row, err := c.Read()
        if err == io.EOF {
            break
        }

        if err != nil {
            return nil, err
        }

        rec := make(map[string]string)
        for i, v := range row {
            rec[headers[i]] = v
        }

        if f(rec) {
            lat, err := strconv.ParseFloat(rec["latitude_deg"], 64)
            if err != nil {
                return nil, err
            }

            lon, err := strconv.ParseFloat(rec["longitude_deg"], 64)
            if err != nil {
                return nil, err
            }

            apt = &Airport{
                Ident:     rec["ident"],
                Name:      rec["name"],
                Country:   rec["iso_country"],
                Region:    rec["iso_region"],
                City:      rec["municipality"],
                Iata:      rec["iata_code"],
                GpsCode:   rec["gps_code"],
                Latitude:  lat,
                Longitude: lon,
            }

            break
        }
    }

    return apt, nil
}

func Lookup(code string) (*Airport, error) {
    return lookup(func(rec map[string]string) bool {
        return rec["ident"] == code
    })
}

func LookupByIata(iata string) (*Airport, error) {
    return lookup(func(rec map[string]string) bool {
        return rec["iata_code"] == iata
    })
}
