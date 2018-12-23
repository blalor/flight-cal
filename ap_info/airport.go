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
    Iata      string
    GpsCode   string
    Latitude  float64
    Longitude float64
}

func Lookup(code string) (*Airport, error) {
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

        if rec["ident"] == code {
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

func LookupByIata(iata string) (*Airport, error) {
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

        if rec["iata_code"] == iata {
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
