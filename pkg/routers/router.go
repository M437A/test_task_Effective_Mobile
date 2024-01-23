package routers

import (
	"name_service/pkg/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateRouters(router *chi.Mux) {

	router.Route("/", func(r chi.Router) {
		r.Post("/", controllers.SaveNewName)

		r.Get("/", controllers.GetUserData)

		r.Put("/", controllers.UpdateUserData)

		r.Delete("/{userID}", controllers.DeleteUserData)
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	controllers.Router = router
}
