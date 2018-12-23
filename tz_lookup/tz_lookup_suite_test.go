package tz_lookup_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    "testing"
)

func TestSubmodule(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Submodule Suite")
}
