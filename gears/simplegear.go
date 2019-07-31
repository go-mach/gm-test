package gears

import (
	"time"

	"github.com/go-mach/machinery/pkg/machinery"
	"github.com/mitchellh/mapstructure"
)

// SimpleGearConf .
type SimpleGearConf struct {
	Name    string
	Version string
	Message string
}

// SimpleGear .
type SimpleGear struct {
	machinery.BaseGear
	canRun bool
	config SimpleGearConf
}

// Provided .
type Provided struct {
	Message string
}

// NewSimpleGear .
func NewSimpleGear(uname string) *SimpleGear {
	//return &SimpleGear{ConfigurableGear: machinery.NewConfigurableGear(uname, nil), canRun: false}
	//return &SimpleGear{uname: uname, canRun: false}
	return &SimpleGear{BaseGear: machinery.BaseGear{Uname: uname, Logger: nil}, canRun: false, config: SimpleGearConf{}}
}

// Start .
func (sg *SimpleGear) Start(m *machinery.Machinery) {
	sg.Logger.Printf("***************** the gear unique name is: %s", sg.Name())
	sg.canRun = true

	go func() {
		// i := 0
		// med := mediator.C
		for sg.canRun {
			sg.Logger.Printf("configuration: %v", sg.config)
			// i = i + 1
			// med <- i
			time.Sleep(2 * time.Second)
		}
	}()
}

// Configure .
func (sg *SimpleGear) Configure(config interface{}) {
	configMap := config.(map[string]interface{})
	mapstructure.Decode(configMap, &sg.config)
}

// Provide .
func (sg *SimpleGear) Provide() interface{} {
	//cfg := sg.Config.(map[string]interface{})
	return Provided{Message: sg.config.Message}
}
