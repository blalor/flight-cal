package main

import (
    "os"
    "fmt"
    "syscall"

    flags "github.com/jessevdk/go-flags"
    log "github.com/sirupsen/logrus"

    "github.com/blalor/flight-cal/cal"
)

var version string = "undef"

type Options struct {
    Debug bool     `env:"DEBUG"    long:"debug"    description:"enable debug"`
    LogFile string `env:"LOG_FILE" long:"log-file" description:"path to JSON log file"`

    FlightDescription string `long:"flight"         description:"flight description (eg. AA 1234)" required:"y"`
    DepartAirport     string `long:"depart-airport" description:"departure airport" required:"y"`
    DepartTime        string `long:"depart-time"    description:"departure time, 2006-01-02T15:04" required:"y"`
    ArriveAirport     string `long:"arrive-airport" description:"arrival airport" required:"y"`
    ArriveTime        string `long:"arrive-time"    description:"arrival time, 2006-01-02T15:04" required:"y"`

    OutFile string `long:"output" description:"output file" required:"y"`
}

func main() {
    var opts Options

    _, err := flags.Parse(&opts)
    if err != nil {
        os.Exit(1)
    }

    if opts.Debug {
        log.SetLevel(log.DebugLevel)
    }

    if opts.LogFile != "" {
        logFp, err := os.OpenFile(opts.LogFile, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0600)
        checkError(fmt.Sprintf("error opening %s", opts.LogFile), err)

        defer logFp.Close()

        // ensure panic output goes to log file
        syscall.Dup2(int(logFp.Fd()), 1)
        syscall.Dup2(int(logFp.Fd()), 2)

        // log as JSON
        log.SetFormatter(&log.JSONFormatter{})

        // send output to file
        log.SetOutput(logFp)
    }

    log.Debug("hi there! (tickertape tickertape)")
    log.Infof("version: %s", version)

    of, err := os.Create(opts.OutFile)
    if err != nil {
        log.Fatalf("unable to open %s: %v", opts.OutFile, err)
    }

    defer of.Close()
    c, err := cal.CreateFlightCal(
        opts.FlightDescription,
        opts.DepartAirport,
        opts.DepartTime,
        opts.ArriveAirport,
        opts.ArriveTime,
    )

    if err != nil {
        log.Fatalf("unable to create event: %v", err)
    }

    _, err = of.WriteString(c.Serialize())
    if err != nil {
        log.Fatalf("unable to write data: %v", err)
    }
}
