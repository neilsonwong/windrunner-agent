package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/neilsonwong/windrunner/config"
)

// ConfigRouter handles viewing and editing agent config settings
func ConfigRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(CORSMiddleware())

	r.Get("/", handleGetConfig)
	r.Put("/", handleEditConfig)

	return r
}

func handleGetConfig(res http.ResponseWriter, req *http.Request) {
	conf := config.Get()
	res.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(res).Encode(conf)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleEditConfig(res http.ResponseWriter, req *http.Request) {
	var conf config.Config

	err := json.NewDecoder(req.Body).Decode(&conf)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	config.Update(&conf)
}
