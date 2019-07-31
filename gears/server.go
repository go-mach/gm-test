package gears

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-mach/machinery/pkg/machinery"
	"github.com/mitchellh/mapstructure"
)

// APIStarterFunc .
type APIStarterFunc func(*machinery.Machinery) APIInfo

// APIInfo .
type APIInfo struct {
	Router *chi.Mux
	Addr   string
}

// ServerGear .
type ServerGear struct {
	machinery.BaseGear
	starter APIStarterFunc
	config  APIConf
}

// NewServerGear .
func NewServerGear(uname string, starter APIStarterFunc) *ServerGear {
	return &ServerGear{BaseGear: machinery.BaseGear{Uname: uname, Logger: nil}, starter: starter, config: APIConf{}}
}

// APIConf is the API Gear configuration structure.
type APIConf struct {
	Endpoint struct {
		Port            int
		BaseRoutingPath string
	}
	Security struct {
		Enabled bool
		Jwt     struct {
			Secret     string
			Expiration struct {
				Enabled bool
				Minutes int32
			}
		}
	}
}

// Start .
func (sg *ServerGear) Start(m *machinery.Machinery) {
	router := chi.NewRouter()
	api := sg.starter(m)
	sg.Logger.Printf("baseRoutingPath config: %v", sg.config.Endpoint.BaseRoutingPath)
	router.Route(sg.config.Endpoint.BaseRoutingPath, func(r chi.Router) {
		r.Mount("/", api.Router)
	})

	log.Fatal(http.ListenAndServe(api.Addr, router))
}

// Configure get the configuration map and struct it into APIConf structure.
func (sg *ServerGear) Configure(config interface{}) {
	configMap := config.(map[string]interface{})
	mapstructure.Decode(configMap, &sg.config)
}

// Shutdown .
func (sg *ServerGear) Shutdown() {
	log.Printf("%s SHUT DOWN", sg.Uname)
}

// Provide .
func (sg *ServerGear) Provide() interface{} {
	return nil
}
