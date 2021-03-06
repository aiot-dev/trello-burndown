package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/swordbeta/trello-burndown/src/backend"
)

// Delete deletes a trello board!
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := backend.GetDatabase()
	defer db.Close()
	db.Delete(&backend.Board{
		ID: vars["board"],
	})
	http.Redirect(w, r, viper.GetString("http.baseURL")+"index", 302)
}
