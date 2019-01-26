# flight-cal

This is a little utility that creates calendar (`.ics`) files for airline flights.  It uses the departure and arrival airport locations to determine the correct timezone offset so that the departure arrival times -- input in the airport's local timezone -- are captured absolutely as UTC.  This makes it so the calendar entries have the correct duration.  I wrote it because flight confirmation emails rarely have either a calendar file attachment nor are the timezone offsets (or even timezone names) presented; this makes it tedious and error-prone if you're math- and/or geography-challenged, such as myself.

`flight-cal` works completely offline, using [go-bindata](https://github.com/kevinburke/go-bindata) to embed the list of airports from [ourairports.com](http://ourairports.com/data/airports.csv) and [go-tz](https://github.com/ugjka/go-tz) to look up the timezones from airport locations.

## building

    make

Creates `stage/flight-cal`.

## usage

    stage/flight-cal \
        --flight='AA 583' \
        --depart-airport=CLT \
        --depart-time=2019-01-25T14:05 \
        --arrive-airport=RIC \
        --arrive-time=2019-01-25T15:10

Generates `20190125T190500Z AA 583.ics`. Add that to your calendar!

# history

I originally tried to do this with a Workflow (now Shortcut) on iOS, but could never get the timezone correction to work. 

The steps were:

* ask for carrier & flight
* ask for departure airport
* look up airport data from `http://www.airport-data.com/api/ap_info.json?iata=<DepartureAirport>`
* use airport location to find its timezone: `http://api.timezonedb.com/v2/get-time-zone?key=…&format=json&by=position&lat=<DepartureAirportInfo (longitude)>&lng=<DepartureAirportInfo (latitude)>``
* ask for departure time
* localize departure time using the airport's timezone
* repeat for arrival airport
