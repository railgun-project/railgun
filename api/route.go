package api

import (
	"github.com/gorilla/mux"
	"github.com/railgun-project/railgun/assets"
	_ "github.com/railgun-project/railgun/assets/statik"
	"github.com/spf13/viper"
	"net/http"
)

func routes() *mux.Router {
	r := mux.NewRouter()
	r.Use(AuthMdiddleware(viper.GetString("dashboard.user"), viper.GetString("dashboard.password")).Middleware)
	api := r.PathPrefix("/api").Subrouter()
	api.Path("/conf/get_keys").HandlerFunc(GetKeysHandle).Methods("GET")
	api.Path("/conf/get_value").HandlerFunc(GetValueHandle).Methods("GET")
	api.Path("/conf/save_value").HandlerFunc(SaveValueHandle).Methods("PUT")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(assets.FileSystem))).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/", http.StatusMovedPermanently)
	})
	return r
}
