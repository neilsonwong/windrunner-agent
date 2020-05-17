package main

import (
	"github.com/neilsonwong/windrunner/api"
	"github.com/neilsonwong/windrunner/config"
	"github.com/neilsonwong/windrunner/tools"
)

func main() {
	// load and print configs
	config.Load().Print()

	//ensure we have mounted our shared folder
	tools.FileOperatorInstance().MountSmb(false)

	// start web server
	api.ListenAndServe()
}
