package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rha7/kubeapi/appinfo"
	"github.com/rha7/kubeapi/controllers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func main() {

	// Controllers //
	homeController := controllers.NewHome()

	r := mux.NewRouter()

	r.HandleFunc("/", homeController.Home)

	bindHost := "0.0.0.0"
	if host := os.Getenv("APP_BIND_HOST"); host != "" {
		bindHost = host
	}

	bindPort := "1690"
	if port := os.Getenv("APP_BIND_PORT"); port != "" {
		bindPort = port
	}

	bindAddress := fmt.Sprintf("%s:%s", bindHost, bindPort)

	logrus.
		WithField("version", appinfo.Version).
		WithField("build_time", appinfo.BuildTime).
		WithField("bind_address", bindAddress).
		Info("starting")

	mwCORS := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
			"HEAD",
			"CONNECT",
		},
	})

	n := negroni.Classic()
	n.Use(mwCORS)
	n.UseHandler(r)

	http.ListenAndServe(bindAddress, n)
}
