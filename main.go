package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-mach/gm-test/gears"
	"github.com/go-mach/machinery-api/pkg/apigear"

	"github.com/go-mach/machinery/pkg/machinery"
)

// // HealthGear .
// type HealthGear struct {
// 	*machinery.BaseGear
// 	Starter func(string, interface{})
// }

// // NewHealthGear returns a new WelcomeGear instance.
// func NewHealthGear(uname string, starter func(string, interface{})) *HealthGear {
// 	return &HealthGear{
// 		BaseGear: machinery.NewBaseGear(uname),
// 		Starter:  starter,
// 	}
// 	//return &HealthGear{machinery.NewBaseGear(uname)}
// }

// // Start startup the health edpoint
// func (hg *HealthGear) Start(m *machinery.Machinery) {
// 	simpleGear := m.GetGear("simple-gear").(gears.SimpleGear)
// 	hg.Starter(hg.UniqueName, &simpleGear)
// 	healthAddr := fmt.Sprintf(":%d", 7777)
// 	go func(addr string) {
// 		if err := health.ServeDefault(addr); err != nil {
// 			log.Fatalf("[main] serveHealth - %s", err.Error())
// 			os.Exit(1)
// 		}
// 	}(healthAddr)
// }

// // Shutdown startup the health edpoint
// func (hg *HealthGear) Shutdown() {
// 	log.Printf("\nHEALTH GEAR SHUTDOWN for signal")
// }

// // Provide .
// func (hg *HealthGear) Provide() interface{} {
// 	return "Hi man, you're welcome!!!"
// }

// // SampleConf .
// type SampleConf struct {
// 	Name    string
// 	Version string
// }

// func composeHealthGear(uname string, simpleGear interface{}) {
// 	if _, ok := simpleGear.(machinery.Configurable); ok {
// 		log.Print("OK")
// 	} else {
// 		log.Print("KO")
// 	}

// 	log.Printf("*** the Health Gear uname is %s ***", uname)
// 	// log.Printf("*** the injected Simple-Gear uname is %s and its canrun is ***", simpleGear.Name())
// }

func apiStarter(m *machinery.Machinery) apigear.APIInfo {
	logger := m.Logger
	logger.Info("******************** api starter ***************")
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		sg := m.GetGear("simple-gear").(*gears.SimpleGear)
		msg, ok := sg.Provide().(gears.Provided)
		if !ok {
			logger.Error("error getting simple-gear provided object")
			w.Write([]byte("ERROR"))
		}
		w.Write([]byte(fmt.Sprintf("It Works: simplegear provided us the message %s, so..... YEPPAAAAAAAA2!!!!", msg.Message)))
		//w.Write([]byte("YEPPA!"))
	})

	middlewares := []func(http.Handler) http.Handler{middleware.Recoverer}

	return apigear.APIInfo{
		Middlewares: middlewares,
		Router:      router,
	}
}

func useMiddlewares() []func(http.Handler) http.Handler {
	return make([]func(http.Handler) http.Handler, 0)
}

func main() {
	fmt.Println(strings.ToUpper("gopher"))
	machinery.NewMachinery().With(
		//NewHealthGear("HealthGear", composeHealthGear),
		gears.NewSimpleGear("simple-gear"),
		apigear.NewAPIGear("api", apiStarter),
		// gears.NewServerGear("api", apiStarter),
	).Start()
	// c := make(chan int)
	// go func() {
	// 	c <- 42
	// }()

	// val := <-c
	// log.Println(val)
}
