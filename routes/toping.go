package routes

import (
	"waysbuck/handlers"
	"waysbuck/pkg/middleware"
	"waysbuck/pkg/mysql"
	"waysbuck/repositories"

	"github.com/gorilla/mux"
)

func TopingRoutes(r *mux.Router) {
	TopingRepository := repositories.RepositoryToping(mysql.DB)
	h := handlers.HandlerToping(TopingRepository)

	r.HandleFunc("/topings", h.FindTopings).Methods("GET")
	r.HandleFunc("/toping/{id}", h.GetToping).Methods("GET")
	r.HandleFunc("/toping", middleware.Auth(middleware.UploadFile(h.CreateToping))).Methods("POST")
}
