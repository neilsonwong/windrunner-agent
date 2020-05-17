package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/spf13/viper"
)

// ListenAndServe listens on the requested port and serves the web api there, also proxies requests
func ListenAndServe() {
	serverPort := viper.GetInt("SERVER_PORT")
	port := strconv.Itoa(serverPort)

	router := chi.NewRouter()
	// router.Use(middleware.Logger)
	router.Use(CORSMiddleware())
	router.Use(middleware.Timeout(60 * time.Second))

	router.Mount("/proxy", ProxyRouter())
	router.Mount("/api", AgentRouter())

	router.HandleFunc("/hello", handleHello)

	err := http.ListenAndServe(":"+port, router)
	log.Fatal(err)
}

func handleHello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, this is a gopher reporting in")
}
