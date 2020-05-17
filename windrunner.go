package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/viper"

	"github.com/neilsonwong/windrunner/api"
)

func main() {
	loadConfig()

	api.ListenAndServe()

	// //ensure mount is successful
	// MountSmb(config.ShareServer, config.ShareFolder, false)

	// //setup http server
	// h := http.NewServeMux()

	// h.HandleFunc("/api/", handleRequestAndRedirect)

	// //setup the play function
	// h.HandleFunc("/play", func(resw http.ResponseWriter, req *http.Request) {
	// 	enableCors(&resw)

	// 	switch req.Method {
	// 	case "POST":
	// 		handlePlay(resw, req, config.ShareServer, config.ShareFolder)
	// 	case "GET":
	// 	case "PUT":
	// 	case "DELETE":
	// 	default:
	// 		fmt.Fprintf(resw, "could not find the resource")
	// 	}
	// })

	// h.HandleFunc("/hello", func(resw http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprintf(resw, "Hello, this is a gopher reporting in")
	// })

	// //setup proxy to fire to listing server
	// // hl := proxy(h, config.ListingServer)

	// err := http.ListenAndServe(":"+port, h)
	// log.Fatal(err)
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// print our config
	shareServer := viper.GetString("SHARE_SERVER")
	shareFolder := viper.GetString("SHARE_FOLDER")
	listingServer := viper.GetString("LISTING_SERVER")
	osxMountPoint := viper.GetString("OSX_MOUNT")
	serverPort := viper.GetInt("SERVER_PORT")
	sharename := "//" + shareServer + "/" + shareFolder
	port := strconv.Itoa(serverPort)

	log.Println("config: share located at " + sharename)
	log.Println("config: listing server at " + listingServer)
	log.Println("config: osx mount point at " + osxMountPoint)
	log.Println("config: agent hosted on port " + port)
}
