package api

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

	"github.com/neilsonwong/windrunner/tools"
	"github.com/spf13/viper"
)

// ListenAndServe listens on the requested port and serves the web api there, also proxies requests
func ListenAndServe() {
	proxyAddr := viper.GetString("LISTING_SERVER")
	serverPort := viper.GetInt("SERVER_PORT")
	port := strconv.Itoa(serverPort)

	//setup http server
	server := http.NewServeMux()

	proxyURL, _ := url.Parse(proxyAddr)
	proxyServer := httputil.NewSingleHostReverseProxy(proxyURL)

	fileOperator := tools.FileOperatorInstance()

	server.Handle("/api/", handleProxy(proxyServer, proxyURL))
	server.Handle("/play", handlePlay(&fileOperator))
	server.HandleFunc("/hello", handleHello)

	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)
}

func handleProxy(proxyServer *httputil.ReverseProxy, proxyURL *url.URL) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// log.Println(config.ListingServer)

		req.URL.Host = proxyURL.Host
		req.URL.Scheme = proxyURL.Scheme
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = proxyURL.Host

		proxyServer.ServeHTTP(res, req)
	})
}

func handleHello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, this is a gopher reporting in")
}

func handlePlay(fo *tools.FileOperator) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		enableCors(&res)
		if req.Method == "POST" {

			//Open(`//RASPBERRYPI/share/anime/Air/[Doki] Air - 01v2 (1280x720 h264 BD FLAC) [E13ADA79].mkv`)
			file := req.FormValue("file")

			log.Println(file)

			// ensure that share is mounted
			err := fo.MountSmb(true)

			if err != nil {
				fmt.Fprintf(res, "unable to mount "+fo.ShareName)
			} else {
				//perhaps cut the share out of filename in future? not sure
				fo.Open(file)
				fmt.Fprintf(res, "opened "+file)
			}
			return
		}

		fmt.Fprintf(res, "could not find the resource")
	})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
