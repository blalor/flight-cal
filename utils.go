package main

import (
    log "github.com/sirupsen/logrus"
)

func checkError(msg string, err error) {
    if err != nil {
        log.Fatalf("%s: %+v", msg, err)
    }
}
