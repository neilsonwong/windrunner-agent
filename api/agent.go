package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/neilsonwong/windrunner/tools"
)

// AgentRouter contains routes for basic windrunner agent functions
func AgentRouter() http.Handler {
	fileOperator := tools.FileOperatorInstance()
	r := chi.NewRouter()

	// define routes
	r.Post("/play", handlePlay(&fileOperator))
	r.Get("/doki", handleHeartbeat)

	return r
}

func handlePlay(fo *tools.FileOperator) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
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
}

func handleHeartbeat(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "ドキドキ")
}
