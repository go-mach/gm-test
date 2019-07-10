package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-mach/gm-health/pkg/health"
	simple "github.com/go-mach/gm-simple-gear"
	"github.com/go-mach/gm/gm"
)

// HealthGear .
type HealthGear struct {
	gm.BaseGear
}

// NewHealthGear returns a new WelcomeGear instance.
func NewHealthGear(uname string) *HealthGear {
	return &HealthGear{gm.NewBaseGear(uname)}
}

// Start startup the health edpoint
func (hg *HealthGear) Start(machinery *gm.Machinery) {
	healthAddr := fmt.Sprintf(":%d", 7777)
	go func(addr string) {
		if err := health.ServeDefault(addr); err != nil {
			log.Fatalf("[main] serveHealth - %s", err.Error())
			os.Exit(1)
		}
	}(healthAddr)
}

// Shutdown startup the health edpoint
func (hg *HealthGear) Shutdown() {
	log.Printf("\nHEALTH GEAR SHUTDOWN for signal")
}

// Provide .
func (hg *HealthGear) Provide() interface{} {
	return "Hi man, you're welcome!!!"
}

// SampleConf .
type SampleConf struct {
	Name    string
	Version string
}

func main() {
	fmt.Println(strings.ToUpper("gopher"))

	gm.NewMachinery().
		With(
			&simple.UUIDGear{},
			simple.NewSimpleGear("simple-gear", nil),
			NewHealthGear("HealthGear")).
		Start()
}
